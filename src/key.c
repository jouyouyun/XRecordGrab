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

#include <stdio.h>
#include <stdlib.h>
#include <X11/Xlib.h>

#include "key.h"

static Key_t *grab_key = NULL;

static Key_t*
new_code(unsigned char code)
{
	Key_t *key = (Key_t*)calloc(1, sizeof(Key_t));
	if (key == NULL) {
		fprintf(stderr, "Alloc Key_t memory failed\n");
		return NULL;
	}

	key->code = code;
	key->next = NULL;

	return key;
}

void
add_key(unsigned char code)
{
	Key_t *key = new_code(code);
	if (key == NULL) {
		fprintf(stderr, "add key '%d' failed\n", code);
		return;
	}

	if (grab_key == NULL) {
		grab_key = key;
	} else {
		Key_t *prev = NULL;
		Key_t *tmp = grab_key;
		while (tmp != NULL) {
			prev = tmp;
			tmp = tmp->next;
		}
		prev->next = key;
	}
}

void
parse_key()
{
	if (grab_key == NULL) {
		return;
	}

	Key_t *tmp = grab_key;
	Display *dsp = XOpenDisplay(0);
	int keysym_return;

	printf("\nKey String: ");
	while (tmp != NULL){
		Key_t *key = tmp;
		tmp = tmp->next;

		/*printf("Parse keycode: %d\n", key->code);*/
		KeySym *keysym = XGetKeyboardMapping(dsp, key->code,
				1, &keysym_return);
		if (keysym == NULL) {
			continue;
		}
		/*printf("KeyCode: %d, KeySym: %d\n", */
				/*key->code, keysym[0]);*/
		printf("%s ",XKeysymToString(keysym[0]));
		XFree(keysym);
	}
	printf("\n");
}

void
free_key()
{
	while (grab_key != NULL){
		Key_t *tmp = grab_key;
		grab_key = grab_key->next;

		tmp->next = NULL;
		free(tmp);
	}

	grab_key = NULL;
}
