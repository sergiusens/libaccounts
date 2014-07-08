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
	"reflect"
	"unsafe"
)

// AuthLoginParameters holds a structure of all the login
// parameters for a Service.
type AuthLoginParameters struct {
	AllowedSchemes        string
	AuthPath              string
	AuthorizationEndpoint string
	Callback              string
	ClientId              string
	ClientSecret          string
	ConsumerKey           string
	ConsumerSecret        string
	Display               string
	Host                  string
	RedirectUri           string
	RequestEndpoint       string
	ResponseType          string
	TokenEndpoint         string
	TokenPath             string
}

// AuthData holds an instance of an AgAuthData with all it's
// related information.
type AuthData struct {
	auth            *C.AgAuthData
	Id              uint
	Mechanism       string
	Method          string
	LoginParameters *AuthLoginParameters
}

// NewAuthData creates a new AuthData instance from an AccountService.
func NewAuthData(acc *AccountService) *AuthData {
	auth := C.ag_account_service_get_auth_data(acc.acc)
	return &AuthData{
		auth:            auth,
		Id:              getId(auth),
		Mechanism:       getMechanism(auth),
		Method:          getMethod(auth),
		LoginParameters: getParameters(auth),
	}
}

// Delete removes the AgAuthData instance.
func (auth *AuthData) Delete() {
	C.ag_auth_data_unref(C.to_AgAuthData(unsafe.Pointer(auth.auth)))
}

func (auth *AuthData) String() string {
	return fmt.Sprintf(
		"Id: '%d' | Mechanism: '%s', Method: '%s' | Login Parameters: {'%s'}",
		auth.Id, auth.Mechanism, auth.Method, auth.LoginParameters.String(),
	)
}

func (loginParam *AuthLoginParameters) String() string {
	return fmt.Sprintf(
		"ClientId: '%s' | ClientSecret: '%s' | ConsumerKey: '%s' | ConsumerSecret: '%s'",
		loginParam.ClientId,
		loginParam.ClientSecret,
		loginParam.ConsumerKey,
		loginParam.ConsumerSecret,
	)
}

func getId(auth *C.AgAuthData) uint {
	return uint(C.ag_auth_data_get_credentials_id(auth))
}

func getMechanism(auth *C.AgAuthData) string {
	return getStringFromGCharPtr(C.ag_auth_data_get_mechanism(auth))
}

func getMethod(auth *C.AgAuthData) string {
	return getStringFromGCharPtr(C.ag_auth_data_get_method(auth))
}

func getParameters(auth *C.AgAuthData) *AuthLoginParameters {
	var param, extraParam *C.GVariant
	param = C.ag_auth_data_get_login_parameters(auth, extraParam)
	defer gVariantUnref(param)
	defer gVariantUnref(extraParam)

	var ap AuthLoginParameters
	rAP := reflect.ValueOf(&ap).Elem()

	var iter C.GVariantIter
	var gValue *C.GVariant
	var gKey *C.gchar

	C.g_variant_iter_init(&iter, param)

	for more := C.gboolean(C.get_dict_entry(&iter, &gKey, &gValue)); more == C.TRUE; {
		key := getStringFromGCharPtr(gKey)
		C.g_free(C.gpointer(gKey))

		gType := getStringFromGCharPtr(C.g_variant_get_type_string(gValue))
		switch gType {
		case "s":
			value := getStringFromGCharPtr(C.g_variant_get_string(gValue, nil))
			if f := rAP.FieldByName(key); f.IsValid() {
				f.SetString(value)
			} else {
				fmt.Println(key, "is not valid")
			}
		case "as":
			fmt.Println(key, "for 'as' is not handled")
		default:
			fmt.Println("Unhandled value for", key, "variant type:", gType)
		}

		C.g_variant_unref(gValue)
		more = C.gboolean(C.get_dict_entry(&iter, &gKey, &gValue))
	}
	return &ap
}
