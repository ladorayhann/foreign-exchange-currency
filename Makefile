build:
	@echo -n "$(NOW) building app... "
	@go build -o mainapp
	@echo "done"

run:
	@echo "$(NOW) starting app... "
	@./mainapp

quick: build run