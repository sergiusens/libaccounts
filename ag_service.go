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
import (
	"fmt"
	"unsafe"
)

// Service holds an instance of an AgService with all it's
// related information.
type Service struct {
	svc         *C.AgService
	Name        string
	DisplayName string
	Description string
	Provider    string
	ServiceType string
	Tags        []string
}

// NewService creates a new Service instance from an AccountService.
func NewService(acc *AccountService) *Service {
	svc := C.ag_account_service_get_service(acc.acc)
	return &Service{
		svc:         svc,
		Name:        getStringFromGCharPtr(C.ag_service_get_name(svc)),
		DisplayName: getStringFromGCharPtr(C.ag_service_get_display_name(svc)),
		Description: getStringFromGCharPtr(C.ag_service_get_description(svc)),
		Provider:    getStringFromGCharPtr(C.ag_service_get_provider(svc)),
		ServiceType: getStringFromGCharPtr(C.ag_service_get_service_type(svc)),
		Tags:        getTags(svc),
	}
}

// Delete removes the AgAccountService instance.
func (svc *Service) Delete() {
	C.ag_service_unref(C.to_AgService(unsafe.Pointer(svc.svc)))
}

func (svc *Service) String() string {
	return fmt.Sprintf(
		"Name: %s(%s); ServiceType: %s; Tags: %s",
		svc.Name, svc.DisplayName, svc.Provider, svc.Tags)
}

func getTags(svc *C.AgService) []string {
	tags := C.ag_service_get_tags(svc)
	defer C.g_list_free(tags)

	length := C.g_list_length(tags)
	result := make([]string, length)

	for n := C.guint(0); n < length; n++ {
		data := C.g_list_nth_data(tags, n)
		pointer := C.to_charptr_from_ptr(data)
		result = append(result, C.GoString(pointer))
	}

	return result
}
