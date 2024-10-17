/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/metadata
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package metadata

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
