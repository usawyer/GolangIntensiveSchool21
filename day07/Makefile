ex00:
	go test -v ./test/ex00/...

ex01:
	go test -bench=. test/ex01/minCoinsBenchmark_test.go -cpuprofile=cpu.prof
	go tool pprof -top -cum cpu.prof > top10.txt
	rm -rf benchmark.test cpu.prof

run_godoc:
	go install golang.org/x/tools/cmd/godoc@latest
	~/go/bin/godoc -http=:8080

ex02:
	@mkdir -p docs
	-wget --mirror --convert-links --adjust-extension --page-requisites --no-parent --no-host-directories -P docs/ http://localhost:8080/pkg/day07/internal
	zip -r docs/docs.zip docs
	rm -rf docs/lib docs/pkg

unzip:
	unzip docs/docs.zip

linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

clean:
	rm -rf docs top10.txt
