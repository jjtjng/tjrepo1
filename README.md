To use, obtain key by signing up https://api.macaddress.io

- To build: make build
- To build docker image: make docker
- To run:
	Invoke 'MACADDRESS_IO_API_KEY=<your-key> go run main.go <mac address>'
	Or 'MACADDRESS_IO_API_KEY=<your-key> maclookuptest_mac <mac address>'
	Or 'MACADDRESS_IO_API_KEY=<your-key> maclookuptest_linux <mac address>'
- TO run in docker,
	Invoke 'docker run --env MACADDRESS_IO_API_KEY=<your key> -it maclookuptest
	Then invoke /usr/local/bin/maclookuptest <mac address>
