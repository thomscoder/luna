// runtime/helpers/reader.js
var WasmReader = class {
  constructor(wasm) {
    this.data = wasm;
    this.pos = 0;
  }
  dword() {
    let prev = this.pos;
    this.pos += 4;
    const arr = this.data.slice(prev, this.pos);
    return arr;
  }
  readBytes(num) {
    let prev = this.pos;
    this.pos += num;
    return this.data.slice(prev, this.pos);
  }
  readByte() {
    let prev = this.pos;
    this.pos += 1;
    return this.data[prev];
  }
};

// runtime/utils/defaults.js
var Opcodes = {
  block: 2,
  loop: 3,
  br: 12,
  br_if: 13,
  end: 11,
  call: 16,
  get_local: 32,
  set_local: 33,
  i32_store_8: 58,
  i32_const: 65,
  f32_const: 67,
  i32_eqz: 69,
  i32_eq: 70,
  f32_eq: 91,
  f32_lt: 93,
  f32_gt: 94,
  i32_and: 113,
  i32_add: 106,
  i32_sub: 107,
  f32_add: 146,
  f32_sub: 147,
  f32_mul: 148,
  f32_div: 149
};
var Section = {
  custom: 0,
  type: 1,
  import: 2,
  func: 3,
  table: 4,
  memory: 5,
  global: 6,
  export: 7,
  code: 10
};
var ExportSection = {
  func: 0,
  table: 1,
  mem: 2,
  global: 3
};

// runtime/utils/errors.js
var RuntimeErrors = {
  ModuleIsEmpty: "Module is empty",
  ModuleTooShort: "Module is too short",
  InvalidMagicHeader: "Invalid magic header",
  InvalidVersionHeader: "Invalid version header",
  InvalidSection: "Invalid section",
  InvalidValueType: "Invalid value type",
  InvalidExportType: "Invalid export type",
  InvalidExportName: "Invalid export name",
  InvalidInstruction: "Invalid instruction",
  InvalidArgumentsNumber: "Invalid number of arguments",
  ExportNotFound: "Export not found"
};
var _RuntimeErrors = RuntimeErrors;

// runtime/utils/types.js
var ValType = {
  i32: 127,
  i64: 126,
  f32: 125,
  f64: 124
};

