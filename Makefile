.PHONY: protoc doc up logs push-gcr

protoc:
	protoc -I proto --go_out=plugins=grpc:proto proto/*.proto

doc:
	protoc --doc_out=./doc --doc_opt=html,index.html proto/*.proto

up:
	docker-compose up -d

upbuild:
	docker-compose up -d --build

logs:
	docker-compose logs -f

push-gcr:
	gcloud auth configure-docker
	docker tag grpc-demo_grpc-server asia.gcr.io/$${GOOGLE_PROJECT_ID}/grpc-demo-server
	docker push asia.gcr.io/$${GOOGLE_PROJECT_ID}/grpc-demo-server