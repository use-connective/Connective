run:
	@mkdir -p bin
	@go build -o bin/connective cmd/*
	@./bin/connective

expose:
	@cloudflared tunnel --url http://localhost:8082
