
.PHONY:default
default:
	@echo "No default action"
	@echo ""
	@echo "    make test    - run all tests"
	@echo "    make new-day - create a new day solution"
	@echo ""

.PHONY:test
test:
	go test -count 1 ./...

.PNONY:new-day
new-day:
	@read -p "Enter day number: " num; \
	go run -tags gen ./tpl/. -day $$num
