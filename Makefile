compile:
	go build -o ./dist/luna main.go && \
	cd dist && ./luna

wasm:
	tinygo build -o ./example/main.wasm -target wasm ./main.go

update: 
	rm ./example/main.wasm && tinygo build -o ./example/main.wasm -target wasm ./main.go