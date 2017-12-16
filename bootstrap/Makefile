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

CONSUL_VERSION          = 1.0.2
CONSUL_ARCHIVE          = $(CACHEDIR)/consul_$(CONSUL_VERSION)_darwin_amd64.zip
CONSUL_DESTDIR          = $(DISTDIR)/consul-$(CONSUL_VERSION)
CONSUL                  = $(CONSUL_DESTDIR)/bin/consul

GOLANG_VERSION          = 1.9.2
GOLANG_ARCHIVE          = $(CACHEDIR)/go$(GOLANG_VERSION).darwin-amd64.tar.gz
GOLANG_DESTDIR          = $(DISTDIR)/go-$(GOLANG_VERSION)
GOLANG                  = $(GOLANG_DESTDIR)/bin/go

PACKER_VERSION          = 1.1.2
PACKER_ARCHIVE          = $(CACHEDIR)/packer_$(PACKER_VERSION)_darwin_amd64.zip
PACKER_DESTDIR          = $(DISTDIR)/packer-$(PACKER_VERSION)
PACKER                  = $(PACKER_DESTDIR)/bin/packer

VAGRANT_VERSION         = 2.0.1
VAGRANT_ARCHIVE         = $(CACHEDIR)/vagrant_$(VAGRANT_VERSION)_x86_64.dmg
VAGRANT_DESTDIR         = /opt/vagrant
VAGRANT                 = $(VAGRANT_DESTDIR)/vagrant

VAULT_VERSION           = 0.9.0
VAULT_ARCHIVE           = $(CACHEDIR)/vault_$(VAULT_VERSION)_darwin_amd64.zip
VAULT_DESTDIR           = $(DISTDIR)/vault-$(VAULT_VERSION)
VAULT                   = $(VAULT_DESTDIR)/bin/vault

VIRTUALBOX_VERSION      = 5.2.2
VIRTUALBOX_ARCHIVE      = $(CACHEDIR)/VirtualBox-$(VIRTUALBOX_VERSION)-119230-OSX.dmg
VIRTUALBOX              = /Applications/VirtualBox.app

BASE_TARGETS           += golang
BASE_TARGETS           += packer
BASE_TARGETS           += vagrant
BASE_TARGETS           += vault
BASE_TARGETS           += virtualbox

INSTALL_TARGETS         = $(BASE_TARGETS:%=%-install)
UNINSTALL_TARGETS       = $(BASE_TARGETS:%=%-uninstall)
ALL_TARGETS             = $(INSTALL_TARGETS) $(UNINSTALL_TARGETS)

# ------------------------------------------------------------------------
# Consul
# ------------------------------------------------------------------------
consul-install: $(CONSUL)

consul-uninstall:
	@echo removing... Consul $(CONSUL_VERSION)
	@rm -rf $(CONSUL_DESTDIR)
	@rm -f $(CONSUL_ARCHIVE)

$(CONSUL_ARCHIVE): 
	@echo downloading... $(@F)
	@mkdir -p $(@D)
	curl -L -o $@ https://releases.hashicorp.com/consul/$(CONSUL_VERSION)/$(@F)
	openssl dgst -sha256 $@

$(CONSUL): 
	@echo installing... Consul $(CONSUL_VERSION)
	@if [ ! -f $(CONSUL_ARCHIVE) ]; then make $(CONSUL_ARCHIVE); fi
	@mkdir -p $(CONSUL_DESTDIR)/bin
	cd $(CONSUL_DESTDIR)/bin && unzip $(CONSUL_ARCHIVE)

# ------------------------------------------------------------------------
# Installation Driver Targets
# ------------------------------------------------------------------------
help:; @for i in $(ALL_TARGETS); do echo make $$i; done
install: $(INSTALL_TARGETS)
uninstall: $(UNINSTALL_TARGETS)

# ------------------------------------------------------------------------
# Golang
# ------------------------------------------------------------------------
golang-install: $(GOLANG)

golang-uninstall: 
	@echo removing... Go $(GOLANG_VERSION)
	@rm -rf $(GOLANG_DESTDIR)
	@rm -f $(GOLANG_ARCHIVE)

$(GOLANG_ARCHIVE): 
	@echo downloading... $(@F)
	@mkdir -p $(@D)
	curl -L -o $@ https://redirector.gvt1.com/edgedl/go/$(@F)
	openssl dgst -sha256 $@

