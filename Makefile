install_hz:
	go install github.com/cloudwego/hertz/cmd/hz@latest

gen_api:
	hz update -module github.com/eeianux/lmq  --service=lmq --idl=idl/lmq.thrift

run:
	./build.sh
	./output/bootstrap.sh