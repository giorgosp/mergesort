# Mergesort

This repo contains an unoptimized sequential and an optimized parallel implemention
of mergesort in Go. **Not for production use**.

There are some posts that complement this repo:

- [Sequential and parallel mergesort in Go](https://medium.com/@giopap/sequential-and-parallel-merge-sort-in-go-74881e92a609). Describes the initial (baseline) implementations.
- [Optimizing parallel Merge Sort with Go tools](https://medium.com/@giopap/optimizing-parallel-mergesort-with-golang-tools-ffca702f1e30). Describes the journey of optimising the parallel implementation using the `pprof` and `trace` tools.

The optimization steps are also distinct branches. `master` contains the most optimized implementation.
