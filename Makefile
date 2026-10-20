.PHONY run:
	docker build -t forum .
	docker run -d -p 8080:8080 --name forum-app forum

.PHONY rebuild:
	docker build -t forum .
	docker rm -f forum-app
	docker run -d -p 8080:8080 --name forum-app forum
