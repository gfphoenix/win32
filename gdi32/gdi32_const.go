// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package gdi32

// GetDeviceCaps index constants
const (
	DRIVERVERSION   = 0
	TECHNOLOGY      = 2
	HORZSIZE        = 4
	VERTSIZE        = 6
	HORZRES         = 8
	VERTRES         = 10
	LOGPIXELSX      = 88
	LOGPIXELSY      = 90
	BITSPIXEL       = 12
	PLANES          = 14
	NUMBRUSHES      = 16
	NUMPENS         = 18
	NUMFONTS        = 22
	NUMCOLORS       = 24
	NUMMARKERS      = 20
	ASPECTX         = 40
	ASPECTY         = 42
	ASPECTXY        = 44
	PDEVICESIZE     = 26
	CLIPCAPS        = 36
	SIZEPALETTE     = 104
	NUMRESERVED     = 106
	COLORRES        = 108
	PHYSICALWIDTH   = 110
	PHYSICALHEIGHT  = 111
	PHYSICALOFFSETX = 112
	PHYSICALOFFSETY = 113
	SCALINGFACTORX  = 114
	SCALINGFACTORY  = 115
	VREFRESH        = 116
	DESKTOPHORZRES  = 118
	DESKTOPVERTRES  = 117
	BLTALIGNMENT    = 119
	SHADEBLENDCAPS  = 120
	COLORMGMTCAPS   = 121
	RASTERCAPS      = 38
	CURVECAPS       = 28
	LINECAPS        = 30
	POLYGONALCAPS   = 32
	TEXTCAPS        = 34
)

// GetDeviceCaps TECHNOLOGY constants
const (
	DT_PLOTTER    = 0
	DT_RASDISPLAY = 1
	DT_RASPRINTER = 2
	DT_RASCAMERA  = 3
	DT_CHARSTREAM = 4
	DT_METAFILE   = 5
	DT_DISPFILE   = 6
)

// GetDeviceCaps SHADEBLENDCAPS constants
const (
	SB_NONE          = 0x00
	SB_CONST_ALPHA   = 0x01
	SB_PIXEL_ALPHA   = 0x02
	SB_PREMULT_ALPHA = 0x04
	SB_GRAD_RECT     = 0x10
	SB_GRAD_TRI      = 0x20
)

// GetDeviceCaps COLORMGMTCAPS constants
const (
	CM_NONE       = 0x00
	CM_DEVICE_ICM = 0x01
	CM_GAMMA_RAMP = 0x02
	CM_CMYK_COLOR = 0x04
)

// GetDeviceCaps RASTERCAPS constants
const (
	RC_BANDING      = 2
	RC_BITBLT       = 1
	RC_BITMAP64     = 8
	RC_DI_BITMAP    = 128
	RC_DIBTODEV     = 512
	RC_FLOODFILL    = 4096
	RC_GDI20_OUTPUT = 16
	RC_PALETTE      = 256
	RC_SCALING      = 4
	RC_STRETCHBLT   = 2048
	RC_STRETCHDIB   = 8192
	RC_DEVBITS      = 0x8000
	RC_OP_DX_OUTPUT = 0x4000
)

// GetDeviceCaps CURVECAPS constants
const (
	CC_NONE       = 0
	CC_CIRCLES    = 1
	CC_PIE        = 2
	CC_CHORD      = 4
	CC_ELLIPSES   = 8
	CC_WIDE       = 16
	CC_STYLED     = 32
	CC_WIDESTYLED = 64
	CC_INTERIORS  = 128
	CC_ROUNDRECT  = 256
)

// GetDeviceCaps LINECAPS constants
const (
	LC_NONE       = 0
	LC_POLYLINE   = 2
	LC_MARKER     = 4
	LC_POLYMARKER = 8
	LC_WIDE       = 16
	LC_STYLED     = 32
	LC_WIDESTYLED = 64
	LC_INTERIORS  = 128
)

// GetDeviceCaps POLYGONALCAPS constants
const (
	PC_NONE        = 0
	PC_POLYGON     = 1
	PC_POLYPOLYGON = 256
	PC_PATHS       = 512
	PC_RECTANGLE   = 2
	PC_WINDPOLYGON = 4
	PC_SCANLINE    = 8
	PC_TRAPEZOID   = 4
	PC_WIDE        = 16
	PC_STYLED      = 32
	PC_WIDESTYLED  = 64
	PC_INTERIORS   = 128
)

