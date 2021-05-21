VERSION		= $(shell /usr/bin/git describe --always)
BUILD_DATE 	= $(shell date -u +%Y-%m-%dT%T:%Z)
HOSTNAME 	= $(shell hostname)

CONFIG	?= config.yaml

GO_LDFLAGS  = -X main.version=$(if $(VERSION),$(VERSION),unknown)
GO_LDFLAGS += -X main.buildDate=$(if $(BUILD_DATE),$(BUILD_DATE),unknown)
GO_LDFLAGS += -X main.buildHost=$(if $(HOSTNAME),$(HOSTNAME),unknown)

build:
	@echo "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	@echo "VERSION ......: $(VERSION)"
	@echo "BUILD HOST ...: $(HOSTNAME)"
	@echo "BUILD DATE ...: $(BUILD_DATE)"
	@echo "GO VERSION ...: $(word 3, $(shell go version))"
	@echo "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	go build -ldflags "$(GO_LDFLAGS)" .

run:
	go run -ldflags "$(GO_LDFLAGS)" . -config $(CONFIG)

test:
	echo "Not implemented"