// runtime/runtime/parser.js
function checkHeader(wasm) {
  if (wasm.data.length < 8) {
    throw new Error(_RuntimeErrors.ModuleTooShort);
  }
  if (wasm.data.length == 8) {
    throw new Error(_RuntimeErrors.ModuleIsEmpty);
  }
  ;
  const magicString = String.fromCharCode(...wasm.readBytes(4));
  if (magicString !== "\0asm") {
    throw new Error(_RuntimeErrors.InvalidMagicHeader);
  }
  const version = wasm.dword();
  if (version.length < 1) {
    throw new Error(_RuntimeErrors.InvalidVersionHeader);
  }
  return true;
}
function parseTypeSection(wasm) {
  const sectionType = wasm.readByte();
  if (sectionType !== Section.type) {
    throw new Error(_RuntimeErrors.InvalidSection);
  }
  function parseValueType(wasm2) {
    const valType = wasm2.readByte();
    switch (valType) {
      case ValType.i32:
        return ValType.i32;
      default:
        throw new Error(_RuntimeErrors.InvalidValueType);
    }
  }
  let sizeOfSection = wasm.readByte();
  const numTypes = wasm.readByte();
  let types = [];
  for (let i = 0; i < numTypes; i++) {
    let func = wasm.readByte();
    let numOfParams = wasm.readByte();
    let params = [];
    for (let j = 0; j < numOfParams; j++) {
      params.push(parseValueType(wasm));
    }
    let numOfResults = wasm.readByte();
    let results = [];
    for (let w = 0; w < numOfResults; w++) {
      results.push(parseValueType(wasm));
    }
    types.push([params, results]);
  }
  return types;
}
function parseFunctionSection(wasm) {
  const isSectionFunc = wasm.readByte();
  if (isSectionFunc !== Section.func) {
    throw new Error(_RuntimeErrors.InvalidSection);
  }
  let sectionSize = wasm.readByte();
  let numberOfFunctions = wasm.readByte();
  let functionSigIndex = [];
  for (let i = 0; i < numberOfFunctions; i++) {
    functionSigIndex.push(wasm.readByte());
  }
  return functionSigIndex;
}
function parseExportSection(wasm) {
  const isExportSection = wasm.readByte();
  if (isExportSection !== Section.export) {
    throw new Error(_RuntimeErrors.InvalidSection);
  }
  let sectionSize = wasm.readByte();
  let numberOfExports = wasm.readByte();
  let exportsArr = [];
  for (let i = 0; i < numberOfExports; i++) {
    const exportNameLength = wasm.readByte();
    const exportName = String.fromCharCode(...wasm.readBytes(exportNameLength));
    if (!!exportName === false)
      throw new Error(_RuntimeErrors.InvalidExportName);
    let zero = wasm.readByte();
    let indexOfExportedFunction = wasm.readByte() == ExportSection.func && ExportSection.func;
    if (typeof indexOfExportedFunction !== "number" && !!indexOfExportedFunction === false) {
      throw new Error(_RuntimeErrors.InvalidExportType);
    }
    exportsArr.push({ name: exportName.replace(/"/g, ""), index: indexOfExportedFunction });
  }
  return exportsArr;
}
function parseCodeSection(wasm) {
  const isCodeSection = wasm.readByte();
  if (isCodeSection !== Section.code) {
    throw new Error(_RuntimeErrors.InvalidSection);
  }
  let sectionSize = wasm.readByte();
  let numberOfFunctions = wasm.readByte();
  let code = [];
  for (let i = 0; i < numberOfFunctions; i++) {
    let funcBodySize = wasm.readByte();
    let numberOfLocals = wasm.readByte();
    let instructions = [];
    let locals = [];
    while (wasm.pos < wasm.data.length) {
      const instruction = wasm.readByte();
      instructions.push(instruction);
      if (instruction == Opcodes.get_local) {
        locals.push(wasm.readByte());
      }
    }
    code.push([locals, instructions]);
  }
  return code;
}

// runtime/runtime/ast.js
function createAST(wasmBinary) {
  const wasm = new WasmReader(wasmBinary);
  checkHeader(wasm);
  const typeSection = parseTypeSection(wasm);
  const functionTypes = parseFunctionSection(wasm);
  const exportSection = parseExportSection(wasm);
  const codeSection = parseCodeSection(wasm);
  const moduleAst = {
    typeSection,
    functionTypes,
    exportSection,
    codeSection
  };
  return moduleAst;
}

// runtime/runtime/processor.js
var Processor = class {
  constructor(func, params) {
    this.func = func;
    this.params = params;
    this.stack = [];
  }
  executeFunc() {
    for (const instruction of this.func[1]) {
      if (instruction == Opcodes.get_local) {
        this.stack.push(this.params[this.func[0].shift()]);
      }
      this.#parseInstruction(instruction);
    }
  }
  #parseInstruction(instruction) {
    let result;
    switch (instruction) {
      case Opcodes.i32_add:
        result = this.stack.reduce((prev, current) => prev + current, 0);
        return this.stack.push(result);
      case Opcodes.i32_sub:
        result = this.stack.reduce((prev, current) => prev - current);
        return this.stack.push(result);
    }
  }
  getResult() {
    return this.stack.pop();
  }
};

// runtime/runtime/invoker.js
function invokeFunction(ast, funcName, params) {
  let exportIndex = ast.exportSection.findIndex((exp) => exp.name === funcName);
  if (exportIndex === -1)
    throw new Error(_RuntimeErrors.ExportNotFound);
  let func = ast.codeSection[exportIndex];
  if (ast.typeSection[exportIndex][0].length !== params.length) {
    throw new Error(_RuntimeErrors.InvalidArgumentsNumber);
  }
  const processor = new Processor(func, params);
  processor.executeFunc();
  return processor.getResult();
}

// runtime/runtime/start.js
var startAeonRuntime = (wasm, ...args) => {
  const ast = createAST(wasm);
  const [funcName, ...rest] = args;
  const params = rest;
  const result = invokeFunction(ast, funcName, params);
  return result;
};
var start_default = startAeonRuntime;
export {
  start_default as default
};