// GetDeviceCaps TEXTCAPS constants
const (
	TC_OP_CHARACTER = 1
	TC_OP_STROKE    = 2
	TC_CP_STROKE    = 4
	TC_CR_90        = 8
	TC_CR_ANY       = 16
	TC_SF_X_YINDEP  = 32
	TC_SA_DOUBLE    = 64
	TC_SA_INTEGER   = 128
	TC_SA_CONTIN    = 256
	TC_EA_DOUBLE    = 512
	TC_IA_ABLE      = 1024
	TC_UA_ABLE      = 2048
	TC_SO_ABLE      = 4096
	TC_RA_ABLE      = 8192
	TC_VA_ABLE      = 16384
	TC_RESERVED     = 32768
	TC_SCROLLBLT    = 65536
)

// Brush styles
const (
	BS_SOLID         = 0
	BS_NULL          = 1
	BS_HOLLOW        = BS_NULL
	BS_HATCHED       = 2
	BS_PATTERN       = 3
	BS_INDEXED       = 4
	BS_DIBPATTERN    = 5
	BS_DIBPATTERNPT  = 6
	BS_PATTERN8X8    = 7
	BS_DIBPATTERN8X8 = 8
	BS_MONOPATTERN   = 9
)

// Hatch styles
const (
	HS_HORIZONTAL = 0
	HS_VERTICAL   = 1
	HS_FDIAGONAL  = 2
	HS_BDIAGONAL  = 3
	HS_CROSS      = 4
	HS_DIAGCROSS  = 5
)

// Pen types
const (
	PS_COSMETIC  = 0x00000000
	PS_GEOMETRIC = 0x00010000
	PS_TYPE_MASK = 0x000F0000
)

// Pen styles
const (
	PS_SOLID       = 0
	PS_DASH        = 1
	PS_DOT         = 2
	PS_DASHDOT     = 3
	PS_DASHDOTDOT  = 4
	PS_NULL        = 5
	PS_INSIDEFRAME = 6
	PS_USERSTYLE   = 7
	PS_ALTERNATE   = 8
	PS_STYLE_MASK  = 0x0000000F
)

// Pen cap types
const (
	PS_ENDCAP_ROUND  = 0x00000000
	PS_ENDCAP_SQUARE = 0x00000100
	PS_ENDCAP_FLAT   = 0x00000200
	PS_ENDCAP_MASK   = 0x00000F00
)

// Pen join types
const (
	PS_JOIN_ROUND = 0x00000000
	PS_JOIN_BEVEL = 0x00001000
	PS_JOIN_MITER = 0x00002000
	PS_JOIN_MASK  = 0x0000F000
)

// Print constants
const (
	PRF_NONCLIENT  = 0x00000002
	PRF_CLIENT     = 0x00000004
	PRF_ERASEBKGND = 0x00000008
	PRF_CHILDREN   = 0x00000010
	PRF_OWNED      = 0x00000020
)

// Stock logical objects
const (
	WHITE_BRUSH         = 0
	LTGRAY_BRUSH        = 1
	GRAY_BRUSH          = 2
	DKGRAY_BRUSH        = 3
	BLACK_BRUSH         = 4
	NULL_BRUSH          = 5
	HOLLOW_BRUSH        = NULL_BRUSH
	WHITE_PEN           = 6
	BLACK_PEN           = 7
	NULL_PEN            = 8
	OEM_FIXED_FONT      = 10
	ANSI_FIXED_FONT     = 11
	ANSI_VAR_FONT       = 12
	SYSTEM_FONT         = 13
	DEVICE_DEFAULT_FONT = 14
	DEFAULT_PALETTE     = 15
	SYSTEM_FIXED_FONT   = 16
	DEFAULT_GUI_FONT    = 17
	DC_BRUSH            = 18
	DC_PEN              = 19
)

const LF_FACESIZE = 32

