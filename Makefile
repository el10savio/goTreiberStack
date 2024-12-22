
fmt:
	@echo "go fmt goTreiberStack"
	go fmt ./...

vet:
	@echo "go vet goTreiberStack"
	go vet ./...

lint:
	@echo "go lint goTreiberStack"
	golint ./...

test:
	@echo "Testing goTreiberStack"
	go test -v -race --cover ./...

bench:
	@echo "Running benchmarks on goTreiberStack"
	go test -bench=. -count=5 ./...

codespell:
	@echo "checking app spellings"
	codespell
