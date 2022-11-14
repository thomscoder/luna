import createAeonRuntime from "./runtime/start.js";

const wasmBinary = new Uint8Array([0, 97, 115, 109, 1, 0, 0, 0, 1, 8, 1, 96, 3, 127, 127, 127, 1, 127, 3, 2, 1, 0, 7, 16, 1, 12, 34, 97, 100, 100, 78, 117, 109, 98, 101, 114, 115, 34, 0, 0, 10, 9, 1, 7, 0, 32, 0, 32, 1, 32, 2, 106, 11]);

const n1 = 8;
const n2 = 20;
const n3 = 23;

const runtime = createAeonRuntime(wasmBinary);
const result = runtime("addNumbers", n1, n2, n3);
console.log(`${n1} + ${n2} + ${n3} =`, result)
