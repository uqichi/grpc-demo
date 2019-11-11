.PHONY: protoc doc up logs

protoc:
	protoc -I proto --go_out=plugins=grpc:proto proto/*.proto

doc:
	protoc --doc_out=./doc --doc_opt=html,index.html proto/*.proto

up:
	docker-compose up -d --build

logs:
	docker-compose logs -f