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
#   $HOME/.woo/bashrc
#   $HOME/.woo/cache/<downloads...>
#   $HOME/.woo/conf/*rc
#   $HOME/.woo/dist/gradle-<version>
#   $HOME/.woo/dist/groovy-<version>
#   $HOME/.woo/dist/go-<version>
#   $HOME/.woo/dist/packer-<version>
#   $HOME/.woo/dist/vault-<version>
#   $HOME/.woo/repo/<repo.git>
#   /usr/local/opt/vagrant -> /opt/vagrant
#   /Applications/VirtualBox.app
#   /Library/Java/JavaVirtualMachines/jdk...
#
# Utilities Purpose
#   * golang        utilities to build, deploy, manage workspaces and services
#   * gradle        utilities to orchestrate tasks, and used a task DAG
#   * groovy        implement gradle custom plugins and tasks
#   * packer        create vm images to host docker engine nodes
#   * vagrant       manage vm nodes
#   * vault         manage workspace secrets (keys, tokens, passwords)
#   * virtualbox    hypervisor to run docker nodes
# ------------------------------------------------------------------------

WOO_HOME                = $(HOME)/.woo
WOO_BASHRC              = $(WOO_HOME)/bashrc
CACHEDIR                = $(WOO_HOME)/cache
CONFDIR                 = $(WOO_HOME)/conf
DISTDIR                 = $(WOO_HOME)/dist
REPODIR                 = $(WOO_HOME)/repo

GOLANG_WEBSITE          = https://golang.org
GOLANG_DOWNLOADS        = https://golang.org/dl/
GOLANG_VERSION          = 1.9.2
GOLANG_ARCHIVE          = $(CACHEDIR)/go$(GOLANG_VERSION).darwin-amd64.tar.gz
GOLANG_DESTDIR          = $(DISTDIR)/go-$(GOLANG_VERSION)
GOLANG                  = $(GOLANG_DESTDIR)/bin/go

GRADLE_WEBSITE          = https://gradle.org
GRADLE_DOWNLOADS        = https://gradle.org/releases/
GRADLE_VERSION          = 4.2.1
GRADLE_ARCHIVE          = $(CACHEDIR)/gradle-$(GRADLE_VERSION)-bin.zip
GRADLE_DESTDIR          = $(DISTDIR)/gradle-$(GRADLE_VERSION)
GRADLE                  = $(GRADLE_DESTDIR)/bin/gradle

GROOVY_WEBSITE          = http://groovy-lang.org
GROOVY_DOWNLOADS        = http://groovy-lang.org/download.html
GROOVY_VERSION          = 2.4.12
GROOVY_ARCHIVE          = $(CACHEDIR)/apache-groovy-binary-$(GROOVY_VERSION).zip
GROOVY_DESTDIR          = $(DISTDIR)/gradle-$(GROOVY_VERSION)
GROOVY                  = $(GROOVY_DESTDIR)/bin/gradle

JDK_WEBSITE             = http://www.oracle.com/technetwork/java/index.html
JDK_DOWNLOADS           = http://www.oracle.com/technetwork/java/javase/downloads/index.html
JDK_VERSION             = 9.0.1
JDK_ARCHIVE             = $(CACHEDIR)/jdk-$(JDK_VERSION)_osx-x64_bin.dmg
JDK                     = /Library/Java/JavaVirtualMachines/jdk-$(JDK_VERSION).jdk

NVM_WEBSITE             = https://github.com/creationix/nvm
NVM_DOWNLOADS           = https://github.com/creationix/nvm#git-install
NVM_VERSION             = v0.33.8
NVM_REPO                = $(REPODIR)/nvm
NVM_BASHRC              = $(CONFDIR)/nvmrc
NVM                     = $(NVM_REPO)/nvm.sh

NODE_WEBSITE            = https://nodejs.org
NODE_DOWNLOADS          = https://nodejs.org/en/download/
NODE_VERSION            = v8.9.3
NODE                    = $(NVM_REPO)/versions/node/$(NODE_VERSION)/bin/node


PACKER_WEBSITE          = https://www.packer.io
PACKER_DOWNLOADS        = https://www.packer.io/downloads.html
PACKER_VERSION          = 1.1.2
PACKER_ARCHIVE          = $(CACHEDIR)/packer_$(PACKER_VERSION)_darwin_amd64.zip
PACKER_DESTDIR          = $(DISTDIR)/packer-$(PACKER_VERSION)
PACKER                  = $(PACKER_DESTDIR)/bin/packer

