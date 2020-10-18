build: 
	go get -d
	go build -o main .


# using build to find compilation errors fast.
# tries to build the application locally and then properly within the container environment.
containers: build
	docker-compose build --force-rm --no-cache 
	docker-compose up -d
	docker ps
	sleep 5
	docker logs my-service

logs: 
	docker logs my-service

stop_containers:
	docker stop my-service
	docker stop mariadb

clean: stop_containers
	rm main
	docker rm my-service
	docker rm mariadb
	docker system prune