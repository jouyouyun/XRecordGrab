/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package main

// #cgo pkg-config: x11 xtst
// #cgo CFLAGS: -Wall -g
// #cgo LDFLAGS: -Wall -g -lpthread
// #include "record.h"
import "C"

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"sort"
)

var _curKeyList []int

//export add_keycode_to_list
func add_keycode_to_list(code int) {
	_curKeyList = append(_curKeyList, code)
}

//export parse_keycode_list
func parse_keycode_list() {
	if len(_curKeyList) < 1 {
		return
	}

	fmt.Println("Before sort:", _curKeyList)
	sort.Ints(_curKeyList)
	fmt.Println("After sort:", _curKeyList)

	if str, err := encodeIntList(_curKeyList); err != nil {
		fmt.Println("Encode int list failed:", err)
	} else {
		if v, ok := bindMap[str]; ok {
			fmt.Println("Exec:", v)
		}
	}

	_curKeyList = []int{}
}

func initRecord() {
	C.record_init()
}

func finalizeRecord() {
	C.record_finalize()
}

func encodeIntList(list []int) (string, error) {
	if len(list) < 1 {
		return "", errors.New("Invalid int list")
	}

	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(&list); err != nil {
		return "", err
	}

	return buf.String(), nil
}
