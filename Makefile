.PHONY: test
test:
	go test ./... -race

.PHONY: bench
bench:
	go test -run=^$$ -bench=. -benchmem -cpuprofile cpu.prof

.PHONY: trace
trace:
	go test -run=^$$ -bench=. -trace=trace.out

