docker-build:
	cd auth-service && make docker-build
	cd user-service && make docker-build
	cd movie-service && make docker-build
	cd docker && docker-compose up -d

docker-run:
	cd docker && docker-compose up -d
