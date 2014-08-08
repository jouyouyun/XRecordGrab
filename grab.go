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
	"errors"
	"fmt"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
)

var X *xgbutil.XUtil
var initFlag bool

func initXUtil() error {
	var err error

	if X, err = xgbutil.NewConn(); err != nil {
		fmt.Println("New XUtil Failed:", err)
		return err
	}

	if !initFlag {
		keybind.Initialize(X)
		initFlag = true
	}

	return nil
}

func stringToKeyCode(key string) (keycode int, err error) {
	if len(key) < 1 {
		return 0, errors.New("Invalid key")
	}

	if X == nil {
		if err = initXUtil(); err != nil {
			return 0, err
		}
	}

	_, codes, e := keybind.ParseString(X, key)
	if e != nil {
		return 0, e
	}

	fmt.Printf("Key: %s, keycode: %v\n", key, codes)

	return int(codes[0]), nil
}
