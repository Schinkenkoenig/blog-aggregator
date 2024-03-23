run:
	go build -o bin/blog-agg && ./bin/blog-agg

gen:
	sqlc generate
