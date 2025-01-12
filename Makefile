build:
	docker build .
compose:
	docker-compose up
run:
	go run cmd/myapp/myapp.go