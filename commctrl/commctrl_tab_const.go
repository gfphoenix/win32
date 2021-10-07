// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import "github.com/Gipcomp/win32/comctl32"

const TCM_FIRST = 0x1300
const TCN_FIRST = -550

const (
	TCS_SCROLLOPPOSITE    = 0x0001
	TCS_BOTTOM            = 0x0002
	TCS_RIGHT             = 0x0002
	TCS_MULTISELECT       = 0x0004
	TCS_FLATBUTTONS       = 0x0008
	TCS_FORCEICONLEFT     = 0x0010
	TCS_FORCELABELLEFT    = 0x0020
	TCS_HOTTRACK          = 0x0040
	TCS_VERTICAL          = 0x0080
	TCS_TABS              = 0x0000
	TCS_BUTTONS           = 0x0100
	TCS_SINGLELINE        = 0x0000
	TCS_MULTILINE         = 0x0200
	TCS_RIGHTJUSTIFY      = 0x0000
	TCS_FIXEDWIDTH        = 0x0400
	TCS_RAGGEDRIGHT       = 0x0800
	TCS_FOCUSONBUTTONDOWN = 0x1000
	TCS_OWNERDRAWFIXED    = 0x2000
	TCS_TOOLTIPS          = 0x4000
	TCS_FOCUSNEVER        = 0x8000
)

const (
	TCS_EX_FLATSEPARATORS = 0x00000001
	TCS_EX_REGISTERDROP   = 0x00000002
)

const (
	TCM_GETIMAGELIST     = TCM_FIRST + 2
	TCM_SETIMAGELIST     = TCM_FIRST + 3
	TCM_GETITEMCOUNT     = TCM_FIRST + 4
	TCM_GETITEM          = TCM_FIRST + 60
	TCM_SETITEM          = TCM_FIRST + 61
	TCM_INSERTITEM       = TCM_FIRST + 62
	TCM_DELETEITEM       = TCM_FIRST + 8
	TCM_DELETEALLITEMS   = TCM_FIRST + 9
	TCM_GETITEMRECT      = TCM_FIRST + 10
	TCM_GETCURSEL        = TCM_FIRST + 11
	TCM_SETCURSEL        = TCM_FIRST + 12
	TCM_HITTEST          = TCM_FIRST + 13
	TCM_SETITEMEXTRA     = TCM_FIRST + 14
	TCM_ADJUSTRECT       = TCM_FIRST + 40
	TCM_SETITEMSIZE      = TCM_FIRST + 41
	TCM_REMOVEIMAGE      = TCM_FIRST + 42
	TCM_SETPADDING       = TCM_FIRST + 43
	TCM_GETROWCOUNT      = TCM_FIRST + 44
	TCM_GETTOOLTIPS      = TCM_FIRST + 45
	TCM_SETTOOLTIPS      = TCM_FIRST + 46
	TCM_GETCURFOCUS      = TCM_FIRST + 47
	TCM_SETCURFOCUS      = TCM_FIRST + 48
	TCM_SETMINTABWIDTH   = TCM_FIRST + 49
	TCM_DESELECTALL      = TCM_FIRST + 50
	TCM_HIGHLIGHTITEM    = TCM_FIRST + 51
	TCM_SETEXTENDEDSTYLE = TCM_FIRST + 52
	TCM_GETEXTENDEDSTYLE = TCM_FIRST + 53
	TCM_SETUNICODEFORMAT = comctl32.CCM_SETUNICODEFORMAT
	TCM_GETUNICODEFORMAT = comctl32.CCM_GETUNICODEFORMAT
)

const (
	TCIF_TEXT       = 0x0001
	TCIF_IMAGE      = 0x0002
	TCIF_RTLREADING = 0x0004
	TCIF_PARAM      = 0x0008
	TCIF_STATE      = 0x0010
)

const (
	TCIS_BUTTONPRESSED = 0x0001
	TCIS_HIGHLIGHTED   = 0x0002
)

const (
	TCHT_NOWHERE     = 0x0001
	TCHT_ONITEMICON  = 0x0002
	TCHT_ONITEMLABEL = 0x0004
	TCHT_ONITEM      = TCHT_ONITEMICON | TCHT_ONITEMLABEL
)

const (
	TCN_KEYDOWN     = TCN_FIRST - 0
	TCN_SELCHANGE   = TCN_FIRST - 1
	TCN_SELCHANGING = TCN_FIRST - 2
	TCN_GETOBJECT   = TCN_FIRST - 3
	TCN_FOCUSCHANGE = TCN_FIRST - 4
)
