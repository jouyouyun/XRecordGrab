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

var (
	_keyToModMap = map[string]string{
		"caps_lock": "lock",
		"alt":       "mod1",
		"meta":      "mod1",
		"num_lock":  "mod2",
		"super":     "mod4",
		"hyper":     "mod4",
	}

	_modToKeyMap = map[string]string{
		"mod1": "alt",
		"mod2": "num_lock",
		"mod4": "super",
		"lock": "caps_lock",
	}
)

func convertkeyToMod(key string) (mod string) {
	mod = key
	if str, ok := _keyToModMap[key]; ok {
		mod = str
	}

	return
}

func convertModToKey(mod string) (key string) {
	key = mod
	if str, ok := _modToKeyMap[mod]; ok {
		key = str
	}

	return
}

func convertShortcutToModStr(shortcut string) (modStr string) {
	list := strings.Split(shortcut, "-")
	lenght := len(list)
	for i, key := range list {
		mod := convertkeyToMod(key)
		modStr += mod
		if i != lenght-1 {
			modStr += "-"
		}
	}

	return
}

func convertModStrToShortcut(modStr string) (shortcut string) {
	list := strings.Split(modStr, "-")
	lenght := len(list)
	for i, mod := range list {
		key := convertModToKey(mod)
		shortcut += key
		if i != lenght-1 {
			shortcut += "-"
		}
	}

	return
}

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

func modifierToKeyName(modStr string) []string {
	if len(modStr) < 1 {
		return []string{}
	}

	tmpStr := strings.ToLower(modStr)
	switch tmpStr {
	case "control":
		return []string{"Control_L", "Control_R"}
	case "shift":
		return []string{"Shift_L", "Shift_R"}
	case "super":
		return []string{"Super_L", "Super_R"}
	case "alt":
		return []string{"Alt_L", "Alt_R"}
	}

	return []string{modStr}
}

func shortcutToKeyNameList(shortcut string) []string {
	if len(shortcut) < 1 {
		return []string{}
	}

	lshortcut := ""
	rshortcut := ""

	shortcut = formatShortcut(shortcut)
	list := strings.Split(shortcut, "-")
	lenght := len(list)
	multiFlag := false
	for i, v := range list {
		if len(v) < 1 {
			continue
		}

		l := modifierToKeyName(v)
		if len(l) == 2 {
			multiFlag = true
			lshortcut += l[0]
			rshortcut += l[1]
		} else {
			lshortcut += v
			rshortcut += v
		}

		if i != lenght-1 {
			lshortcut += "-"
			rshortcut += "-"
		}
	}

	if multiFlag {
		return []string{lshortcut}
	}

	return []string{lshortcut, rshortcut}
}

func shortcutToXgbShortcut(shortcut string) (xgbShortcut string) {
	shortcut = formatShortcut(shortcut)
	if len(shortcut) < 1 {
		return
	}

	shortcut = strings.ToLower(shortcut)
	xgbShortcut = convertShortcutToModStr(shortcut)
	return
}

func xgbShortcutToShortcut(xgbShortcut string) (shortcut string) {
	if len(xgbShortcut) < 1 {
		return
	}

	shortcut = convertModStrToShortcut(xgbShortcut)
	return
}
