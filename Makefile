app_path = cmd/app.go
app_name = update_packages

build:
	go build -o $(app_name) $(app_path)

run:
	./$(app_name)

start: build run