VAGRANT_WEBSITE         = https://www.vagrantup.com
VAGRANT_DOWNLOADS       = https://www.vagrantup.com/downloads.html
VAGRANT_VERSION         = 2.0.1
VAGRANT_ARCHIVE         = $(CACHEDIR)/vagrant_$(VAGRANT_VERSION)_x86_64.dmg
VAGRANT_DESTDIR         = /opt/vagrant
VAGRANT                 = $(VAGRANT_DESTDIR)/vagrant

VAULT_WEBSITE           = https://www.vaultproject.io
VAULT_DOWNLOADS         = https://www.vaultproject.io/downloads.html
VAULT_VERSION           = 0.9.0
VAULT_ARCHIVE           = $(CACHEDIR)/vault_$(VAULT_VERSION)_darwin_amd64.zip
VAULT_DESTDIR           = $(DISTDIR)/vault-$(VAULT_VERSION)
VAULT                   = $(VAULT_DESTDIR)/bin/vault

VIRTUALBOX_WEBSITE      = https://www.virtualbox.org
VIRTUALBOX_DOWNLOADS    = https://www.virtualbox.org/wiki/Downloads
VIRTUALBOX_VERSION      = 5.2.2
VIRTUALBOX_ARCHIVE      = $(CACHEDIR)/VirtualBox-$(VIRTUALBOX_VERSION)-119230-OSX.dmg
VIRTUALBOX              = /Applications/VirtualBox.app

BASE_TARGETS           += golang
BASE_TARGETS           += gradle
BASE_TARGETS           += groovy
BASE_TARGETS           += packer
BASE_TARGETS           += jdk
BASE_TARGETS           += node
BASE_TARGETS           += vagrant
BASE_TARGETS           += vault
BASE_TARGETS           += virtualbox

INSTALL_TARGETS         = $(BASE_TARGETS:%=%-install)
UNINSTALL_TARGETS       = $(BASE_TARGETS:%=%-uninstall)
ALL_TARGETS             = $(INSTALL_TARGETS) $(UNINSTALL_TARGETS)

# ------------------------------------------------------------------------
# Installation Driver Targets
# ------------------------------------------------------------------------
help:; @for i in $(BASE_TARGETS); do make $$i; echo; done
install: $(INSTALL_TARGETS)
uninstall: $(UNINSTALL_TARGETS)

# ------------------------------------------------------------------------
# Golang
# ------------------------------------------------------------------------
golang:
	@echo "Component : Go"
	@echo "Version   : $(GOLANG_VERSION)"
	@echo "Cache     : $(GOLANG_ARCHIVE)"
	@echo "Install   : $(GOLANG_DESTDIR)"
	@echo "Website   : $(GOLANG_WEBSITE)"
	@echo "Downloads : $(GOLANG_DOWNLOADS)"
	@echo "Targets   : golang-install golang-uninstall"

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
	tar -C $(GOLANG_DESTDIR) -xvf $(GOLANG_ARCHIVE) --strip-components 1

# ------------------------------------------------------------------------
# Gradle 
# ------------------------------------------------------------------------
gradle:
	@echo "Component : Gradle"
	@echo "Version   : $(GRADLE_VERSION)"
	@echo "Cache     : $(GRADLE_ARCHIVE)"
	@echo "Install   : $(GRADLE_DESTDIR)"
	@echo "Website   : $(GRADLE_WEBSITE)"
	@echo "Downloads : $(GRADLE_DOWNLOADS)"
	@echo "Targets   : gradle-install gradle-uninstall"

gradle-install: $(GRADLE)

gradle-uninstall: 
	@echo removing... Gradle $(GRADLE_VERSION)
	rm -rf $(GRADLE_DESTDIR)
	rm -f $(GRADLE_ARCHIVE)
	rm -f $(GRADLE_ARCHIVE).sha256

$(GRADLE_ARCHIVE): 
	@echo downloading... $(@F)
	@mkdir -p $(@D)
	curl -L -o $@ https://services.gradle.org/distributions/$(@F)
	curl -fsSL -o $(@).sha256 https://services.gradle.org/distributions/$(@F).sha256
	openssl dgst -sha256 $@

$(GRADLE): 
	@echo installing... Gradle $(GRADLE_VERSION)
	@if [ ! -f $(GRADLE_ARCHIVE) ]; then make $(GRADLE_ARCHIVE); fi
	@mkdir -p $(DISTDIR)
	cd $(DISTDIR) && unzip $(GRADLE_ARCHIVE)

