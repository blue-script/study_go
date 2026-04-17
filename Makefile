include .env
export

add_new_user:
	export NEW_USER=YES && \
	go run main.go

server:
	go run main.go