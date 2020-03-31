.PHONY: test
test:
	go test ./... -race

.PHONY: bench
bench:
	go test ./parallel -run=^$$ -bench=. -benchmem -cpuprofile cpu.prof -count 5

.PHONY: trace
trace:
	go test ./parallel -run=^$$ -bench=. -trace=trace.out

