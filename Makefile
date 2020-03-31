.PHONY: test
test:
	go test ./... -race

.PHONY: bench
bench:
	go test -run=^$$ -bench=. -benchmem -cpuprofile cpu.prof -count 5

.PHONY: trace
trace:
	go test -run=^$$ -bench=. -trace=trace.out

