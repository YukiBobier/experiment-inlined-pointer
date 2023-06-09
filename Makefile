.PHONY: benchmark benchmark-no-inline benchmark-show-inlining benchmark-comparison benchmark-print-optimizations

benchmark:
	@echo "##### Benchmarking... #####"
	@go test -bench=. -benchmem

benchmark-no-inline:
	@echo "##### Benchmarking without inlining... #####"
	@go test -bench=. -benchmem -gcflags "-l"

benchmark-comparison: benchmark benchmark-no-inline

benchmark-print-optimizations:
	@echo "##### Printing optimizations... #####"
	@go test -bench=. -benchmem -gcflags "-m"
