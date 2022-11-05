import { Section, ExportSection, Opcodes } from "../utils/defaults.js";
import { RuntimeErrors } from "../utils/errors.js";
import { ValType } from "../utils/types.js";

// I'm creating a very basic AST https://en.wikipedia.org/wiki/Abstract_syntax_tree
// from the wasm binary
// If you do not come from Luna https://github.com/thomscoder/luna,
// our final wasm binary resembles this

// 0000000: 0061 736d                                 ; WASM_BINARY_MAGIC
// 0000004: 0100 0000                                 ; WASM_BINARY_VERSION
// ; section "Type" (1)
// 0000008: 01                                        ; section code
// 0000009: 00                                        ; section size (guess)
// 000000a: 01                                        ; num types
// ; func type 0
// 000000b: 60                                        ; func
// 000000c: 02                                        ; num params
// 000000d: 7f                                        ; i32
// 000000e: 7f                                        ; i32
// 000000f: 01                                        ; num results
// 0000010: 7f                                        ; i32
// 0000009: 07                                        ; FIXUP section size
// ; section "Function" (3)
// 0000011: 03                                        ; section code
// 0000012: 00                                        ; section size (guess)
// 0000013: 01                                        ; num functions
// 0000014: 00                                        ; function 0 signature index
// 0000012: 02                                        ; FIXUP section size
// ; section "Export" (7)
// 0000015: 07                                        ; section code
// 0000016: 00                                        ; section size (guess)
// 0000017: 01                                        ; num exports
// 0000018: 06                                        ; string length
// 0000019: 6164 6454 776f                           addTwo  ; export name
// 000001f: 00                                        ; export kind
// 0000020: 00                                        ; export func index
// 0000016: 0a                                        ; FIXUP section size
// ; section "Code" (10)
// 0000021: 0a                                        ; section code
// 0000022: 00                                        ; section size (guess)
// 0000023: 01                                        ; num functions
// ; function body 0
// 0000024: 00                                        ; func body size (guess)
// 0000025: 00                                        ; local decl count
// 0000026: 20                                        ; local.get
// 0000027: 00                                        ; local index
// 0000028: 20                                        ; local.get
// 0000029: 01                                        ; local index
// 000002a: 6a                                        ; i32.add
// 000002b: 0b                                        ; end
// 0000024: 07                                        ; FIXUP func body size
// 0000022: 09                                        ; FIXUP section size
// ; section "name"
// 000002c: 00                                        ; section code
// 000002d: 00                                        ; section size (guess)
// 000002e: 04                                        ; string length
// 000002f: 6e61 6d65                                name  ; custom section name
// 0000033: 02                                        ; local name type
// 0000034: 00                                        ; subsection size (guess)
// 0000035: 01                                        ; num functions
// 0000036: 00                                        ; function index
// 0000037: 00                                        ; num locals
// 0000034: 03                                        ; FIXUP subsection size
// 000002d: 0a                                        ; FIXUP section size


// if you need a refresher on modules
// see https://webassembly.github.io/spec/core/binary/modules.html

// IMPORTANT - I created a `wasmParser` helper function to
// keep track of our position in the binary


// We parse and check the headers (MAGIC and VERSION)
export function checkHeader(wasm) {
    if (wasm.data.length < 8) {
        throw new Error(RuntimeErrors.ModuleTooShort);
    }

    if (wasm.data.length == 8) {
        throw new Error(RuntimeErrors.ModuleIsEmpty)
    };

    const magicString = String.fromCharCode(...wasm.readBytes(4));
    if (magicString !== "\0asm") {
        throw new Error(RuntimeErrors.InvalidMagicHeader);
    }

    const version = wasm.dword();
    if (version.length < 1) {
        throw new Error(RuntimeErrors.InvalidVersionHeader);
    }

    return true;
}

// For every section in the parser we check:
// - if the section is valid
// - the section's size 
// - the number, type or the number of types of the content inside the section

// We parse and check the type section
// See https://webassembly.github.io/spec/core/binary/modules.html#type-section
export function parseTypeSection(wasm) {
    
    const sectionType = wasm.readByte();
    if (sectionType !== Section.type) {
        throw new Error(RuntimeErrors.InvalidSection);
    }

    // helper function to map the valType (right now accepts only i32)
    function parseValueType(wasm) {
        const valType = wasm.readByte();

        switch (valType) {
            case ValType.i32:
                return ValType.i32;
            default:
                throw new Error(RuntimeErrors.InvalidValueType);
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
            params.push(parseValueType(wasm))
        }

        let numOfResults = wasm.readByte();
        let results = [];
        for (let w = 0; w < numOfResults; w++) {
            results.push(parseValueType(wasm))
        }
        types.push([params, results])
    }

    return types;
}

// we parse the function section (3)
// See https://webassembly.github.io/spec/core/binary/modules.html#function-section
export function parseFunctionSection(wasm) {
    const isSectionFunc = wasm.readByte()
    if (isSectionFunc !== Section.func) {
        throw new Error(RuntimeErrors.InvalidSection);
    }

    let sectionSize = wasm.readByte();
    // number of functions
    let numberOfFunctions = wasm.readByte();
    // Index of the function aka the position
    let functionSigIndex = [];

    for (let i = 0; i < numberOfFunctions; i++) {
        functionSigIndex.push(wasm.readByte())
    }

    return functionSigIndex;
}

// We parse export section
// See https://webassembly.github.io/spec/core/binary/modules.html#export-section
export function parseExportSection(wasm) {
    const isExportSection = wasm.readByte();
    if (isExportSection !== Section.export) {
        throw new Error(RuntimeErrors.InvalidSection);
    }

    let sectionSize = wasm.readByte();
    // How many exports in the module (we only have one)
    let numberOfExports = wasm.readByte();
    let exportsArr = [];

    for (let i = 0; i < numberOfExports; i++) {
        // we get the exported name
        const exportNameLength = wasm.readByte();
        const exportName = String.fromCharCode(...wasm.readBytes(exportNameLength));

        if (!!exportName === false) throw new Error(RuntimeErrors.InvalidExportName);
        // export kind
        let zero = wasm.readByte();
        // Index of the exported function
        let indexOfExportedFunction = (wasm.readByte() == ExportSection.func) && ExportSection.func;
        
        if (typeof indexOfExportedFunction !== 'number' && !!indexOfExportedFunction === false) {
            throw new Error(RuntimeErrors.InvalidExportType)
        }
     
        exportsArr.push({name: exportName.replace(/"/g, ""), index: indexOfExportedFunction})
    }

    return exportsArr;
}

// we parse the code section
// See https://webassembly.github.io/spec/core/binary/modules.html#code-section
export function parseCodeSection(wasm) {
    const isCodeSection = wasm.readByte();

    
    if (isCodeSection !== Section.code) {
        throw new Error(RuntimeErrors.InvalidSection);
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
            
            instructions.push(instruction)
            
            if (instruction == Opcodes.get_local) {
                locals.push(wasm.readByte())
            }
        } 
        code.push([locals, instructions]);
    }

    return code
}
