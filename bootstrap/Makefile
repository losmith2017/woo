# ------------------------------------------------------------------------
# Purpose
#   Downloads and installs a small set of utilities to faciliate
#   Woo boostrap development 
#
# Usage
#   Default target will print help message
#
# Note
#   Installation does not check for correct versions of virtualbox and 
#   vagrant, yet. 
#
# Default Installation
#   $HOME/.woo/cache/<downloads...>
#   $HOME/.woo/dist/go-<version>
#   $HOME/.woo/dist/packer-<version>
#   $HOME/.woo/dist/vault-<version>
#   /usr/local/opt/vagrant -> /opt/vagrant
#   /Applications/VirtualBox.app
#
# Utilities Purpose
#   golang, utilities to build, deploy, manage workspaces and services
#   packer, create vm images to host docker engine nodes
#   vagrant, manage vm nodes
#   vault, manage workspace secrets (keys, tokens, passwords)
#   virtualbox, hypervisor to run docker nodes
# ------------------------------------------------------------------------

WOO_HOME                = $(HOME)/.woo
CACHEDIR                = $(WOO_HOME)/cache
DISTDIR                 = $(WOO_HOME)/dist
TEMPDIR                 = $(WOO_HOME)/tmp

GOLANG_VERSION          = 1.9.2
GOLANG_ARCHIVE          = $(CACHEDIR)/go$(GOLANG_VERSION).darwin-amd64.tar.gz
GOLANG                  = $(DISTDIR)/go-$(GOLANG_VERSION)

PACKER_VERSION          = 1.1.2
PACKER_ARCHIVE          = $(CACHEDIR)/packer_$(PACKER_VERSION)_darwin_amd64.zip
PACKER                  = $(DISTDIR)/packer-$(PACKER_VERSION)/bin/packer

VAGRANT_VERSION         = 2.0.1
VAGRANT_ARCHIVE         = $(CACHEDIR)/vagrant_$(VAGRANT_VERSION)_x86_64.dmg
VAGRANT                 = /usr/local/bin/vagrant

VAULT_VERSION           = 0.9.0
VAULT_ARCHIVE           = $(CACHEDIR)/vault_$(VAULT_VERSION)_darwin_amd64.zip
VAULT                   = $(DISTDIR)/vault-$(VAULT_VERSION)/bin/vault

VIRTUALBOX_VERSION      = 5.2.2
VIRTUALBOX_ARCHIVE      = $(CACHEDIR)/VirtualBox-$(VIRTUALBOX_VERSION)-119230-OSX.dmg
VIRTUALBOX              = /Applications/VirtualBox.app

ALL_TARGETS             = golang packer vagrant vault virtualbox


# ------------------------------------------------------------------------
# Driver Targets
# ------------------------------------------------------------------------
help:; @for i in $(ALL_TARGETS); do echo make $$i; done
all: $(ALL_TARGETS)
golang: dirs $(GOLANG) 
packer: dirs $(PACKER) 
vagrant: dirs $(VAGRANT)
vault: dirs $(VAULT)
virtualbox: dirs $(VIRTUALBOX)
dirs: $(CACHEDIR) $(DISTDIR)


# ------------------------------------------------------------------------
# 
# ------------------------------------------------------------------------
$(DISTDIR):
	@echo making distribution directory... $@
	@mkdir -p $@

$(CACHEDIR):
	@echo making cache directory... $@
	@mkdir -p $@

$(GOLANG_ARCHIVE):
	@echo downloading... $(@F)
	curl -L -o $@ https://redirector.gvt1.com/edgedl/go/$(@F)
	openssl dgst -sha256 $@

$(GOLANG): 
	@echo installing... $(@F)
	@if [ ! -f $(GOLANG_ARCHIVE) ]; then make $(GOLANG_ARCHIVE); fi
	mkdir -p $@
	tar -C $@ -xvf $(GOLANG_ARCHIVE)
	mv $(@)/go/* $@
	rm -rf $(@)/go

$(PACKER_ARCHIVE):
	@echo downloading... $(@F)
	curl -L -o $@ https://releases.hashicorp.com/packer/$(PACKER_VERSION)/$(@F)
	openssl dgst -sha256 $@

$(PACKER): 
	@echo installing... $(@F) $(PACKER_VERSION)
	@if [ ! -f $(PACKER_ARCHIVE) ]; then make $(PACKER_ARCHIVE); fi
	mkdir -p $(@D)
	cd $(@D) && unzip $(PACKER_ARCHIVE)

$(VAGRANT_ARCHIVE):
	@echo downloading... $(@F)
	curl -L -o $@ https://releases.hashicorp.com/vagrant/$(VAGRANT_VERSION)/$(@F)
	openssl dgst -sha256 $@

$(VAGRANT): 
	@echo installing... $(@F) $(VAGRANT_VERSION)
	@if [ ! -f $(VAGRANT_ARCHIVE) ]; then make $(VAGRANT_ARCHIVE); fi
	hdiutil attach $(VAGRANT_ARCHIVE)
	sudo installer -pkg /Volumes/Vagrant/vagrant.pkg -target /Volumes/Macintosh\ HD
	hdiutil detach /Volumes/Vagrant

$(VAULT_ARCHIVE):
	@echo downloading... $(@F)
	curl -L -o $@ https://releases.hashicorp.com/vault/$(VAULT_VERSION)/$(@F)
	openssl dgst -sha256 $@

$(VAULT): 
	@echo installing... $(@F) $(VAULT_VERSION)
	@if [ ! -f $(VAULT_ARCHIVE) ]; then make $(VAULT_ARCHIVE); fi
	mkdir -p $(@D)
	cd $(@D) && unzip $(VAULT_ARCHIVE)

$(VIRTUALBOX_ARCHIVE):
	@echo downloading... $(@F)
	curl -L -o $@ http://download.virtualbox.org/virtualbox/$(VIRTUALBOX_VERSION)/$(@F)
	openssl dgst -sha256 $@

$(VIRTUALBOX): 
	@echo installing... $@ $(VIRTUALBOX_VERSION)
	@if [ ! -f $(VIRTUALBOX_ARCHIVE) ]; then make $(VIRTUALBOX_ARCHIVE); fi
	hdiutil attach $(VIRTUALBOX_ARCHIVE)
	sudo installer -pkg /Volumes/VirtualBox/VirtualBox.pkg -target /Volumes/Macintosh\ HD
	hdiutil detach /Volumes/VirtualBox

.PHONY: all dirs golang help packer vagrant vault virtualbox 