// Font weight constants
const (
	FW_DONTCARE   = 0
	FW_THIN       = 100
	FW_EXTRALIGHT = 200
	FW_ULTRALIGHT = FW_EXTRALIGHT
	FW_LIGHT      = 300
	FW_NORMAL     = 400
	FW_REGULAR    = 400
	FW_MEDIUM     = 500
	FW_SEMIBOLD   = 600
	FW_DEMIBOLD   = FW_SEMIBOLD
	FW_BOLD       = 700
	FW_EXTRABOLD  = 800
	FW_ULTRABOLD  = FW_EXTRABOLD
	FW_HEAVY      = 900
	FW_BLACK      = FW_HEAVY
)

// Charset constants
const (
	ANSI_CHARSET        = 0
	DEFAULT_CHARSET     = 1
	SYMBOL_CHARSET      = 2
	SHIFTJIS_CHARSET    = 128
	HANGEUL_CHARSET     = 129
	HANGUL_CHARSET      = 129
	GB2312_CHARSET      = 134
	CHINESEBIG5_CHARSET = 136
	GREEK_CHARSET       = 161
	TURKISH_CHARSET     = 162
	HEBREW_CHARSET      = 177
	ARABIC_CHARSET      = 178
	BALTIC_CHARSET      = 186
	RUSSIAN_CHARSET     = 204
	THAI_CHARSET        = 222
	EASTEUROPE_CHARSET  = 238
	OEM_CHARSET         = 255
	JOHAB_CHARSET       = 130
	VIETNAMESE_CHARSET  = 163
	MAC_CHARSET         = 77
)

// Font output precision constants
const (
	OUT_DEFAULT_PRECIS   = 0
	OUT_STRING_PRECIS    = 1
	OUT_CHARACTER_PRECIS = 2
	OUT_STROKE_PRECIS    = 3
	OUT_TT_PRECIS        = 4
	OUT_DEVICE_PRECIS    = 5
	OUT_RASTER_PRECIS    = 6
	OUT_TT_ONLY_PRECIS   = 7
	OUT_OUTLINE_PRECIS   = 8
	OUT_PS_ONLY_PRECIS   = 10
)

// Font clipping precision constants
const (
	CLIP_DEFAULT_PRECIS   = 0
	CLIP_CHARACTER_PRECIS = 1
	CLIP_STROKE_PRECIS    = 2
	CLIP_MASK             = 15
	CLIP_LH_ANGLES        = 16
	CLIP_TT_ALWAYS        = 32
	CLIP_EMBEDDED         = 128
)

// Font output quality constants
const (
	DEFAULT_QUALITY        = 0
	DRAFT_QUALITY          = 1
	PROOF_QUALITY          = 2
	NONANTIALIASED_QUALITY = 3
	ANTIALIASED_QUALITY    = 4
	CLEARTYPE_QUALITY      = 5
)

// Font pitch constants
const (
	DEFAULT_PITCH  = 0
	FIXED_PITCH    = 1
	VARIABLE_PITCH = 2
)

// Font family constants
const (
	FF_DECORATIVE = 80
	FF_DONTCARE   = 0
	FF_MODERN     = 48
	FF_ROMAN      = 16
	FF_SCRIPT     = 64
	FF_SWISS      = 32
)

// DeviceCapabilities capabilities
const (
	DC_FIELDS            = 1
	DC_PAPERS            = 2
	DC_PAPERSIZE         = 3
	DC_MINEXTENT         = 4
	DC_MAXEXTENT         = 5
	DC_BINS              = 6
	DC_DUPLEX            = 7
	DC_SIZE              = 8
	DC_EXTRA             = 9
	DC_VERSION           = 10
	DC_DRIVER            = 11
	DC_BINNAMES          = 12
	DC_ENUMRESOLUTIONS   = 13
	DC_FILEDEPENDENCIES  = 14
	DC_TRUETYPE          = 15
	DC_PAPERNAMES        = 16
	DC_ORIENTATION       = 17
	DC_COPIES            = 18
	DC_BINADJUST         = 19
	DC_EMF_COMPLIANT     = 20
	DC_DATATYPE_PRODUCED = 21
	DC_COLLATE           = 22
	DC_MANUFACTURER      = 23
	DC_MODEL             = 24
	DC_PERSONALITY       = 25
	DC_PRINTRATE         = 26
	DC_PRINTRATEUNIT     = 27
	DC_PRINTERMEM        = 28
	DC_MEDIAREADY        = 29
	DC_STAPLE            = 30
	DC_PRINTRATEPPM      = 31
	DC_COLORDEVICE       = 32
	DC_NUP               = 33
	DC_MEDIATYPENAMES    = 34
	DC_MEDIATYPES        = 35
)

