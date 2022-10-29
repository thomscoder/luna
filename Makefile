compile:
	go build -o ./dist/luna main.go && \
	cd dist && ./luna