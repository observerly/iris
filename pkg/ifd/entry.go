/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/metadata
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package metadata

/*****************************************************************************************************************/

import "encoding/binary"

/*****************************************************************************************************************/

// An IFDEntry is a single entry in an Image File Directory.
// A value of type DataTypeRational is composed of two 32-bit values,
// thus data contains two uints (numerator and denominator) for a single number.
type IFDEntry struct {
	Tag      TagType
	DataType DataType
	Data     []uint32
}

/*****************************************************************************************************************/

func (e IFDEntry) PutData(p []byte) {
	enc := binary.LittleEndian

	for _, d := range e.Data {
		switch e.DataType {
		case DataTypeByte, DataTypeASCII:
			p[0] = byte(d)
			p = p[1:]
		case DataTypeShort:
			enc.PutUint16(p, uint16(d))
			p = p[2:]
		case DataTypeLong, DataTypeRational:
			enc.PutUint32(p, uint32(d))
			p = p[4:]
		}
	}
}

/*****************************************************************************************************************/
