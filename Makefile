
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
	folder="./day$$num"; \
	if test -d "$$folder"; \
	then echo Folder already exists; \
	else echo "Creating $$folder"; \
		mkdir -p "$$folder";  \
		touch "$$folder/sol.go"; \
		echo "package day$$num\n\n\
func Part1(lines []string) int{\n\
\t//todo: write code here\n\
\treturn -1\n\
}\n" >> "./day$$num/sol.go"; \
	fi

