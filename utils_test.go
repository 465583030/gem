// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package gem

import (
	"bytes"
	"strings"
	"testing"
)

var (
	s = "str"
	b = []byte("str")
)

func TestBytes2String(t *testing.T) {
	if strings.Compare(s, Bytes2String(b)) != 0 {
		t.Errorf(`unexpected: strings.Compare("%s", Bytes2String([]byte("%s"))) got false want true.`, s, s)
	}
}

func TestString2Bytes(t *testing.T) {
	if !bytes.Equal(b, String2Bytes(s)) {
		t.Errorf(`unexpected: bytes.Equal([]byte("%s"), String2Bytes("%s")) got false want true.`, s, s)
	}
}
