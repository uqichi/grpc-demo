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

push-server:
	gcloud auth configure-docker
	docker tag grpc-demo_grpc-server asia.gcr.io/$${GOOGLE_PROJECT_ID}/grpc-demo-server:tls
	docker push asia.gcr.io/$${GOOGLE_PROJECT_ID}/grpc-demo-server:tls

push-client:
	gcloud auth configure-docker
	docker tag grpc-demo_grpc-client asia.gcr.io/$${GOOGLE_PROJECT_ID}/grpc-demo-client:tls
	docker push asia.gcr.io/$${GOOGLE_PROJECT_ID}/grpc-demo-client:tls
