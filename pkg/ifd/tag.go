/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/metadata
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package metadata

/*****************************************************************************************************************/

type TagType uint16

/*****************************************************************************************************************/

type (
	TagValueNewSubfileType     TagType
	TagValueSubfileType        TagType
	TagValueCompressionType    TagType
	TagValuePhotometricType    TagType
	TagValuePredictorType      TagType
	TagValueResolutionUnitType TagType
	TagValueSampleFormatType   TagType
)

/*****************************************************************************************************************/

// Type(A/B/C/*), Num(1/*), Required, # comment
const (
	TagTypeNewSubfileType                 TagType                    = 254   // LONG , 1, # Default=0. subfile data descriptor
	TagValueNewSubfileTypeNil             TagValueNewSubfileType     = 0     //
	TagValueNewSubfileTypeReduced         TagValueNewSubfileType     = 1     // # bit0, reduced resolution version
	TagValueNewSubfileTypePage            TagValueNewSubfileType     = 2     // # bit1, one page of many
	TagValueNewSubfileTypeReducedPage     TagValueNewSubfileType     = 3     //
	TagValueNewSubfileTypeMask            TagValueNewSubfileType     = 4     // # bit2, transparency mask
	TagValueNewSubfileTypeReducedMask     TagValueNewSubfileType     = 5     //
	TagValueNewSubfileTypePageMask        TagValueNewSubfileType     = 6     //
	TagValueNewSubfileTypeReducedPageMask TagValueNewSubfileType     = 7     //
	TagTypeSubfileType                    TagType                    = 255   // SHORT, 1, # kind of data in subfile
	TagValueSubfileTypeImage              TagValueSubfileType        = 1     // # full resolution image data
	TagValueSubfileTypeReducedImage       TagValueSubfileType        = 2     // # reduced size image data
	TagValueSubfileTypePage               TagValueSubfileType        = 3     // # one page of many
	TagTypeImageWidth                     TagType                    = 256   // SHORT/LONG/LONG8, 1, # Required
	TagTypeImageLength                    TagType                    = 257   // SHORT/LONG/LONG8, 1, # Required
	TagTypeBitsPerSample                  TagType                    = 258   // SHORT, *, # Default=1. See SamplesPerPixel
	TagTypeCompression                    TagType                    = 259   // SHORT, 1, # Default=1
	TagValueCompressionTypeNil            TagValueCompressionType    = 0     //
	TagValueCompressionTypeNone           TagValueCompressionType    = 1     //
	TagValueCompressionTypeCCITT          TagValueCompressionType    = 2     //
	TagValueCompressionTypeG3             TagValueCompressionType    = 3     // # Group 3 Fax.
	TagValueCompressionTypeG4             TagValueCompressionType    = 4     // # Group 4 Fax.
	TagValueCompressionTypeLZW            TagValueCompressionType    = 5     //
	TagValueCompressionTypeJPEGOld        TagValueCompressionType    = 6     // # Superseded by cJPEG.
	TagValueCompressionTypeJPEG           TagValueCompressionType    = 7     //
	TagValueCompressionTypeDeflate        TagValueCompressionType    = 8     // # zlib compression.
	TagValueCompressionTypePackBits       TagValueCompressionType    = 32773 //
	TagValueCompressionTypeDeflateOld     TagValueCompressionType    = 32946 // # Superseded by cDeflate.
	TagTypePhotometricInterpretation      TagType                    = 262   // SHORT, 1,
	TagValuePhotometricTypeWhiteIsZero    TagValuePhotometricType    = 0     //
	TagValuePhotometricTypeBlackIsZero    TagValuePhotometricType    = 1     //
	TagValuePhotometricTypeRGB            TagValuePhotometricType    = 2     //
	TagValuePhotometricTypePaletted       TagValuePhotometricType    = 3     //
	TagValuePhotometricTypeTransMask      TagValuePhotometricType    = 4     // # transparency mask
	TagValuePhotometricTypeCMYK           TagValuePhotometricType    = 5     //
	TagValuePhotometricTypeYCbCr          TagValuePhotometricType    = 6     //
	TagValuePhotometricTypeCIELab         TagValuePhotometricType    = 8     //
	TagTypeThreshholding                  TagType                    = 263   // SHORT, 1, # Default=1
	TagTypeCellWidth                      TagType                    = 264   // SHORT, 1,
	TagTypeCellLenght                     TagType                    = 265   // SHORT, 1,
	TagTypeFillOrder                      TagType                    = 266   // SHORT, 1, # Default=1
	TagTypeDocumentName                   TagType                    = 269   // ASCII
	TagTypeImageDescription               TagType                    = 270   // ASCII
	TagTypeMake                           TagType                    = 271   // ASCII
	TagTypeModel                          TagType                    = 272   // ASCII
	TagTypeStripOffsets                   TagType                    = 273   // SHORT/LONG/LONG8, *, # StripsPerImage
	TagTypeOrientation                    TagType                    = 274   // SHORT, 1, # Default=1
	TagTypeSamplesPerPixel                TagType                    = 277   // SHORT, 1, # Default=1
	TagTypeRowsPerStrip                   TagType                    = 278   // SHORT/LONG/LONG8, 1,
	TagTypeStripByteCounts                TagType                    = 279   // SHORT/LONG/LONG8, *, # StripsPerImage
	TagTypeMinSampleValue                 TagType                    = 280   // SHORT,    *, # Default=0
	TagTypeMaxSampleValue                 TagType                    = 281   // SHORT,    *, # Default=2^BitsPerSample-1
	TagTypeXResolution                    TagType                    = 282   // RATIONAL, 1, # Required?
	TagTypeYResolution                    TagType                    = 283   // RATIONAL, 1, # Required?
	TagTypePlanarConfiguration            TagType                    = 284   // SHORT,    1, # Defaule=1
	TagTypePageName                       TagType                    = 285   // ASCII
	TagTypeXPosition                      TagType                    = 286   // RATIONAL,   1
	TagTypeYPosition                      TagType                    = 287   // RATIONAL,   1
	TagTypeFreeOffsets                    TagType                    = 288   // LONG/LONG8, *
	TagTypeFreeByteCounts                 TagType                    = 289   // LONG/LONG8, *
	TagTypeGrayResponseUnit               TagType                    = 290   // SHORT, 1,
	TagTypeGrayResponseCurve              TagType                    = 291   // SHORT, *, # 2**BitPerSample
	TagTypeT4Options                      TagType                    = 292   // LONG,  1, # Default=0
	TagTypeT6Options                      TagType                    = 293   // LONG,  1, # Default=0
	TagTypeResolutionUnit                 TagType                    = 296   // SHORT, 1, # Default=2
	TagValueResolutionUnitTypeNone        TagValueResolutionUnitType = 1     //
	TagValueResolutionUnitTypePerInch     TagValueResolutionUnitType = 2     // # Dots per inch.
	TagValueResolutionUnitTypePerCM       TagValueResolutionUnitType = 3     // # Dots per centimeter.
	TagTypePageNumber                     TagType                    = 297   // SHORT, 2,
	TagTypeTransferFunction               TagType                    = 301   // SHORT, *, # {1 or SamplesPerPixel}*2**BitPerSample
	TagTypeSoftware                       TagType                    = 305   // ASCII
	TagTypeDateTime                       TagType                    = 306   // ASCII, 20, # YYYY:MM:DD HH:MM:SS, include NULL
	TagTypeArtist                         TagType                    = 315   // ASCII
	TagTypeHostComputer                   TagType                    = 316   // ASCII
	TagTypePredictor                      TagType                    = 317   // SHORT, 1, # Default=1
	TagValuePredictorTypeNone             TagValuePredictorType      = 1     //
	TagValuePredictorTypeHorizontal       TagValuePredictorType      = 2     //
	TagTypeWhitePoint                     TagType                    = 318   // RATIONAL, 2
	TagTypePrimaryChromaticities          TagType                    = 319   // RATIONAL, 6
	TagTypeColorMap                       TagType                    = 320   // SHORT, *, # 3*(2**BitPerSample)
	TagTypeHalftoneHints                  TagType                    = 321   // SHORT, 2
	TagTypeTileWidth                      TagType                    = 322   // SHORT/LONG, 1
	TagTypeTileLength                     TagType                    = 323   // SHORT/LONG, 1
	TagTypeTileOffsets                    TagType                    = 324   // LONG/LONG8, *, # TilesPerImage
	TagTypeTileByteCounts                 TagType                    = 325   // SHORT/LONG, *, # TilesPerImage
	TagTypeBadFaxLines                    TagType                    = 326   // ingore # Used in the TIFF-F standard, denotes the number of 'bad' scan lines encountered by the facsimile device.
	TagTypeCleanFaxData                   TagType                    = 327   // ingore # Used in the TIFF-F standard, indicates if 'bad' lines encountered during reception are stored in the data, or if 'bad' lines have been replaced by the receiver.
	TagTypeConsecutiveBadFaxLines         TagType                    = 328   // ingore # Used in the TIFF-F standard, denotes the maximum number of consecutive 'bad' scanlines received.
	TagTypeSubIFD                         TagType                    = 330   // IFD,   *  # IFD pointer
	TagTypeInkSet                         TagType                    = 332   // SHORT, 1, # Default=1
	TagTypeInkNames                       TagType                    = 333   // ASCII
	TagTypeNumberOfInks                   TagType                    = 334   // SHORT, 1, # Default=4
	TagTypeDotRange                       TagType                    = 336   // BYTE/SHORT, # Default=[0,2^BitsPerSample-1]
	TagTypeTargetPrinter                  TagType                    = 337   // ASCII
	TagTypeExtraSamples                   TagType                    = 338   // BYTE,  1,
	TagTypeSampleFormat                   TagType                    = 339   // SHORT, *, # SamplesPerPixel. Default=1
	TagValueSampleFormatTypeUint          TagValueSampleFormatType   = 1     //
	TagValueSampleFormatTypeTwoInt        TagValueSampleFormatType   = 2     //
	TagValueSampleFormatTypeFloat         TagValueSampleFormatType   = 3     //
	TagValueSampleFormatTypeUndefined     TagValueSampleFormatType   = 4     //
	TagTypeSMinSampleValue                TagType                    = 340   // *,     *, # SamplesPerPixel, try double
	TagTypeSMaxSampleValue                TagType                    = 341   // *,     *, # SamplesPerPixel, try double
	TagTypeTransferRange                  TagType                    = 342   // SHORT, 6,
	TagTypeClipPath                       TagType                    = 343   // ingore # Mirrors the essentials of PostScript's path creation functionality.
	TagTypeXClipPathUnits                 TagType                    = 344   // ingore # The number of units that span the width of the image, in terms of integer ClipPath coordinates.
	TagTypeYClipPathUnits                 TagType                    = 345   // ingore # The number of units that span the height of the image, in terms of integer ClipPath coordinates.
	TagTypeIndexed                        TagType                    = 346   // ingore # Aims to broaden the support for indexed images to include support for any color space.
	TagTypeJPEGTables                     TagType                    = 347   // ingore # JPEG quantization and/or Huffman tables.
	TagTypeOPIProxy                       TagType                    = 351   // ingore # OPI-related.
	TagTypeGlobalParametersIFD            TagType                    = 400   // ingore # Used in the TIFF-FX standard to point to an IFD containing tags that are globally applicable to the complete TIFF file.
	TagTypeProfileType                    TagType                    = 401   // ingore # Used in the TIFF-FX standard, denotes the type of data stored in this file or IFD.
	TagTypeFaxProfile                     TagType                    = 402   // ingore # Used in the TIFF-FX standard, denotes the 'profile' that applies to this file.
	TagTypeCodingMethods                  TagType                    = 403   // ingore # Used in the TIFF-FX standard, indicates which coding methods are used in the file.
	TagTypeVersionYear                    TagType                    = 404   // ingore # Used in the TIFF-FX standard, denotes the year of the standard specified by the FaxProfile field.
	TagTypeModeNumber                     TagType                    = 405   // ingore # Used in the TIFF-FX standard, denotes the mode of the standard specified by the FaxProfile field.
	TagTypeDecode                         TagType                    = 433   // ingore # Used in the TIFF-F and TIFF-FX standards, holds information about the ITULAB (PhotometricInterpretation = 10) encoding.
	TagTypeDefaultImageColor              TagType                    = 434   // ingore # Defined in the Mixed Raster Content part of RFC 2301, is the default color needed in areas where no image is available.
	TagTypeJPEGProc                       TagType                    = 512   // SHORT, 1,
	TagTypeJPEGInterchangeFormat          TagType                    = 513   // LONG,  1,
	TagTypeJPEGInterchangeFormatLength    TagType                    = 514   // LONG,  1,
	TagTypeJPEGRestartInterval            TagType                    = 515   // SHORT, 1,
	TagTypeJPEGLosslessPredictors         TagType                    = 517   // SHORT, *, # SamplesPerPixel
	TagTypeJPEGPointTransforms            TagType                    = 518   // SHORT, *, # SamplesPerPixel
	TagTypeJPEGQTables                    TagType                    = 519   // LONG,  *, # SamplesPerPixel
	TagTypeJPEGDCTables                   TagType                    = 520   // LONG,  *, # SamplesPerPixel
	TagTypeJPEGACTables                   TagType                    = 521   // LONG,  *, # SamplesPerPixel
	TagTypeYCbCrCoefficients              TagType                    = 529   // RATIONAL, 3
	TagTypeYCbCrSubSampling               TagType                    = 530   // SHORT, 2, # Default=[2,2]
	TagTypeYCbCrPositioning               TagType                    = 531   // SHORT, 1, # Default=1
	TagTypeReferenceBlackWhite            TagType                    = 532   // LONG , *, # 2*SamplesPerPixel
	TagTypeStripRowCounts                 TagType                    = 559   // ingore # Defined in the Mixed Raster Content part of RFC 2301, used to replace RowsPerStrip for IFDs with variable-sized strips.
	TagTypeXMP                            TagType                    = 700   // ingore # XML packet containing XMP metadata
	TagTypeImageID                        TagType                    = 32781 // ingore # OPI-related.
	TagTypeImageLayer                     TagType                    = 34732 // ingore # Defined in the Mixed Raster Content part of RFC 2301, used to denote the particular function of this Image in the mixed raster scheme.
	TagTypeCopyright                      TagType                    = 33432 // ASCII
	TagTypeWangAnnotation                 TagType                    = 32932 // ingore # Annotation data, as used in 'Imaging for Windows'.
	TagTypeMDFileTag                      TagType                    = 33445 // ingore # Specifies the pixel data format encoding in the Molecular Dynamics GEL file format.
	TagTypeMDScalePixel                   TagType                    = 33446 // ingore # Specifies a scale factor in the Molecular Dynamics GEL file format.
	TagTypeMDColorTable                   TagType                    = 33447 // ingore # Used to specify the conversion from 16bit to 8bit in the Molecular Dynamics GEL file format.
	TagTypeMDLabName                      TagType                    = 33448 // ingore # Name of the lab that scanned this file, as used in the Molecular Dynamics GEL file format.
	TagTypeMDSampleInfo                   TagType                    = 33449 // ingore # Information about the sample, as used in the Molecular Dynamics GEL file format.
	TagTypeMDPrepDate                     TagType                    = 33450 // ingore # Date the sample was prepared, as used in the Molecular Dynamics GEL file format.
	TagTypeMDPrepTime                     TagType                    = 33451 // ingore # Time the sample was prepared, as used in the Molecular Dynamics GEL file format.
	TagTypeMDFileUnits                    TagType                    = 33452 // ingore # Units for data in this file, as used in the Molecular Dynamics GEL file format.
	TagTypeModelPixelScaleTag             TagType                    = 33550 // DOUBLE # Used in interchangeable GeoTIFF files.
	TagTypeIPTC                           TagType                    = 33723 // ingore # IPTC (International Press Telecommunications Council) metadata.
	TagTypeINGRPacketDataTag              TagType                    = 33918 // ingore # Intergraph Application specific storage.
	TagTypeINGRFlagRegisters              TagType                    = 33919 // ingore # Intergraph Application specific flags.
	TagTypeIrasBTransformationMatrix      TagType                    = 33920 // DOUBLE, 17 # Originally part of Intergraph's GeoTIFF tags, but likely understood by IrasB only.
	TagTypeModelTiepointTag               TagType                    = 33922 // DOUBLE # Originally part of Intergraph's GeoTIFF tags, but now used in interchangeable GeoTIFF files.
	TagTypeModelTransformationTag         TagType                    = 34264 // DOUBLE, 16 # Used in interchangeable GeoTIFF files.
	TagTypePhotoshop                      TagType                    = 34377 // ingore # Collection of Photoshop 'Image Resource Blocks'.
	TagTypeExifIFD                        TagType                    = 34665 // IFD    # A pointer to the Exif IFD.
	TagTypeICCProfile                     TagType                    = 34675 // ingore # ICC profile data.
	TagTypeGeoKeyDirectoryTag             TagType                    = 34735 // SHORT, *, # >= 4
	TagTypeGeoDoubleParamsTag             TagType                    = 34736 // DOUBLE
	TagTypeGeoAsciiParamsTag              TagType                    = 34737 // ASCII
	TagTypeGPSIFD                         TagType                    = 34853 // IFD    # A pointer to the Exif-related GPS Info IFD.
	TagTypeHylaFAXFaxRecvParams           TagType                    = 34908 // ingore # Used by HylaFAX.
	TagTypeHylaFAXFaxSubAddress           TagType                    = 34909 // ingore # Used by HylaFAX.
	TagTypeHylaFAXFaxRecvTime             TagType                    = 34910 // ingore # Used by HylaFAX.
	TagTypeImageSourceData                TagType                    = 37724 // ingore # Used by Adobe Photoshop.
	TagTypeInteroperabilityIFD            TagType                    = 40965 // IFD    # A pointer to the Exif-related Interoperability IFD.
	TagTypeGDALMETADATA                   TagType                    = 42112 // ingore # Used by the GDAL library, holds an XML list of name=value 'metadata' values about the image as a whole, and about specific samples.
	TagTypeGDALNODATA                     TagType                    = 42113 // ingore # Used by the GDAL library, contains an ASCII encoded nodata or background pixel value.
	TagTypeOceScanjobDescription          TagType                    = 50215 // ingore # Used in the Oce scanning process.
	TagTypeOceApplicationSelector         TagType                    = 50216 // ingore # Used in the Oce scanning process.
	TagTypeOceIdentificationNumber        TagType                    = 50217 // ingore # Used in the Oce scanning process.
	TagTypeOceImageLogicCharacteristics   TagType                    = 50218 // ingore # Used in the Oce scanning process.
	TagTypeDNGVersion                     TagType                    = 50706 // ingore # Used in IFD 0 of DNG files.
	TagTypeDNGBackwardVersion             TagType                    = 50707 // ingore # Used in IFD 0 of DNG files.
	TagTypeUniqueCameraModel              TagType                    = 50708 // ingore # Used in IFD 0 of DNG files.
	TagTypeLocalizedCameraModel           TagType                    = 50709 // ingore # Used in IFD 0 of DNG files.
	TagTypeCFAPlaneColor                  TagType                    = 50710 // ingore # Used in Raw IFD of DNG files.
	TagTypeCFALayout                      TagType                    = 50711 // ingore # Used in Raw IFD of DNG files.
	TagTypeLinearizationTable             TagType                    = 50712 // ingore # Used in Raw IFD of DNG files.
	TagTypeBlackLevelRepeatDim            TagType                    = 50713 // ingore # Used in Raw IFD of DNG files.
	TagTypeBlackLevel                     TagType                    = 50714 // ingore # Used in Raw IFD of DNG files.
	TagTypeBlackLevelDeltaH               TagType                    = 50715 // ingore # Used in Raw IFD of DNG files.
	TagTypeBlackLevelDeltaV               TagType                    = 50716 // ingore # Used in Raw IFD of DNG files.
	TagTypeWhiteLevel                     TagType                    = 50717 // ingore # Used in Raw IFD of DNG files.
	TagTypeDefaultScale                   TagType                    = 50718 // ingore # Used in Raw IFD of DNG files.
	TagTypeDefaultCropOrigin              TagType                    = 50719 // ingore # Used in Raw IFD of DNG files.
	TagTypeDefaultCropSize                TagType                    = 50720 // ingore # Used in Raw IFD of DNG files.
	TagTypeColorMatrix1                   TagType                    = 50721 // ingore # Used in IFD 0 of DNG files.
	TagTypeColorMatrix2                   TagType                    = 50722 // ingore # Used in IFD 0 of DNG files.
	TagTypeCameraCalibration1             TagType                    = 50723 // ingore # Used in IFD 0 of DNG files.
	TagTypeCameraCalibration2             TagType                    = 50724 // ingore # Used in IFD 0 of DNG files.
	TagTypeReductionMatrix1               TagType                    = 50725 // ingore # Used in IFD 0 of DNG files.
	TagTypeReductionMatrix2               TagType                    = 50726 // ingore # Used in IFD 0 of DNG files.
	TagTypeAnalogBalance                  TagType                    = 50727 // ingore # Used in IFD 0 of DNG files.
	TagTypeAsShotNeutral                  TagType                    = 50728 // ingore # Used in IFD 0 of DNG files.
	TagTypeAsShotWhiteXY                  TagType                    = 50729 // ingore # Used in IFD 0 of DNG files.
	TagTypeBaselineExposure               TagType                    = 50730 // ingore # Used in IFD 0 of DNG files.
	TagTypeBaselineNoise                  TagType                    = 50731 // ingore # Used in IFD 0 of DNG files.
	TagTypeBaselineSharpness              TagType                    = 50732 // ingore # Used in IFD 0 of DNG files.
	TagTypeBayerGreenSplit                TagType                    = 50733 // ingore # Used in Raw IFD of DNG files.
	TagTypeLinearResponseLimit            TagType                    = 50734 // ingore # Used in IFD 0 of DNG files.
	TagTypeCameraSerialNumber             TagType                    = 50735 // ingore # Used in IFD 0 of DNG files.
	TagTypeLensInfo                       TagType                    = 50736 // ingore # Used in IFD 0 of DNG files.
	TagTypeChromaBlurRadius               TagType                    = 50737 // ingore # Used in Raw IFD of DNG files.
	TagTypeAntiAliasStrength              TagType                    = 50738 // ingore # Used in Raw IFD of DNG files.
	TagTypeDNGPrivateData                 TagType                    = 50740 // ingore # Used in IFD 0 of DNG files.
	TagTypeMakerNoteSafety                TagType                    = 50741 // ingore # Used in IFD 0 of DNG files.
	TagTypeCalibrationIlluminant1         TagType                    = 50778 // ingore # Used in IFD 0 of DNG files.
	TagTypeCalibrationIlluminant2         TagType                    = 50779 // ingore # Used in IFD 0 of DNG files.
	TagTypeBestQualityScale               TagType                    = 50780 // ingore # Used in Raw IFD of DNG files.
	TagTypeAliasLayerMetadata             TagType                    = 50784 // ingore # Alias Sketchbook Pro layer usage description.
)

/*****************************************************************************************************************/
