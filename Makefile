GOPATH=$(shell go env GOPATH)
PRJ_DIR=$(shell pwd)

.PHONY: print_env
print_env:
	@echo $(GOPATH)
	@echo $(PRJ_DIR)

.PHONY: start
start:
	docker compose up

.PHONY: stop
stop:
	docker compose down
