## set up go-mqtt

### set up go

	cd /root/download
	wget https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
	tar -zvxf go1.6.linux-amd64.tar.gz
	mv go ~/go
	vim ~/.bashrc
	/*
	export GOROOT=$HOME/go
	export GOPATH=/workspace/github/go-mqtt
	export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
	*/
	source ~/.bashrc
	

### set up surgemq

	mkdir -p /workspace/github/go-mqtt
	cd /workspace/github/go-mqtt
	mkdir bin kpg src
	go get github.com/surgemq/surgemq
	go get ./...


### run surgemq

	cd /workspace/github/go-mqtt/src/github.com/surgemq/surgemq/examples/surgemq/
	go run surgemq.go
	// test by emqtt-benchmark
	cd /root/download
	wget https://packages.erlang-solutions.com/erlang/esl-erlang/FLAVOUR_1_general/esl-erlang_18.2-1~centos~6_amd64.rpm
	yum -y install esl-erlang_18.2-1~centos~6_amd64.rpm
	cd /workspace/github
	git clone https://github.com/emqtt/emqtt_benchmark.git
	cd emqtt_benchmark
	./emqtt_bench_sub -h 192.168.1.249 -p 1883 -c 3000 -t bench_mark -q 1 -u xxx -P xxx -C
	./emqtt_bench_pub -h 192.168.1.249 -p 1883 -c 2 -u xxx -P xxx -t bench_mark -s 10 -q 1


