ROOT_DIR=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/
RELEASE_VERSION=V2.0

GOOS=linux
CGO_ENABLED=0
DIST_DIR=$(ROOT_DIR)dist/

.PHONY: release
release: dist_dir sender receiver;

.PHONY: release-darwin
release-darwin: darwin release;

.PHONY: dist_dir
dist_dir: ; $(info ======== prepare distribute dir:)
	mkdir -p $(DIST_DIR)
	@rm -rf $(DIST_DIR)*

.PHONY: sender
sender: ; $(info ==============go build sender==============)
	env CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) go build -a -installsuffix cgo -o $(DIST_DIR)sender $(ROOT_DIR)cmd/sender/*.go

.PHONY: receiver
receiver: ; $(info ==============go build receiver==============)
	env CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) go build -a -installsuffix cgo -o $(DIST_DIR)receiver $(ROOT_DIR)cmd/receiver/*.go

.PHONY: docker
docker: release;
	@echo ==============docker build==============
	docker build -t domgoer/msg-pusher:$(RELEASE_VERSION) -f Dockerfile .

.PHONY: sender-docker
sender-docker: dist_dir sender;
	@echo ==============docker build==============
	docker build -t domgoer/msg-pusher-sender:$(RELEASE_VERSION) -f Dockerfile-sender .

.PHONY: receiver-docker
receiver-docker: dist_dir receiver;
	@echo ==============docker build==============
	docker build -t domgoer/msg-pusher-receiver:$(RELEASE_VERSION) -f Dockerfile-receiver .

.PHONY: darwin
darwin:
	$(eval GOOS := darwin)
