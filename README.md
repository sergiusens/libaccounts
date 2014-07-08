# Online Accounts bindings for golang.

Introduction
------------

This is an Alpha release. While the API presented here could be considered stable, 
due to the matureness of the development of this package it could potentially
change. The API is also far from complete.

This wraps arround [accounts-sso](https://code.google.com/p/accounts-sso/),
specifically using the 
[glib API](http://docs.accounts-sso.googlecode.com/git/libaccounts-glib/html/index.html)

These binding are highly inspired by mandel's 
[blog post](http://www.themacaque.com/?p=1133)
for online accounts.

Requirements
------------

To build this package you will need to have on Ubuntu:

    $ sudo apt-get install libaccounts-glib-dev


Using
-----

Here's an example of usage:

    m := libaccounts.NewManager()
    defer m.Delete()
    ac := m.GetEnabledAccountServices()

	for i := range ac {
        defer ac[i].Delete()
        svc := ac[i].GetService()
        defer svc.Delete()
        auth := ac[i].GetAuthData()
        defer auth.Delete()
        fmt.Println(svc)
        fmt.Println("Auth data :", auth)
        fmt.Println()
	}

