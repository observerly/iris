/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/metadata
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package metadata

/*****************************************************************************************************************/

type DataType uint16

/*****************************************************************************************************************/

const (
	DataTypeNil       DataType = 0  // placeholder, invalid
	DataTypeByte      DataType = 1  // 8-bit unsigned integer
	DataTypeASCII     DataType = 2  // 8-bit bytes w/ last byte null
	DataTypeShort     DataType = 3  // 16-bit unsigned integer
	DataTypeLong      DataType = 4  // 32-bit unsigned integer
	DataTypeRational  DataType = 5  // 64-bit unsigned fraction
	DataTypeSByte     DataType = 6  // !8-bit signed integer
	DataTypeUndefined DataType = 7  // !8-bit untyped data
	DataTypeSShort    DataType = 8  // !16-bit signed integer
	DataTypeSLong     DataType = 9  // !32-bit signed integer
	DataTypeSRational DataType = 10 // !64-bit signed fraction
	DataTypeFloat     DataType = 11 // !32-bit IEEE floating point
	DataTypeDouble    DataType = 12 // !64-bit IEEE floating point
	DataTypeIFD       DataType = 13 // %32-bit unsigned integer (offset)
	DataTypeLong8     DataType = 16 // BigTIFF 64-bit unsigned integer
	DataTypeSLong8    DataType = 17 // BigTIFF 64-bit signed integer
	DataTypeIFD8      DataType = 18 // BigTIFF 64-bit unsigned integer (offset)
)

/*****************************************************************************************************************/
