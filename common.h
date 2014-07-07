/*
 * Copyright 2014 Canonical Ltd.
 *
 * Authors:
 * Sergio Schvezov <sergio.schvezov@canonical.com>
 *
 * This file is part of libaccounts.
 */

#ifndef __COMMON_ACCOUNTS_H__
#define __COMMON_ACCOUNTS_H__
static inline void free_string(char* s) { free(s); }

static GError* to_error(void* err) { return (GError*)err; }

static inline char* to_charptr(const gchar* s) { return (char*)s; }
static inline char* to_charptr_from_ptr(gpointer s) { return (char*)s; }
static inline gpointer to_gpointer_from_charptr(const gchar *s) { return (gpointer)s; }
static inline gpointer to_gpointer_from_gvalue(const GValue *s) { return (gpointer)s; }
static inline gchar* to_gcharptr(const char* s) { return (gchar*)s; }

static inline AgService* to_AgService(void* o) { return (AgService *)o; }
static inline AgAccountService* to_AgAccountService(void* o) { return (AgAccountService *)o; }
static inline AgApplication* to_AgApplication(void* o) { return (AgApplication *)o; }
static inline AgProvider* to_AgProvider(void* o) { return (AgProvider *)o; }
static inline AgServiceType* to_AgServiceType(void* o) { return (AgServiceType *)o; }
static inline AgAuthData* to_AgAuthData(void* o) { return (AgAuthData *)o; }

static inline gboolean get_dict_entry(GVariantIter *iter, gchar **key, GVariant **value) {
    return g_variant_iter_next(iter, "{sv}", key, value);
}
#endif
