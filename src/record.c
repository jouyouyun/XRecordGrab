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
#include <pthread.h>
#include <X11/Xlib.h>
#include <X11/extensions/record.h>

typedef struct _RecordData {
	Display *ctrl_dsp;
	Display *data_dsp;
	XRecordRange *range;
	XRecordContext context;
} RecordData;

static void *enable_ctx_thread(void *user_data);
static void intercept_cb (XPointer user_data, XRecordInterceptData *hook);

static RecordData *grab_data = NULL;

void
record_init()
{
	grab_data = (RecordData*)calloc(1, sizeof(RecordData));
	if (grab_data == NULL) {
		fprintf(stderr, "Alloc RecordData memory failed\n");
		return;
	}

	grab_data->ctrl_dsp = XOpenDisplay(NULL);
	grab_data->data_dsp = XOpenDisplay(NULL);
	if (grab_data->ctrl_dsp == NULL || grab_data->data_dsp == NULL) {
		printf(stderr, "Open Display Failed\n");
		record_finalize();
		return;
	}

	int major, first_event, first_error;
	if (!XQueryExtension(grab_data->ctrl_dsp, "XTEST",
	                     &dummy, &first_event, &first_error)) {
		fprintf(stderr, "XTest extension missing...\n");
		record_finalize();
		return;
	}

	int minor;
	if (!XRecordQueryVersion(grab_data->ctrl_dsp, &major, &minor)) {
		fprintf(stderr, "Failed to obtain XRecord version\n");
		record_finalize();
		return
	}

	grab_data->range = XRecordAllocRange();
	if (!grab_data->range) {
		fprintf(stderr, "Alloc XRecordRange memory failed\n");
		record_finalize();
		return;
	}

	grab_data->range->device_events.first = KeyPress;
	grab_data->range->device_events.last = KeyRelease;
	/*grab_data->range->device_events.last = ButtonRelease;*/

	XRecordClientSpec spec = XRecordAllClients;
	grab_data->context = XRecordCreateContext(grab_data->data_dsp,
	                     0, &spec, 1, &grab_data->range, 1);
	if (!grab_data->context) {
		fprintf(stderr, "Unable to create context...\n");
		record_finalize();
		return;
	}

	XSynchronize(grab_data->ctrl_dsp, True);
	XFlush(grab_data->ctrl_dsp);

	pthread_t thrd;
	pthread_attr_t attr;

	// Free thread resource when thread terminates
	pthread_attr_init(&attr);
	pthread_attr_setdetachstate(&attr, PTHREAD_CREATE_DETACHED);
	int ret = pthread_create(&thrd, &attr, enable_ctx_thread, NULL);
	pthread_attr_destroy(&attr);

	if (ret != 0 ) {
		fprintf(stderr, "Create context thread failed...\n");
		record_finalize();
		return;
	}
}

void
record_finalize()
{
	if (!grab_data) {
		return;
	}

	if (grab_data->context) {
		XRecordDisableContext(grab_data->data_dsp, grab_data->context);
		XRecordFreeContext(grab_data->data_dsp, grab_data->context);
		grab_data->context = 0;
	}

	if (grab_data->range) {
		XFree(grab_data->range);
		grab_data->range = NULL;
	}

	if (grab_data->data_dsp) {
		XCloseDisplay(grab_data->data_dsp);
		grab_data->data_dsp = NULL;
	}

	if (grab_data->ctrl_dsp) {
		XCloseDisplay(grab_data->ctrl_dsp);
		grab_data->ctrl_dsp = NULL;
	}

	if (grab_data) {
		free(grab_data);
		grab_data = NULL;
	}
}

static void*
enable_ctx_thread(void *user_data)
{
	if (!XRecordEnableContext(grab_data->data_dsp, grab_data->context,
				intercept_cb, NULL)) {
		fprintf(stderr, "Unable to enable context...\n");
		record_finalize();
	}

	pthread_exit(NULL);
}

static void
intercept_cb (XPointer user_data, XRecordInterceptData *hook)
{
}
