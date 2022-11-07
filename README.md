# Luna 🌙

Luna is a reeeaaally tiny, yet expanding, compiler for WebAssembly Text Format, written in Go and built as one of my quest to conquer the WebAssembly dungeon.

(I just wanted to build something like wat2wasm)

<img src="https://i.ibb.co/hdcV1h0/Screenshot-2022-11-01-alle-17-47-06.png" alt="luna" />

It is so tiny that can only make additions lol.

# Why ❓

I've built Luna because I wanted to learn how to build a compiler while learning WebAssembly. 
<strong>So Luna was built for DEMONSTRATION and EDUCATIONAL purposes first.</strong>

The goal of Luna is not to do fancy stuff to replace (in a long distant future) solid tools like <a href="https://webassembly.github.io/wabt/demo/wat2wasm/" target="_blank">wat2wasm</a>, <a href="https://github.com/wasmerio/wasmer" target="_blank">wasmer</a> or others...

The goal of this project is to become a `useful landmark` for anyone approaching WebAssembly and/or for anyone that wants to develop a compiled-to-wasm programming language. 

> I tried to document each section of the code as much as I could (I'm still doing it) with link to resources I've studied while building this, but if you want to improve it, feel free to open issues and pull requests.

# How it works ❓

Followed the amazing articles about the <a href="https://blog.scottlogic.com/2019/05/17/webassembly-compiler.html" target="_blank">Chasm compiler</a> (a WAT compiler written in Typescript for the Chasm language) and a guide to write a <a href="https://www.bitfalter.com/webassembly-compiler-building-a-compiler" target="_blank">WAT compiler</a> in Rust 

- Luna takes a `.wat` file (or string if used in the browser)
- Splits it into tokens `./compiler/tokenizer.go`
- Creates a very simple AST of the tokens `./compiler/parser.go`
- Compiles `./compiler/compiler.go`

# Use it in the browser 🌐

Luna can also be used in the browsers

Demo: https://luna-demo.vercel.app/

The

```bash
make wasm
```

command will build (or update) with TinyGo the `./example/main.wasm` file to be imported in the browser
In the `./example` directory there's a working example of how to do that.

> 💡 - `make update` will simply update/replace the existing main.wasm

> 💡 - Check the console to see the tokenizer and the parser outputs

# Aeon Runtime

Luna also implements a really tiny runtime that can run the exported functions.
Read more about it in `./runtime/README.md`

# Requirements ✋

- Go
- Tinygo
- Make
- syscall/js set up (or simply comment out the startLuna function and the syscall/js import)

# Roadmap

1. <h3>More interactivity</h3>
    Currently Luna supports only the renaming of the exported function and some order scrumbling

2. <h3>More arithmetics</h3>
    Currently Luna supports only addition

3. <h3>Expansion of Wat syntax</h3>

# Contributing

To contribute simply 
- create a branch with the feature or the bug fix
- open pull requests

Luna is by no means finished there are a lot of things that can be implemented and A LOT of things that can be improved. Any suggestions, pull requests, issues or feedback is greatly welcomed!!!

Per aspera ad astra!