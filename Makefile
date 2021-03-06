# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=rbot
GOFILES=rbot.go handler.go auth.go cmd-access.go cmd-admin.go cmd-op.go cmd-google.go cmd-booru.go cmd-ap.go bitly.go cmd-help.go cmd-misc.go
pkgdir=$(QUOTED_GOROOT)/pkg/$(GOOS)_$(GOARCH)
PREREQ=$(pkgdir)/github.com/fluffle/goirc/event.a $(pkgdir)/github.com/fluffle/goirc/client.a $(pkgdir)/goconfig.a

all: rbot.conf auth.conf bans.list ap.conf

.PHONY: event client goconfig

$(pkgdir)/github.com/fluffle/goirc/event.a: event
	@true
$(pkgdir)/github.com/fluffle/goirc/client.a: client
	@true
$(pkgdir)/goconfig.a: goconfig
	@true

event:
	$(MAKE) -sC event install
client:
	$(MAKE) -sC client install
goconfig:
	$(MAKE) -sC goconfig install

include $(GOROOT)/src/Make.cmd

rbot.conf: rbot.conf.example
	@if [ -f $@ ] ; then \
		echo "rbot.conf exists, but rbot.conf.example is newer." ; \
	else \
		echo cp $< $@ ; \
		cp $< $@ ; \
	fi

auth.conf: auth.conf.example
	@if [ -f $@ ] ; then \
		echo "auth.conf exists, but auth.conf.example is newer." ; \
	else \
		echo cp $< $@ ; \
		cp $< $@ ; \
	fi

bans.list:
	@if [ ! -f $@ ] ; then \
		touch $@ ; \
		echo "created $@" ; \
	fi

ap.conf:
	@if [ ! -f $@ ] ; then \
		touch $@ ; \
		echo "created $@" ; \
	fi

clean: clean-deps
clean-deps:
	$(MAKE) -C event clean
	$(MAKE) -C client clean
	$(MAKE) -C goconfig clean

nuke: nuke-deps
nuke-deps:
	$(MAKE) -C event nuke
	$(MAKE) -C client nuke
	$(MAKE) -C goconfig nuke

uninstall:
	@echo Perhaps you meant \"make nuke\"