const (
	CCHDEVICENAME = 32
	CCHFORMNAME   = 32
)

const (
	DM_UPDATE      = 1
	DM_COPY        = 2
	DM_PROMPT      = 4
	DM_MODIFY      = 8
	DM_IN_BUFFER   = DM_MODIFY
	DM_IN_PROMPT   = DM_PROMPT
	DM_OUT_BUFFER  = DM_COPY
	DM_OUT_DEFAULT = DM_UPDATE
)

// DEVMODE field selection bits
const (
	DM_ORIENTATION        = 0x00000001
	DM_PAPERSIZE          = 0x00000002
	DM_PAPERLENGTH        = 0x00000004
	DM_PAPERWIDTH         = 0x00000008
	DM_SCALE              = 0x00000010
	DM_POSITION           = 0x00000020
	DM_NUP                = 0x00000040
	DM_DISPLAYORIENTATION = 0x00000080
	DM_COPIES             = 0x00000100
	DM_DEFAULTSOURCE      = 0x00000200
	DM_PRINTQUALITY       = 0x00000400
	DM_COLOR              = 0x00000800
	DM_DUPLEX             = 0x00001000
	DM_YRESOLUTION        = 0x00002000
	DM_TTOPTION           = 0x00004000
	DM_COLLATE            = 0x00008000
	DM_FORMNAME           = 0x00010000
	DM_LOGPIXELS          = 0x00020000
	DM_BITSPERPEL         = 0x00040000
	DM_PELSWIDTH          = 0x00080000
	DM_PELSHEIGHT         = 0x00100000
	DM_DISPLAYFLAGS       = 0x00200000
	DM_DISPLAYFREQUENCY   = 0x00400000
	DM_ICMMETHOD          = 0x00800000
	DM_ICMINTENT          = 0x01000000
	DM_MEDIATYPE          = 0x02000000
	DM_DITHERTYPE         = 0x04000000
	DM_PANNINGWIDTH       = 0x08000000
	DM_PANNINGHEIGHT      = 0x10000000
	DM_DISPLAYFIXEDOUTPUT = 0x20000000
)

// Orientation constants
const (
	DMORIENT_PORTRAIT  = 1
	DMORIENT_LANDSCAPE = 2
)

