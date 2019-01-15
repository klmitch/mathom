REPO      = github.com/klmitch/mathom

GO        = go
GOFMT     = gofmt
GOIMPORTS = goimports
GOLINT    = golint

SOURCES   = $(shell find . -name .git -prune -o -name \*.go -print)

all: test

format:
ifeq ($(CI_TEST),true)
	@imports=`$(GOIMPORTS) -l $(SOURCES)`; \
	fmts=`$(GOFMT) -l -s $(SOURCES)`; \
	all=`(echo $$imports; echo $$fmts) | sort -u`; \
	if [ "$$all" != "" ]; then \
		echo "Following files need updates:"; \
		echo; \
		echo $$all; \
		exit 1;\
	fi
else
	$(GOIMPORTS) -l -local $(REPO) -w $(SOURCES)
	$(GOFMT) -l -s -w $(SOURCES)
endif

lint:
	$(GOLINT) -set_exit_status ./...

vet:
	$(GO) vet ./...

test: format lint vet
	$(GO) test -coverprofile=cover.out ./...

cover: test
	$(GO) tool cover -html=cover.out -o coverage.html
