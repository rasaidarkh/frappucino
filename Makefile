build:
	docker build .
compose:
	docker-compose up
run:
	go run cmd/api/main.go