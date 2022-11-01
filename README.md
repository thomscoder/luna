# Luna

Luna is a reeeaaally minimal compiler for WebAssembly Text Format, written in Go and built as one of my quest to conquer the WebAssembly dungeon.

It is so tiny that can only make additions lol.

I've built Luna because I wanted to learn how to build a compiler and I would like for it to become on.
The goal of Luna is not to replace solid tools like <a href target="_blank">wat2wasm</a>, <a href target="_blank">wasmer</a> 
or to do fancy stuff, the goal of this project is to keep it as minimal as I can yet to be useful for anyone approaching WebAssembly.

I tried to document each section of the code as much as I could (I'm still doing it), but if you want to improve it, feel free to open pull requests.

Luna is by no means finished, while I'll try to keep it as simple and tiny as possible (to ease the reading and learning), there are some things that can be implemented and A LOT of things that can be improved.


# How it works

Followed the amazing article for the <a href>chasm compiler</a>

- Luna takes a `.wat` file (or string if used in the browser)
- Splits it into tokens
- Creates a very simple AST of the tokens
- Compiles

# Use it in the browser

Luna can also be used in the browser

The

```bash
make wasm
```
or

```bash
make update
```
commands will build (or update) with TinyGo the `./example/main.wasm` file to be imported in the browser

In the `./example` directory there's a working example of how to do that.

# Requirements

- Go
- Tinygo
- Make
( - syscall/js )