.PHONY: protoc doc

protoc:
	protoc -I proto --go_out=plugins=grpc:proto proto/*.proto

doc:
	protoc --doc_out=./doc --doc_opt=html,index.html proto/*.proto