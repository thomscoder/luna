compile:
	go build -o ./dist/hali main.go && \
	cd dist && ./hali