# ------------------------------------------------------------------------
# Groovy 
# ------------------------------------------------------------------------
groovy:
	@echo "Component : Groovy"
	@echo "Version   : $(GROOVY_VERSION)"
	@echo "Cache     : $(GROOVY_ARCHIVE)"
	@echo "Install   : $(GROOVY_DESTDIR)"
	@echo "Website   : $(GROOVY_WEBSITE)"
	@echo "Downloads : $(GROOVY_DOWNLOADS)"
	@echo "Targets   : groovy-install groovy-uninstall"

groovy-install: $(GROOVY)

groovy-uninstall: 
	@echo removing... Gradle $(GROOVY_VERSION)
	rm -rf $(GROOVY_DESTDIR)
	rm -f $(GROOVY_ARCHIVE)

$(GROOVY_ARCHIVE): 
	@echo downloading... $(@F)
	@mkdir -p $(@D)
	curl -L -o $@ https://dl.bintray.com/groovy/maven/$(@F)
	openssl dgst -sha256 $@

$(GROOVY): 
	@echo installing... Groovy $(GROOVY_VERSION)
	@if [ ! -f $(GROOVY_ARCHIVE) ]; then make $(GROOVY_ARCHIVE); fi
	@mkdir -p $(DISTDIR)
	cd $(DISTDIR) && unzip $(GROOVY_ARCHIVE)

# ------------------------------------------------------------------------
# JDK
# ------------------------------------------------------------------------
jdk:
	@echo "Component : JDK"
	@echo "Version   : $(JDK_VERSION)"
	@echo "Cache     : $(JDK_ARCHIVE)"
	@echo "Install   : $(JDK)"
	@echo "Website   : $(JDK_WEBSITE)"
	@echo "Downloads : $(JDK_DOWNLOADS)"
	@echo "Targets   : jdk-install jdk-uninstall"

jdk-install: $(JDK)

jdk-uninstall:
	@echo uninstalling... JDK $(JDK_VERSION)
	sudo rm -rf /Library/Java/JavaVirtualMachines/jdk-9.0.1.jdk
	sudo rm -rf /Library/Internet\ Plug-Ins/JavaAppletPlugin.plugin
	sudo rm -rf /Library/PreferencePanes/JavaControlPanel.prefPane
	rm -f $(JDK_ARCHIVE)

$(JDK): 
	@echo installing... JDK $(JDK_VERSION)
	@if [ ! -f $(JDK_ARCHIVE) ]; then make $(JDK_ARCHIVE); fi
	hdiutil attach $(JDK_ARCHIVE)
	sudo installer -pkg /Volumes/JDK\ $(JDK_VERSION)/JDK\ $(JDK_VERSION).pkg -target /Volumes/Macintosh\ HD
	hdiutil detach /Volumes/JDK\ $(JDK_VERSION)

$(JDK_ARCHIVE): 
	@echo downloading... $(@F)
	@mkdir -p $(@D)
	curl -jkL -o $@  -H "Cookie: oraclelicense=accept-securebackup-cookie"  http://download.oracle.com/otn-pub/java/jdk/$(JDK_VERSION)+11/jdk-$(JDK_VERSION)_osx-x64_bin.dmg
	openssl dgst -sha256 $@

jdk-checksum:
	curl -fsSL https://www.oracle.com/webfolder/s/digest/9-0-1checksum.html | awk 'f{print;f=0} /jdk-9.0.1_osx-x64_bin/{f=1}'

# ------------------------------------------------------------------------
# Node
# ------------------------------------------------------------------------
node:
	@echo "Component : Node"
	@echo "Version   : $(NODE_VERSION)"
	@echo "Install   : $(NODE)"
	@echo "Website   : $(NODE_WEBSITE)"
	@echo "Downloads : $(NODE_DOWNLOADS)"
	@echo "Targets   : node-install node-uninstall"

node-install: $(NVM) $(NVM_BASHRC) $(NODE)

node-uninstall: 
	@echo uninstalling... NVM $(NVM_VERSION)
	rm -rf $(NVM_REPO)
	rm -rf $(NVM_BASHRC)

$(NODE):
	@echo install... Node $(NODE_VERSION)
	(source $(NVM_BASHRC) &&  nvm install $(NODE_VERSION) && nvm alias default $(NODE_VERSION))