// Paper sizes
const (
	DMPAPER_FIRST                         = DMPAPER_LETTER
	DMPAPER_LETTER                        = 1   /* Letter 8 1/2 x 11 in               */
	DMPAPER_LETTERSMALL                   = 2   /* Letter Small 8 1/2 x 11 in         */
	DMPAPER_TABLOID                       = 3   /* Tabloid 11 x 17 in                 */
	DMPAPER_LEDGER                        = 4   /* Ledger 17 x 11 in                  */
	DMPAPER_LEGAL                         = 5   /* Legal 8 1/2 x 14 in                */
	DMPAPER_STATEMENT                     = 6   /* Statement 5 1/2 x 8 1/2 in         */
	DMPAPER_EXECUTIVE                     = 7   /* Executive 7 1/4 x 10 1/2 in        */
	DMPAPER_A3                            = 8   /* A3 297 x 420 mm                    */
	DMPAPER_A4                            = 9   /* A4 210 x 297 mm                    */
	DMPAPER_A4SMALL                       = 10  /* A4 Small 210 x 297 mm              */
	DMPAPER_A5                            = 11  /* A5 148 x 210 mm                    */
	DMPAPER_B4                            = 12  /* B4 (JIS) 250 x 354                 */
	DMPAPER_B5                            = 13  /* B5 (JIS) 182 x 257 mm              */
	DMPAPER_FOLIO                         = 14  /* Folio 8 1/2 x 13 in                */
	DMPAPER_QUARTO                        = 15  /* Quarto 215 x 275 mm                */
	DMPAPER_10X14                         = 16  /* 10x14 in                           */
	DMPAPER_11X17                         = 17  /* 11x17 in                           */
	DMPAPER_NOTE                          = 18  /* Note 8 1/2 x 11 in                 */
	DMPAPER_ENV_9                         = 19  /* Envelope #9 3 7/8 x 8 7/8          */
	DMPAPER_ENV_10                        = 20  /* Envelope #10 4 1/8 x 9 1/2         */
	DMPAPER_ENV_11                        = 21  /* Envelope #11 4 1/2 x 10 3/8        */
	DMPAPER_ENV_12                        = 22  /* Envelope #12 4 \276 x 11           */
	DMPAPER_ENV_14                        = 23  /* Envelope #14 5 x 11 1/2            */
	DMPAPER_CSHEET                        = 24  /* C size sheet                       */
	DMPAPER_DSHEET                        = 25  /* D size sheet                       */
	DMPAPER_ESHEET                        = 26  /* E size sheet                       */
	DMPAPER_ENV_DL                        = 27  /* Envelope DL 110 x 220mm            */
	DMPAPER_ENV_C5                        = 28  /* Envelope C5 162 x 229 mm           */
	DMPAPER_ENV_C3                        = 29  /* Envelope C3  324 x 458 mm          */
	DMPAPER_ENV_C4                        = 30  /* Envelope C4  229 x 324 mm          */
	DMPAPER_ENV_C6                        = 31  /* Envelope C6  114 x 162 mm          */
	DMPAPER_ENV_C65                       = 32  /* Envelope C65 114 x 229 mm          */
	DMPAPER_ENV_B4                        = 33  /* Envelope B4  250 x 353 mm          */
	DMPAPER_ENV_B5                        = 34  /* Envelope B5  176 x 250 mm          */
	DMPAPER_ENV_B6                        = 35  /* Envelope B6  176 x 125 mm          */
	DMPAPER_ENV_ITALY                     = 36  /* Envelope 110 x 230 mm              */
	DMPAPER_ENV_MONARCH                   = 37  /* Envelope Monarch 3.875 x 7.5 in    */
	DMPAPER_ENV_PERSONAL                  = 38  /* 6 3/4 Envelope 3 5/8 x 6 1/2 in    */
	DMPAPER_FANFOLD_US                    = 39  /* US Std Fanfold 14 7/8 x 11 in      */
	DMPAPER_FANFOLD_STD_GERMAN            = 40  /* German Std Fanfold 8 1/2 x 12 in   */
	DMPAPER_FANFOLD_LGL_GERMAN            = 41  /* German Legal Fanfold 8 1/2 x 13 in */
	DMPAPER_ISO_B4                        = 42  /* B4 (ISO) 250 x 353 mm              */
	DMPAPER_JAPANESE_POSTCARD             = 43  /* Japanese Postcard 100 x 148 mm     */
	DMPAPER_9X11                          = 44  /* 9 x 11 in                          */
	DMPAPER_10X11                         = 45  /* 10 x 11 in                         */
	DMPAPER_15X11                         = 46  /* 15 x 11 in                         */
	DMPAPER_ENV_INVITE                    = 47  /* Envelope Invite 220 x 220 mm       */
	DMPAPER_RESERVED_48                   = 48  /* RESERVED--DO NOT USE               */
	DMPAPER_RESERVED_49                   = 49  /* RESERVED--DO NOT USE               */
	DMPAPER_LETTER_EXTRA                  = 50  /* Letter Extra 9 \275 x 12 in        */
	DMPAPER_LEGAL_EXTRA                   = 51  /* Legal Extra 9 \275 x 15 in         */
	DMPAPER_TABLOID_EXTRA                 = 52  /* Tabloid Extra 11.69 x 18 in        */
	DMPAPER_A4_EXTRA                      = 53  /* A4 Extra 9.27 x 12.69 in           */
	DMPAPER_LETTER_TRANSVERSE             = 54  /* Letter Transverse 8 \275 x 11 in   */
	DMPAPER_A4_TRANSVERSE                 = 55  /* A4 Transverse 210 x 297 mm         */
	DMPAPER_LETTER_EXTRA_TRANSVERSE       = 56  /* Letter Extra Transverse 9\275 x 12 in */
	DMPAPER_A_PLUS                        = 57  /* SuperA/SuperA/A4 227 x 356 mm      */
	DMPAPER_B_PLUS                        = 58  /* SuperB/SuperB/A3 305 x 487 mm      */
	DMPAPER_LETTER_PLUS                   = 59  /* Letter Plus 8.5 x 12.69 in         */
	DMPAPER_A4_PLUS                       = 60  /* A4 Plus 210 x 330 mm               */
	DMPAPER_A5_TRANSVERSE                 = 61  /* A5 Transverse 148 x 210 mm         */
	DMPAPER_B5_TRANSVERSE                 = 62  /* B5 (JIS) Transverse 182 x 257 mm   */
	DMPAPER_A3_EXTRA                      = 63  /* A3 Extra 322 x 445 mm              */
	DMPAPER_A5_EXTRA                      = 64  /* A5 Extra 174 x 235 mm              */
	DMPAPER_B5_EXTRA                      = 65  /* B5 (ISO) Extra 201 x 276 mm        */
	DMPAPER_A2                            = 66  /* A2 420 x 594 mm                    */
	DMPAPER_A3_TRANSVERSE                 = 67  /* A3 Transverse 297 x 420 mm         */
	DMPAPER_A3_EXTRA_TRANSVERSE           = 68  /* A3 Extra Transverse 322 x 445 mm   */
	DMPAPER_DBL_JAPANESE_POSTCARD         = 69  /* Japanese Double Postcard 200 x 148 mm */
	DMPAPER_A6                            = 70  /* A6 105 x 148 mm                 */
	DMPAPER_JENV_KAKU2                    = 71  /* Japanese Envelope Kaku #2       */
	DMPAPER_JENV_KAKU3                    = 72  /* Japanese Envelope Kaku #3       */
	DMPAPER_JENV_CHOU3                    = 73  /* Japanese Envelope Chou #3       */
	DMPAPER_JENV_CHOU4                    = 74  /* Japanese Envelope Chou #4       */
	DMPAPER_LETTER_ROTATED                = 75  /* Letter Rotated 11 x 8 1/2 11 in */
	DMPAPER_A3_ROTATED                    = 76  /* A3 Rotated 420 x 297 mm         */
	DMPAPER_A4_ROTATED                    = 77  /* A4 Rotated 297 x 210 mm         */
	DMPAPER_A5_ROTATED                    = 78  /* A5 Rotated 210 x 148 mm         */
	DMPAPER_B4_JIS_ROTATED                = 79  /* B4 (JIS) Rotated 364 x 257 mm   */
	DMPAPER_B5_JIS_ROTATED                = 80  /* B5 (JIS) Rotated 257 x 182 mm   */
	DMPAPER_JAPANESE_POSTCARD_ROTATED     = 81  /* Japanese Postcard Rotated 148 x 100 mm */
	DMPAPER_DBL_JAPANESE_POSTCARD_ROTATED = 82  /* Double Japanese Postcard Rotated 148 x 200 mm */
	DMPAPER_A6_ROTATED                    = 83  /* A6 Rotated 148 x 105 mm         */
	DMPAPER_JENV_KAKU2_ROTATED            = 84  /* Japanese Envelope Kaku #2 Rotated */
	DMPAPER_JENV_KAKU3_ROTATED            = 85  /* Japanese Envelope Kaku #3 Rotated */
	DMPAPER_JENV_CHOU3_ROTATED            = 86  /* Japanese Envelope Chou #3 Rotated */
	DMPAPER_JENV_CHOU4_ROTATED            = 87  /* Japanese Envelope Chou #4 Rotated */
	DMPAPER_B6_JIS                        = 88  /* B6 (JIS) 128 x 182 mm           */
	DMPAPER_B6_JIS_ROTATED                = 89  /* B6 (JIS) Rotated 182 x 128 mm   */
	DMPAPER_12X11                         = 90  /* 12 x 11 in                      */
	DMPAPER_JENV_YOU4                     = 91  /* Japanese Envelope You #4        */
	DMPAPER_JENV_YOU4_ROTATED             = 92  /* Japanese Envelope You #4 Rotated*/
	DMPAPER_P16K                          = 93  /* PRC 16K 146 x 215 mm            */
	DMPAPER_P32K                          = 94  /* PRC 32K 97 x 151 mm             */
	DMPAPER_P32KBIG                       = 95  /* PRC 32K(Big) 97 x 151 mm        */
	DMPAPER_PENV_1                        = 96  /* PRC Envelope #1 102 x 165 mm    */
	DMPAPER_PENV_2                        = 97  /* PRC Envelope #2 102 x 176 mm    */
	DMPAPER_PENV_3                        = 98  /* PRC Envelope #3 125 x 176 mm    */
	DMPAPER_PENV_4                        = 99  /* PRC Envelope #4 110 x 208 mm    */
	DMPAPER_PENV_5                        = 100 /* PRC Envelope #5 110 x 220 mm    */
	DMPAPER_PENV_6                        = 101 /* PRC Envelope #6 120 x 230 mm    */
	DMPAPER_PENV_7                        = 102 /* PRC Envelope #7 160 x 230 mm    */
	DMPAPER_PENV_8                        = 103 /* PRC Envelope #8 120 x 309 mm    */
	DMPAPER_PENV_9                        = 104 /* PRC Envelope #9 229 x 324 mm    */
	DMPAPER_PENV_10                       = 105 /* PRC Envelope #10 324 x 458 mm   */
	DMPAPER_P16K_ROTATED                  = 106 /* PRC 16K Rotated                 */
	DMPAPER_P32K_ROTATED                  = 107 /* PRC 32K Rotated                 */
	DMPAPER_P32KBIG_ROTATED               = 108 /* PRC 32K(Big) Rotated            */
	DMPAPER_PENV_1_ROTATED                = 109 /* PRC Envelope #1 Rotated 165 x 102 mm */
	DMPAPER_PENV_2_ROTATED                = 110 /* PRC Envelope #2 Rotated 176 x 102 mm */
	DMPAPER_PENV_3_ROTATED                = 111 /* PRC Envelope #3 Rotated 176 x 125 mm */
	DMPAPER_PENV_4_ROTATED                = 112 /* PRC Envelope #4 Rotated 208 x 110 mm */
	DMPAPER_PENV_5_ROTATED                = 113 /* PRC Envelope #5 Rotated 220 x 110 mm */
	DMPAPER_PENV_6_ROTATED                = 114 /* PRC Envelope #6 Rotated 230 x 120 mm */
	DMPAPER_PENV_7_ROTATED                = 115 /* PRC Envelope #7 Rotated 230 x 160 mm */
	DMPAPER_PENV_8_ROTATED                = 116 /* PRC Envelope #8 Rotated 309 x 120 mm */
	DMPAPER_PENV_9_ROTATED                = 117 /* PRC Envelope #9 Rotated 324 x 229 mm */
	DMPAPER_PENV_10_ROTATED               = 118 /* PRC Envelope #10 Rotated 458 x 324 mm */
	DMPAPER_LAST                          = DMPAPER_PENV_10_ROTATED
	DMPAPER_USER                          = 256
)

