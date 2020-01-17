build:
	GOOS=linux go build -o maclookuptest_linux main.go
	GOOS=darwin go build -o maclookuptest_mac main.go

docker:
	docker build -t maclookuptest -f Dockerfile .

run:
	echo "Invoke 'MACADDRESS_IO_API_KEY=<your-key> go run main.go <mac address>'"
	echo "Or 'MACADDRESS_IO_API_KEY=<your-key> maclookuptest_mac <mac address>'"
	echo "Or 'MACADDRESS_IO_API_KEY=<your-key> maclookuptest_linux <mac address>'"

rundocker:
	echo "Invoke 'docker run --env MACADDRESS_IO_API_KEY=<your key> -it maclookuptest"
	echo "Then invoke /usr/local/bin/maclookuptest <mac address>"
