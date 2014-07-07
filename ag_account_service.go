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
import

// AccountService holds an instance of an AgAccountService
// with all it's related information.
"unsafe"

type AccountService struct {
	acc *C.AgAccountService
}

// GetService returns the Service for the AccountService
func (acc *AccountService) GetService() *Service {
	return NewService(acc)
}

// GetAuthData returns the AuthData for the AccountService
func (acc *AccountService) GetAuthData() *AuthData {
	return NewAuthData(acc)
}

// Delete removes the AgAccountService instance.
func (acc *AccountService) Delete() {
	C.g_object_unref(C.gpointer(unsafe.Pointer(acc.acc)))
}
