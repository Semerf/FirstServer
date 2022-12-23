.PHONY: run
run:
	go run cmd/storage/main/main.go
	go run cmd/storage/calculate/calculate.go
	
server:
	go run cmd/storage/main/main.go
calculate:
	go run cmd/storage/calculate/calcmain.go

