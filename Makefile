ROOT=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))
GO=$(shell which go)


all: build
	$(ROOT)/fodcmd/fodcmd -domains=$(ROOT)/domains -foxyproxy=$(ROOT)/foxyproxy-patterns.json -pac=$(ROOT)/OmegaProfile.pac -proxifier=$(ROOT)/FOD.ppx

build:
	cd $(ROOT)/fodcmd && $(GO) build .

foxyproxy: build
	$(ROOT)/fodcmd/fodcmd -domains=$(ROOT)/domains -foxyproxy=$(ROOT)/foxyproxy-patterns.json

pac: build
	$(ROOT)/fodcmd/fodcmd -domains=$(ROOT)/domains -pac=$(ROOT)/OmegaProfile.pac

proxifier: build
	$(ROOT)/fodcmd/fodcmd -domains=$(ROOT)/domains -proxifier=$(ROOT)/FOD.ppx

