server:
	go run ./cmd/entry.go
client:
	cd ./web/banana-frontend
	npm start

#build
docker build:
	docker build -t banana .
kuber:
	kubectl apply -f deployment.yaml
	kubectl apply -f service.yaml