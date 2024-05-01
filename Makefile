APPS := sonic
ELASTIC_VERSION := 7.17.0
KIBANA_VERSION := 7.17.0

all: $(APPS)

$(APPS):
	go build --ldflags "-X main.Version=$$(cat VERSION)" -o bin/$@ app.go

test:
	go test ./...

asset:
	cd web && \
		rm -rf ../public/assets && \
		rm -rf ./build && \
		yarn install && \
		./scripts/build.sh && \
		cp -r build ../public/assets

clean:
	rm -rf bin/*

start:
	# Elasticsearch ve Kibana Docker imajlarını başlat
	docker-compose up -d elasticsearch kibana

stop:
	docker-compose down