$(NVM_BASHRC):
	@echo initializing... NVM $(NVM_VERSION) $@
	@echo "export NVM_DIR=$(NVM_REPO)" > $@
	@echo '[ -s "$$NVM_DIR/nvm.sh" ] && source "$$NVM_DIR/nvm.sh"' >> $@
	@echo '[ -s "$$NVM_DIR/bash_completion" ] && source "$$NVM_DIR/bash_completion"' >> $@

$(NVM_REPO):
	@mkdir -p $(@D)
	cd $(@D) && git clone https://github.com/creationix/nvm.git 

$(NVM):
	@echo installing... NVM $(NVM_VERSION)
	@if [ ! -f $(NVM) ]; then make $(NVM_REPO); fi
	cd $(@D) && git checkout $(NVM_VERSION)
	@if [ ! -f $(NVM_BASHRC) ]; then make $(NVM_BASHRC); fi

# ------------------------------------------------------------------------
# Packer
# ------------------------------------------------------------------------
packer:
	@echo "Component : Packer"
	@echo "Version   : $(PACKER_VERSION)"
	@echo "Cache     : $(PACKER_ARCHIVE)"
	@echo "Install   : $(PACKER_DESTDIR)"
	@echo "Website   : $(PACKER_WEBSITE)"
	@echo "Downloads : $(PACKER_DOWNLOADS)"
	@echo "Targets   : packer-install packer-uninstall"

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
# Ruby 
# ------------------------------------------------------------------------
# ruby website https://www.ruby-lang.org/en/
# rbenv website https://github.com/rbenv/rbenv
#
# rbenv install
#   brew install rbenv
#
# rbenv initialize and validate rbenv
#   rbenv init
#   curl -fsSL https://github.com/rbenv/rbenv-installer/raw/master/bin/rbenv-doctor | bash
#
# list available ruby versions
#   rbenv install -l
#
# list installed versions
#   rbenv versions
#
# install a ruby version
#   rbenv install 2.3.5
#
# set a ruby version as default
#   rbenv global 2.3.5

# ------------------------------------------------------------------------
# Traefik
# ------------------------------------------------------------------------
# website https://www.vaultproject.io
# downloads https://www.vaultproject.io/downloads.html
#
# download
#   cd .woo/cache
#   curl -L -O https://github.com/containous/traefik/releases/download/v1.4.4/traefik_darwin-amd64
#   openssl dgst -sha256 traefik_darwin-amd64
#
# install
#   mkdir .woo/dist/traefik-1.1.4
#   cd .woo/dist/traefik-1.1.4
#   cp ../../cache/traefik_darwin-amd64 traefik
#
# uninstall
#   rm -rf .woo/dist/traefik-1.1.4

# ------------------------------------------------------------------------
# Vagrant
# ------------------------------------------------------------------------
vagrant:
	@echo "Component : Vagrant"
	@echo "Version   : $(VAGRANT_VERSION)"
	@echo "Cache     : $(VAGRANT_ARCHIVE)"
	@echo "Install   : $(VAGRANT_DESTDIR)"
	@echo "Website   : $(VAGRANT_WEBSITE)"
	@echo "Downloads : $(VAGRANT_DOWNLOADS)"
	@echo "Targets   : vagrant-install vagrant-uninstall"

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
vault:
	@echo "Component : Vault"
	@echo "Version   : $(VAULT_VERSION)"
	@echo "Cache     : $(VAULT_ARCHIVE)"
	@echo "Install   : $(VAULT_DESTDIR)"
	@echo "Website   : $(VAULT_WEBSITE)"
	@echo "Downloads : $(VAULT_DOWNLOADS)"
	@echo "Targets   : vault-install vault-uninstall"

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
virtualbox:
	@echo "Component : VirtualBox"
	@echo "Version   : $(VIRTUALBOX_VERSION)"
	@echo "Cache     : $(VIRTUALBOX_ARCHIVE)"
	@echo "Install   : $(VIRTUALBOX)"
	@echo "Website   : $(VIRTUALBOX_WEBSITE)"
	@echo "Downloads : $(VIRTUALBOX_DOWNLOADS)"
	@echo "Targets   : virtualbox-install virtualbox-uninstall"

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

# ------------------------------------------------------------------------
# Woorc
# ------------------------------------------------------------------------
woo-bashrc: $(WOO_BASHRC)

$(WOO_BASHRC):
	@echo initialize... $@
	@echo 'for i in `find $$WOO_HOME/conf -name \*rc` ; do echo $$i; source $$i; done' > $@


.PHONY: $(ALL_TARGETS)
