run:
	go run main.go
build:
	docker-compose build
up:
	docker-compose up
upb:
	docker-compose up --build
upd:
	docker-compose up -d
down:
	docker-compose down --remove-orphans
downv:
	docker-compose down --remove-orphans --volumes
pgstop:
	sudo service postgresql stop
pgstart:
	sudo service postgresql start
