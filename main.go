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

import "fmt"
import "sort"

var bindMap map[string]string

func main() {
	fmt.Println("Hello World")

	bindMap = make(map[string]string)

	keyList := []string{
		"Super_L", "Super_R",
		"Control_L", "Control_R",
		"Alt_L", "Alt_R",
		"Shift_L", "Shift_R",
		"y", "Y",
	}

	list := []string{"Control_L", "Alt_L", "y"}
	for _, v := range keyList {
		stringToKeyCode(v)
	}

	codes := []int{}
	for _, v := range list {
		if code, err := stringToKeyCode(v); err != nil {
			continue
		} else {
			codes = append(codes, code)
		}
	}
	sort.Ints(codes)

	if str, err := encodeIntList(codes); err == nil {
		bindMap[str] = "Test grab keycode"
	}

	defer finalizeRecord()
	initRecord()

	select {}
}