// Bin constants
const (
	DMBIN_FIRST         = DMBIN_UPPER
	DMBIN_UPPER         = 1
	DMBIN_ONLYONE       = 1
	DMBIN_LOWER         = 2
	DMBIN_MIDDLE        = 3
	DMBIN_MANUAL        = 4
	DMBIN_ENVELOPE      = 5
	DMBIN_ENVMANUAL     = 6
	DMBIN_AUTO          = 7
	DMBIN_TRACTOR       = 8
	DMBIN_SMALLFMT      = 9
	DMBIN_LARGEFMT      = 10
	DMBIN_LARGECAPACITY = 11
	DMBIN_CASSETTE      = 14
	DMBIN_FORMSOURCE    = 15
	DMBIN_LAST          = DMBIN_FORMSOURCE
	DMBIN_USER          = 256
)

// Quality constants
const (
	DMRES_DRAFT  = -1
	DMRES_LOW    = -2
	DMRES_MEDIUM = -3
	DMRES_HIGH   = -4
)

// Color/monochrome constants
const (
	DMCOLOR_MONOCHROME = 1
	DMCOLOR_COLOR      = 2
)

// Duplex constants
const (
	DMDUP_SIMPLEX    = 1
	DMDUP_VERTICAL   = 2
	DMDUP_HORIZONTAL = 3
)

