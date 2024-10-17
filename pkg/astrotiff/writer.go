/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/astrotiff
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package astrotiff

/*****************************************************************************************************************/

import (
	"io"
)

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
