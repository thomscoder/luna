import { createAST } from "./ast.js";
import { invokeFunction } from "./invoker.js";

const startAeonRuntime = (wasm, ...args) => {
    const ast = createAST(wasm);
    const [funcName, ...rest] = args;
    const params = rest;

    const result = invokeFunction(ast, funcName, params);
    return result;
}

export default startAeonRuntime;