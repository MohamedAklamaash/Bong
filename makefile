networkrun:
	go run packetsniffing/main.go

networkbuild:
	go build -o bin/main packetsniffing/main.go

portrun:
	go run port-scanning/main.go