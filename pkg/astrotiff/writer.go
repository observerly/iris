/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/astrotiff
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package astrotiff

/*****************************************************************************************************************/

import (
	"encoding/binary"
	"io"
	"sort"

	metadata "github.com/observerly/iris/pkg/ifd"
)

/*****************************************************************************************************************/

// We only write little-endian TIFF files.
var enc = binary.LittleEndian

/*****************************************************************************************************************/

// writePixels writes the internal byte array of an image to w. It is less general
// but much faster then encode. writePixels is used when pix directly
// corresponds to one of the TIFF image types.
func writePixels(w io.Writer, pix []byte, nrows, length, stride int) error {
	if length == stride {
		_, err := w.Write(pix[:nrows*length])
		return err
	}

	for ; nrows > 0; nrows-- {
		if _, err := w.Write(pix[:length]); err != nil {
			return err
		}
		pix = pix[stride:]
	}

	return nil
}

/*****************************************************************************************************************/

// writeIFD writes the Image File Directory to w. The IFD is written at the given offset.
// The IFD is written in ascending order of the tag values.
func writeIFD(w io.Writer, ifdOffset int, d []metadata.IFDEntry) error {
	var buf [metadata.IFDLengthInBytes]byte
	// Make space for "pointer area" containing IFD entry data
	// longer than 4 bytes.
	parea := make([]byte, 1024)
	pstart := ifdOffset + metadata.IFDLengthInBytes*len(d) + 6
	var o int // Current offset in parea.

	// The IFD has to be written with the tags in ascending order.
	sort.Sort(metadata.SortByTagInterface(d))

	// Write the number of entries in this IFD.
	if err := binary.Write(w, enc, uint16(len(d))); err != nil {
		return err
	}

	for _, entry := range d {
		enc.PutUint16(buf[0:2], uint16(entry.Tag))
		enc.PutUint16(buf[2:4], uint16(entry.DataType))
		count := uint32(len(entry.Data))
		if entry.DataType == metadata.DataTypeRational {
			count /= 2
		}
		enc.PutUint32(buf[4:8], count)
		datalen := int(count * uint32(entry.DataType.ByteSize()))

		if datalen <= 4 {
			entry.PutData(buf[8:12])
		} else {
			if (o + datalen) > len(parea) {
				newlen := len(parea) + 1024
				for (o + datalen) > newlen {
					newlen += 1024
				}
				newarea := make([]byte, newlen)
				copy(newarea, parea)
				parea = newarea
			}
			entry.PutData(parea[o : o+datalen])
			enc.PutUint32(buf[8:12], uint32(pstart+o))
			o += datalen
		}

		if _, err := w.Write(buf[:]); err != nil {
			return err
		}
	}

	// The IFD ends with the offset of the next IFD in the file,
	// or zero if it is the last one (page 14).
	if err := binary.Write(w, enc, uint32(0)); err != nil {
		return err
	}

	_, err := w.Write(parea[:o])

	return err
}

/*****************************************************************************************************************/
