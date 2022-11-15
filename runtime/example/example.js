import createAeonRuntime from "../runtime/start.js";

// Compile the ./main.wat file with https://luna-demo.vercel.app
// Get the hex dump from either the middle div or the console

// The below example
// - exports aeonAddition
// - takes two parameters (i32)
// - performs an addition (i32)
const wasmBinary = new Uint8Array([0, 97, 115, 109, 1, 0, 0, 0, 1, 7, 1, 96, 2, 127, 127, 1, 127, 3, 2, 1, 0, 7, 18, 1, 14, 34, 97, 101, 111, 110, 65, 100, 100, 105, 116, 105, 111, 110, 34, 0, 0, 10, 9, 1, 7, 0, 32, 0, 32, 1, 106, 11]);

const runtime = createAeonRuntime(wasmBinary)
const result = runtime("aeonAddition", 2, 3);

console.log("Result", result) // prints 5
