
.PHONY build: 
	docker build -t forum .
	-docker rm -f forum-app 
	docker run -d -p 8080:8080 --name forum-app forum
	docker image prune -f