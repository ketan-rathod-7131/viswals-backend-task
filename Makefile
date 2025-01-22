build-producer: 
	docker build -f build/Dockerfile-producer -t producer-image:latest .
build-consumer: 
	docker build -f build/Dockerfile-consumer -t consumer-image:latest .
deploy-dev:
	docker compose up --build;