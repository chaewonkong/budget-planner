# github.com/chaewonkong/budget-planner

gen:
	goa gen github.com/chaewonkong/budget-planner/design/ledger --output gen/ledger
	goa example github.com/chaewonkong/budget-planner/design/ledger
.PHONY: gen


clean:
	rm -rf cmd gen bin
	rm -rf ./*.go
.PHONY: clean

ref:
	@go mod tidy -v
	@go mod vendor -v