// TrueType constants
const (
	DMTT_BITMAP           = 1
	DMTT_DOWNLOAD         = 2
	DMTT_SUBDEV           = 3
	DMTT_DOWNLOAD_OUTLINE = 4
)

// Collation constants
const (
	DMCOLLATE_FALSE = 0
	DMCOLLATE_TRUE  = 1
)

// Background modes
const (
	TRANSPARENT = 1
	OPAQUE      = 2
)

// Ternary raster operations
const (
	SRCCOPY        = 0x00CC0020
	SRCPAINT       = 0x00EE0086
	SRCAND         = 0x008800C6
	SRCINVERT      = 0x00660046
	SRCERASE       = 0x00440328
	NOTSRCCOPY     = 0x00330008
	NOTSRCERASE    = 0x001100A6
	MERGECOPY      = 0x00C000CA
	MERGEPAINT     = 0x00BB0226
	PATCOPY        = 0x00F00021
	PATPAINT       = 0x00FB0A09
	PATINVERT      = 0x005A0049
	DSTINVERT      = 0x00550009
	BLACKNESS      = 0x00000042
	WHITENESS      = 0x00FF0062
	NOMIRRORBITMAP = 0x80000000
	CAPTUREBLT     = 0x40000000
)

// StretchBlt modes
const (
	BLACKONWHITE        = 1
	WHITEONBLACK        = 2
	COLORONCOLOR        = 3
	HALFTONE            = 4
	MAXSTRETCHBLTMODE   = 4
	STRETCH_ANDSCANS    = BLACKONWHITE
	STRETCH_ORSCANS     = WHITEONBLACK
	STRETCH_DELETESCANS = COLORONCOLOR
	STRETCH_HALFTONE    = HALFTONE
)