$(GOLANG): 
	@echo installing... Go $(GOLANG_VERSION)
	@if [ ! -f $(GOLANG_ARCHIVE) ]; then make $(GOLANG_ARCHIVE); fi
	@mkdir -p $(GOLANG_DESTDIR)
	tar -C $(GOLANG_DESTDIR) -xvf $(GOLANG_ARCHIVE)
	mv $(GOLANG_DESTDIR)/go/* $(GOLANG_DESTDIR)
	rm -rf $(GOLANG_DESTDIR)/go

# ------------------------------------------------------------------------
# Packer
# ------------------------------------------------------------------------
packer-install: $(PACKER)

packer-uninstall:
	@echo removing... Packer $(PACKER_VERSION)
	@rm -rf $(PACKER_DESTDIR)
	@rm -f $(PACKER_ARCHIVE)

$(PACKER_ARCHIVE): 
	@echo downloading... $(@F)
	@mkdir -p $(@D)
	curl -L -o $@ https://releases.hashicorp.com/packer/$(PACKER_VERSION)/$(@F)
	openssl dgst -sha256 $@

$(PACKER): 
	@echo installing... Packer $(PACKER_VERSION)
	@if [ ! -f $(PACKER_ARCHIVE) ]; then make $(PACKER_ARCHIVE); fi
	@mkdir -p $(PACKER_DESTDIR)/bin
	cd $(PACKER_DESTDIR)/bin && unzip $(PACKER_ARCHIVE)

# ------------------------------------------------------------------------
# Vagrant
# ------------------------------------------------------------------------
vagrant-install: $(VAGRANT)

vagrant-uninstall:
	@echo uninstalling... Vagrant $(VAGRANT_VERSION)
	@if [ ! -f $(VAGRANT_ARCHIVE) ]; then make $(VAGRANT_ARCHIVE); fi
	hdiutil attach $(VAGRANT_ARCHIVE)
	/Volumes/Vagrant/uninstall.tool
	hdiutil detach /Volumes/Vagrant
	rm -f $(VAGRANT_ARCHIVE)

$(VAGRANT_ARCHIVE):
	@echo downloading... Vagrant $(VAGRANT_VERSION)
	@mkdir -p $(@D)
	curl -L -o $@ https://releases.hashicorp.com/vagrant/$(VAGRANT_VERSION)/$(@F)
	openssl dgst -sha256 $@

$(VAGRANT): 
	@echo installing... Vagrant $(VAGRANT_VERSION)
	@if [ ! -f $(VAGRANT_ARCHIVE) ]; then make $(VAGRANT_ARCHIVE); fi
	hdiutil attach $(VAGRANT_ARCHIVE)
	sudo installer -pkg /Volumes/Vagrant/vagrant.pkg -target /Volumes/Macintosh\ HD
	hdiutil detach /Volumes/Vagrant


# ------------------------------------------------------------------------
# Vault
# ------------------------------------------------------------------------
vault-install: $(VAULT)

vault-uninstall:
	@echo removing... Vault $(VAULT_VERSION)
	@rm -rf $(VAULT_DESTDIR)
	@rm -f $(VAULT_ARCHIVE)

$(VAULT_ARCHIVE): 
	@echo downloading... Vault $(VAULT_VERSION)
	@mkdir -p $(@D)
	curl -L -o $@ https://releases.hashicorp.com/vault/$(VAULT_VERSION)/$(@F)
	openssl dgst -sha256 $@

$(VAULT): 
	@echo installing... Vault $(VAULT_VERSION)
	@if [ ! -f $(VAULT_ARCHIVE) ]; then make $(VAULT_ARCHIVE); fi
	@mkdir -p $(VAULT_DESTDIR)/bin
	cd $(VAULT_DESTDIR)/bin && unzip $(VAULT_ARCHIVE)

# ------------------------------------------------------------------------
# VirtualBox 
# ------------------------------------------------------------------------
virtualbox-install: $(VIRTUALBOX)

virtualbox-uninstall: 
	@if [ -d $(VIRTUALBOX) ]; then make virtualbox-shutdown-uninstall; fi

virtualbox-shutdown-uninstall:
	@echo shutdown... VirtualBox $(VIRTUALBOX_VERSION)
	@vboxmanage list runningvms | sed -E 's/.*\{(.*)\}/\1/' | xargs -L1 -I {} VBoxManage controlvm {} savestate
	@pkill -HUP -f /Applications/VirtualBox.app/Contents/MacOS/VirtualBox || true
	@sudo /Library/Application\ Support/VirtualBox/LaunchDaemons/VirtualBoxStartup.sh  stop
	@echo uninstalling... VirtualBox $(VIRTUALBOX_VERSION)
	@if [ ! -f $(VIRTUALBOX_ARCHIVE) ]; then make $(VIRTUALBOX_ARCHIVE); fi
	hdiutil attach $(VIRTUALBOX_ARCHIVE)
	/Volumes/VirtualBox/VirtualBox_Uninstall.tool
	hdiutil detach /Volumes/VirtualBox
	rm -f $(VIRTUALBOX_ARCHIVE)

$(VIRTUALBOX): 
	@echo installing... VirtualBox $(VIRTUALBOX_VERSION)
	@if [ ! -f $(VIRTUALBOX_ARCHIVE) ]; then make $(VIRTUALBOX_ARCHIVE); fi
	hdiutil attach $(VIRTUALBOX_ARCHIVE)
	sudo installer -pkg /Volumes/VirtualBox/VirtualBox.pkg -target /Volumes/Macintosh\ HD
	hdiutil detach /Volumes/VirtualBox

$(VIRTUALBOX_ARCHIVE): 
	@echo downloading... $(@F)
	@mkdir -p $(@D)
	curl -L -o $@ http://download.virtualbox.org/virtualbox/$(VIRTUALBOX_VERSION)/$(@F)
	openssl dgst -sha256 $@


.PHONY: $(ALL_TARGETS)
