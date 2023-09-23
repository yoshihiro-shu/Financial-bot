GOPATH=$(shell go env GOPATH)
PRJ_DIR=$(shell pwd)

.PHONY: print_env
print_env:
	@echo $(GOPATH)
	@echo $(PRJ_DIR)

.PHONY: start
start:
	sh ./script/docker-compose-up.sh

.PHONY: stop
stop:
	sh ./script/docker-compose-down.sh

.PHONY: gen-orm
gen-orm:
	sh ./script/gen-orm.sh

.PHONY: gen-migration
gen-migration:
	sh ./script/gen-migration.sh

