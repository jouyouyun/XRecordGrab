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

import (
	"strings"
)

func formatShortcut(shortcut string) (retStr string) {
	if len(shortcut) < 1 {
		return
	}

	flag := false
	start := 0
	end := 0
	for i, ch := range shortcut {
		if ch == '<' {
			flag = true
			start = i
			continue
		}

		if ch == '>' && flag {
			flag = false
			end = i

			for j := start + 1; j < end; j++ {
				retStr += string(shortcut[j])
			}
			retStr += "-"
			continue
		}

		if !flag {
			retStr += string(ch)
		}
	}

	// convert 'primary' to 'control'
	list := strings.Split(retStr, "-")
	retStr = ""
	for i, v := range list {
		if strings.ToLower(v) == "primary" ||
			strings.ToLower(v) == "control" {
			//multiple 'control'
			if !strings.Contains(strings.ToLower(retStr), "control") {
				if i != 0 {
					retStr += "-"
				}
				retStr = "Control"
			}
			continue
		}

		if i != 0 {
			retStr += "-"
		}
		retStr += v
	}

	return
}