// Bitmap compression constants
const (
	BI_RGB       = 0
	BI_RLE8      = 1
	BI_RLE4      = 2
	BI_BITFIELDS = 3
	BI_JPEG      = 4
	BI_PNG       = 5
)

// Bitmap color table usage
const (
	DIB_RGB_COLORS = 0
	DIB_PAL_COLORS = 1
)

const CBM_INIT = 4

const (
	CLR_INVALID = 0xFFFFFFFF
	CLR_NONE    = CLR_INVALID
	CLR_DEFAULT = 0xFF000000
)

const (
	/* pixel types */
	PFD_TYPE_RGBA       = 0
	PFD_TYPE_COLORINDEX = 1

	/* layer types */
	PFD_MAIN_PLANE     = 0
	PFD_OVERLAY_PLANE  = 1
	PFD_UNDERLAY_PLANE = (-1)

	/* PIXELFORMATDESCRIPTOR flags */
	PFD_DOUBLEBUFFER        = 0x00000001
	PFD_STEREO              = 0x00000002
	PFD_DRAW_TO_WINDOW      = 0x00000004
	PFD_DRAW_TO_BITMAP      = 0x00000008
	PFD_SUPPORT_GDI         = 0x00000010
	PFD_SUPPORT_OPENGL      = 0x00000020
	PFD_GENERIC_FORMAT      = 0x00000040
	PFD_NEED_PALETTE        = 0x00000080
	PFD_NEED_SYSTEM_PALETTE = 0x00000100
	PFD_SWAP_EXCHANGE       = 0x00000200
	PFD_SWAP_COPY           = 0x00000400
	PFD_SWAP_LAYER_BUFFERS  = 0x00000800
	PFD_GENERIC_ACCELERATED = 0x00001000
	PFD_SUPPORT_DIRECTDRAW  = 0x00002000

	/* PIXELFORMATDESCRIPTOR flags for use in ChoosePixelFormat only */
	PFD_DEPTH_DONTCARE        = 0x20000000
	PFD_DOUBLEBUFFER_DONTCARE = 0x40000000
	PFD_STEREO_DONTCARE       = 0x80000000
)

// GradientFill constants
const (
	GRADIENT_FILL_RECT_H   = 0x00
	GRADIENT_FILL_RECT_V   = 0x01
	GRADIENT_FILL_TRIANGLE = 0x02
)

// Region Combine Modes
const (
	RGN_AND  = 1
	RGN_OR   = 2
	RGN_XOR  = 3
	RGN_DIFF = 4
	RGN_COPY = 5
)

// Region Types
const (
	REGIONERROR   = 0
	NULLREGION    = 1
	SIMPLEREGION  = 2
	COMPLEXREGION = 3
)

// AlphaBlend operations
const (
	AC_SRC_ALPHA = 0x1
)

// AddFontResourceEx flags
const (
	FR_PRIVATE  = 0x10
	FR_NOT_ENUM = 0x20
)
