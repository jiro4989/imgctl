BIN_FILE := tkimgutil
MAIN_FILES := main.go version.go commands.go

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

.PHONY: build
build:
	go build .

.PHONY: setup
setup:
	dep ensure

.PHONY: clean
clean:
	-rm -rf dist/*
	-rm ${BIN_FILE}
