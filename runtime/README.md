# Aeon Runtime

Aeon will be a separate project (repository) and once done it'll get imported in Luna.

Aeon is a really minimal WASM runtime built for demonstration and educational purposes.

It is so tiny that can only handle `ì32` functions.

# How?

The runtime will be divide in three parts:

- <strong>A parser</strong> for the WASM binary that will output an AST (Abstract Syntax Tree). The parser will run a check against all the sections

- <strong>An executor</strong> of the code in the WASM binary that will implement a stack 
> remember WebAssembly's stack machine concept

- <strong>A processor</strong> where the instructions (e.g. `i32.add`) will take place

# Contributing

Aeon is still under development, but if you want to contribute to either one of these two projects, feel free to share your thoughts and code!!!

Any feedback is greatly appreciated.
