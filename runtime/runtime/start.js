import { createAST } from "./ast.js";
import { invokeFunction } from "./invoker.js";

/**
  @param wasm - A wasm binary.
*/
const createAeonRuntime = (wasm) => {
  const ast = createAST(wasm);
  return (funcName, ...params) => invokeFunction(ast, funcName, params);
}

export default createAeonRuntime;
