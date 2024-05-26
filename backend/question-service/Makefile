build-go:
	go build -o build/quest cmd/question/main.go
run:
	go run cmd/question/main.go
clear:
	 docker ps -a --format "{{.ID}} {{.Names}}" | grep "question-service" | awk '{print $1}' | xargs docker rm -f
