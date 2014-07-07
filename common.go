/*
 * Copyright 2014 Canonical Ltd.
 *
 * Authors:
 * Sergio Schvezov <sergio.schvezov@canonical.com>
 *
 * This file is part of libaccounts.
 */

package libaccounts

/*
#cgo pkg-config: libaccounts-glib
#include <stdlib.h>
#include <glib.h>
#include <libaccounts-glib/accounts-glib.h>
#include "common.h"
*/
import "C"

func getStringFromGCharPtr(data *C.gchar) string {
	if data == nil {
		return ""
	}
	return C.GoString(C.to_charptr(data))
}

func gVariantUnref(v *C.GVariant) {
	if v != nil {
		C.g_variant_unref(v)
	}
}
