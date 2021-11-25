
test:
	go test -count=1 -v ./...

doc:
	godoc -http=localhost:6060 -play

