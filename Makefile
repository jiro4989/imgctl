NAME := tkimgutil
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -ldflags="-s -w \
	-X \"main.Name=$(NAME)\" \
	-X \"main.Version=$(VERSION)\" \
	-X \"main.Revision=$(REVISION)\" \
	-extldflags \"-static\""

SRCS := $(shell find . -type f -name '*.go')

MAIN_FILES := main.go commands.go

# 配布物に含めるファイル
COPY_FILES := README.md CHANGELOG.md

# 配布物の出力先
DIST_DIR   := dist/$(VERSION)

.PHONY: run
run:
	go run ${MAIN_FILES} ${ARGS}

.PHONY: run-all
run-all:
	go run ${MAIN_FILES} generate | \
		go run ${MAIN_FILES} scale -s 50 | \
		go run ${MAIN_FILES} trim -x 100 -y 290 | \
		sort | \
		go run ${MAIN_FILES} paste

.PHONY: run-find
run-find: clean
	go run ${MAIN_FILES} generate
	go run ${MAIN_FILES} find -d dist/generate | sort | \
		go run ${MAIN_FILES} scale -s 50 | \
		go run ${MAIN_FILES} trim -x 100 -y 290 | \
		sort | \
		go run ${MAIN_FILES} paste

.PHONY: build
build: ${SRCS}
	mkdir -p bin
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) \
		-o bin/$(NAME) .

.PHONY: install
install: build
	go install ${LDFLAGS}

.PHONY: cross-build
cross-build: var-check
	-rm -rf $(DIST_DIR)
	bash ./script/cross-build.sh $(NAME) $(VERSION) $(LDFLAGS) $(MAIN_FILES)

.PHONY: archive
archive: cross-build
	ls -d $(DIST_DIR)/* | while read -r d; do cp $(COPY_FILES) $$d/; done
	bash ./script/arch.sh $(DIST_DIR)

.PHONY: release
release: archive
	ghr $(VERSION) $(DIST_DIR)/

.PHONY: setup
setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/tcnksm/ghr
	dep ensure

.PHONY: clean
clean:
	-rm -rf dist/generate
	-rm -rf dist/paste
	-rm -rf dist/scale
	-rm -rf dist/trim

.PHONY: var-check
var-check:
	if [ "${VERSION}" = "" ]; then echo "VERSIONは必須です。"; exit 1; fi
