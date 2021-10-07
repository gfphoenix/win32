// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import (
	"github.com/Gipcomp/win32/comctl32"
	"github.com/Gipcomp/win32/user32"
)

const (
	TB_THUMBPOSITION         = 4
	TB_THUMBTRACK            = 5
	TB_ENDTRACK              = 8
	TB_ENABLEBUTTON          = user32.WM_USER + 1
	TB_CHECKBUTTON           = user32.WM_USER + 2
	TB_PRESSBUTTON           = user32.WM_USER + 3
	TB_HIDEBUTTON            = user32.WM_USER + 4
	TB_INDETERMINATE         = user32.WM_USER + 5
	TB_MARKBUTTON            = user32.WM_USER + 6
	TB_ISBUTTONENABLED       = user32.WM_USER + 9
	TB_ISBUTTONCHECKED       = user32.WM_USER + 10
	TB_ISBUTTONPRESSED       = user32.WM_USER + 11
	TB_ISBUTTONHIDDEN        = user32.WM_USER + 12
	TB_ISBUTTONINDETERMINATE = user32.WM_USER + 13
	TB_ISBUTTONHIGHLIGHTED   = user32.WM_USER + 14
	TB_SETSTATE              = user32.WM_USER + 17
	TB_GETSTATE              = user32.WM_USER + 18
	TB_ADDBITMAP             = user32.WM_USER + 19
	TB_DELETEBUTTON          = user32.WM_USER + 22
	TB_GETBUTTON             = user32.WM_USER + 23
	TB_BUTTONCOUNT           = user32.WM_USER + 24
	TB_COMMANDTOINDEX        = user32.WM_USER + 25
	TB_SAVERESTORE           = user32.WM_USER + 76
	TB_CUSTOMIZE             = user32.WM_USER + 27
	TB_ADDSTRING             = user32.WM_USER + 77
	TB_GETITEMRECT           = user32.WM_USER + 29
	TB_BUTTONSTRUCTSIZE      = user32.WM_USER + 30
	TB_SETBUTTONSIZE         = user32.WM_USER + 31
	TB_SETBITMAPSIZE         = user32.WM_USER + 32
	TB_AUTOSIZE              = user32.WM_USER + 33
	TB_GETTOOLTIPS           = user32.WM_USER + 35
	TB_SETTOOLTIPS           = user32.WM_USER + 36
	TB_SETPARENT             = user32.WM_USER + 37
	TB_SETROWS               = user32.WM_USER + 39
	TB_GETROWS               = user32.WM_USER + 40
	TB_GETBITMAPFLAGS        = user32.WM_USER + 41
	TB_SETCMDID              = user32.WM_USER + 42
	TB_CHANGEBITMAP          = user32.WM_USER + 43
	TB_GETBITMAP             = user32.WM_USER + 44
	TB_GETBUTTONTEXT         = user32.WM_USER + 75
	TB_REPLACEBITMAP         = user32.WM_USER + 46
	TB_GETBUTTONSIZE         = user32.WM_USER + 58
	TB_SETBUTTONWIDTH        = user32.WM_USER + 59
	TB_SETINDENT             = user32.WM_USER + 47
	TB_SETIMAGELIST          = user32.WM_USER + 48
	TB_GETIMAGELIST          = user32.WM_USER + 49
	TB_LOADIMAGES            = user32.WM_USER + 50
	TB_GETRECT               = user32.WM_USER + 51
	TB_SETHOTIMAGELIST       = user32.WM_USER + 52
	TB_GETHOTIMAGELIST       = user32.WM_USER + 53
	TB_SETDISABLEDIMAGELIST  = user32.WM_USER + 54
	TB_GETDISABLEDIMAGELIST  = user32.WM_USER + 55
	TB_SETSTYLE              = user32.WM_USER + 56
	TB_GETSTYLE              = user32.WM_USER + 57
	TB_SETMAXTEXTROWS        = user32.WM_USER + 60
	TB_GETTEXTROWS           = user32.WM_USER + 61
	TB_GETOBJECT             = user32.WM_USER + 62
	TB_GETBUTTONINFO         = user32.WM_USER + 63
	TB_SETBUTTONINFO         = user32.WM_USER + 64
	TB_INSERTBUTTON          = user32.WM_USER + 67
	TB_ADDBUTTONS            = user32.WM_USER + 68
	TB_HITTEST               = user32.WM_USER + 69
	TB_SETDRAWTEXTFLAGS      = user32.WM_USER + 70
	TB_GETHOTITEM            = user32.WM_USER + 71
	TB_SETHOTITEM            = user32.WM_USER + 72
	TB_SETANCHORHIGHLIGHT    = user32.WM_USER + 73
	TB_GETANCHORHIGHLIGHT    = user32.WM_USER + 74
	TB_GETINSERTMARK         = user32.WM_USER + 79
	TB_SETINSERTMARK         = user32.WM_USER + 80
	TB_INSERTMARKHITTEST     = user32.WM_USER + 81
	TB_MOVEBUTTON            = user32.WM_USER + 82
	TB_GETMAXSIZE            = user32.WM_USER + 83
	TB_SETEXTENDEDSTYLE      = user32.WM_USER + 84
	TB_GETEXTENDEDSTYLE      = user32.WM_USER + 85
	TB_GETPADDING            = user32.WM_USER + 86
	TB_SETPADDING            = user32.WM_USER + 87
	TB_SETINSERTMARKCOLOR    = user32.WM_USER + 88
	TB_GETINSERTMARKCOLOR    = user32.WM_USER + 89
	TB_MAPACCELERATOR        = user32.WM_USER + 90
	TB_GETSTRING             = user32.WM_USER + 91
	TB_GETIDEALSIZE          = user32.WM_USER + 99
	TB_GETMETRICS            = user32.WM_USER + 101
	TB_SETCOLORSCHEME        = comctl32.CCM_SETCOLORSCHEME
	TB_GETCOLORSCHEME        = comctl32.CCM_GETCOLORSCHEME
	TB_SETUNICODEFORMAT      = comctl32.CCM_SETUNICODEFORMAT
	TB_GETUNICODEFORMAT      = comctl32.CCM_GETUNICODEFORMAT
)

