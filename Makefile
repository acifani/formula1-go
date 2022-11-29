f1go:
	go build ./cmd/f1go

run:
	go run ./cmd/f1go

clean:
	rm -rf f1go

demo: f1go
	vhs < demo.tape