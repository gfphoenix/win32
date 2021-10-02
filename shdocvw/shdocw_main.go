// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package shdocvw

import "github.com/Gipcomp/winapi/ole32"

var (
	CLSID_WebBrowser            = ole32.CLSID{0x8856F961, 0x340A, 0x11D0, [8]byte{0xA9, 0x6B, 0x00, 0xC0, 0x4F, 0xD7, 0x05, 0xA2}}
	DIID_DWebBrowserEvents2     = ole32.IID{0x34A715A0, 0x6587, 0x11D0, [8]byte{0x92, 0x4A, 0x00, 0x20, 0xAF, 0xC7, 0xAC, 0x4D}}
	IID_IWebBrowser2            = ole32.IID{0xD30C1661, 0xCDAF, 0x11D0, [8]byte{0x8A, 0x3E, 0x00, 0xC0, 0x4F, 0xC9, 0xE2, 0x6E}}
	IID_IDocHostUIHandler       = ole32.IID{0xBD3F23C0, 0xD43E, 0x11CF, [8]byte{0x89, 0x3B, 0x00, 0xAA, 0x00, 0xBD, 0xCE, 0x1A}}
	IID_IOleInPlaceActiveObject = ole32.IID{0x00000117, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
)