// ToolBar notifications
const (
	TBN_FIRST    = -700
	TBN_DROPDOWN = TBN_FIRST - 10
)

// TBN_DROPDOWN return codes
const (
	TBDDRET_DEFAULT      = 0
	TBDDRET_NODEFAULT    = 1
	TBDDRET_TREATPRESSED = 2
)

// ToolBar state constants
const (
	TBSTATE_CHECKED       = 1
	TBSTATE_PRESSED       = 2
	TBSTATE_ENABLED       = 4
	TBSTATE_HIDDEN        = 8
	TBSTATE_INDETERMINATE = 16
	TBSTATE_WRAP          = 32
	TBSTATE_ELLIPSES      = 0x40
	TBSTATE_MARKED        = 0x0080
)

// ToolBar style constants
const (
	TBSTYLE_BUTTON       = 0
	TBSTYLE_SEP          = 1
	TBSTYLE_CHECK        = 2
	TBSTYLE_GROUP        = 4
	TBSTYLE_CHECKGROUP   = TBSTYLE_GROUP | TBSTYLE_CHECK
	TBSTYLE_DROPDOWN     = 8
	TBSTYLE_AUTOSIZE     = 16
	TBSTYLE_NOPREFIX     = 32
	TBSTYLE_TOOLTIPS     = 256
	TBSTYLE_WRAPABLE     = 512
	TBSTYLE_ALTDRAG      = 1024
	TBSTYLE_FLAT         = 2048
	TBSTYLE_LIST         = 4096
	TBSTYLE_CUSTOMERASE  = 8192
	TBSTYLE_REGISTERDROP = 0x4000
	TBSTYLE_TRANSPARENT  = 0x8000
)

// ToolBar extended style constants
const (
	TBSTYLE_EX_DRAWDDARROWS       = 0x00000001
	TBSTYLE_EX_MIXEDBUTTONS       = 8
	TBSTYLE_EX_HIDECLIPPEDBUTTONS = 16
	TBSTYLE_EX_DOUBLEBUFFER       = 0x80
)

// ToolBar button style constants
const (
	BTNS_BUTTON        = TBSTYLE_BUTTON
	BTNS_SEP           = TBSTYLE_SEP
	BTNS_CHECK         = TBSTYLE_CHECK
	BTNS_GROUP         = TBSTYLE_GROUP
	BTNS_CHECKGROUP    = TBSTYLE_CHECKGROUP
	BTNS_DROPDOWN      = TBSTYLE_DROPDOWN
	BTNS_AUTOSIZE      = TBSTYLE_AUTOSIZE
	BTNS_NOPREFIX      = TBSTYLE_NOPREFIX
	BTNS_WHOLEDROPDOWN = 0x0080
	BTNS_SHOWTEXT      = 0x0040
)

// TBBUTTONINFO mask flags
const (
	TBIF_IMAGE   = 0x00000001
	TBIF_TEXT    = 0x00000002
	TBIF_STATE   = 0x00000004
	TBIF_STYLE   = 0x00000008
	TBIF_LPARAM  = 0x00000010
	TBIF_COMMAND = 0x00000020
	TBIF_SIZE    = 0x00000040
	TBIF_BYINDEX = 0x80000000
)

// TBMETRICS mask flags
const (
	TBMF_PAD           = 0x00000001
	TBMF_BARPAD        = 0x00000002
	TBMF_BUTTONSPACING = 0x00000004
)
