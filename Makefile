cover:
	go test -coverprofile cover.out -covermode count
	go tool cover -html cover.out

bench:
	go test -bench Bench -benchmem
