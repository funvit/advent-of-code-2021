
.PHONY:default
default:
	@echo "No default action"

.PHONY:test
test:
	go test -count 1 ./...
