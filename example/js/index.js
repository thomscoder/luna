import startAeonRuntime from '../dist/bundle.js';
import { defaultText } from './inputDefaultText.js';
import { CONST_MODE } from './texts/texts.js';

const go = new Go(); // Defined in wasm_exec.js. Don't forget to add this in your index.html.


const wasmBrowserInstantiate = async (wasmModuleUrl, importObject) => {
  let response = undefined;

  // Check if the browser supports streaming instantiation
  if (WebAssembly.instantiateStreaming) {
    // Fetch the module, and instantiate it as it is downloading
    response = await WebAssembly.instantiateStreaming(
      fetch(wasmModuleUrl),
      importObject
    );
  } else {
    // Fallback to using fetch to download the entire module
    // And then instantiate the module
    const fetchAndInstantiateTask = async () => {
      const wasmArrayBuffer = await fetch(wasmModuleUrl).then(response =>
        response.arrayBuffer()
      );
      return WebAssembly.instantiate(wasmArrayBuffer, importObject);
    };

    response = await fetchAndInstantiateTask();
  }

  return response;
};

const runLunaAddition = async () => {
  // Get the importObject from the go instance.
  const importObject = go.importObject;

  // Instantiate our wasm module
  const wasmModule = await wasmBrowserInstantiate("./main.wasm", importObject);
  // Allow the wasm_exec go instance, bootstrap and execute our wasm module
  go.run(wasmModule.instance);



  // Set the result onto the doc
  const input1 = document.getElementById('input-1');
  const input2 = document.getElementById('input-2');
  const label2 = document.getElementById('label-2')
  const codeContainer = document.getElementById('code');
  const moduleContainer = document.getElementById('module');
  const selectInstruction = document.getElementById('instruction');

  const compile = document.getElementById('compile');
  const btn = document.getElementById('btn')

  const res = document.getElementById('result');
  const aeonRes = document.getElementById('result-aeon');


  btn.setAttribute('disabled', true)

  let editor = CodeMirror(codeContainer, {
    value: defaultText(selectInstruction),
    mode:  "wast",
    lineNumbers: true,
  });

  // Change editor content based on selected mode
  selectInstruction.onchange = (e) => {
    editor.setValue(defaultText(selectInstruction));
    // Hide second input on const
    if (e.target.value === CONST_MODE) {
      input2.style.display = 'none';
      label2.style.display = 'none';
    }
  }

  compile.addEventListener('click', async () => {
    moduleContainer.innerHTML = ""
    const textContent = editor.getValue()
    // We call the startLuna function
    const _wasm = startLuna(textContent).module.split(" ").map(v => parseInt(v, 10))

    const wasm = Uint8Array.from(_wasm)

    try {
      const wasmer = await WebAssembly.instantiate(wasm);
      const fn = Object.keys(wasmer.instance.exports)[0]
      moduleContainer.innerHTML = `<p>Compiled successfully!\nExported function <span class="exported">${fn}</span></p>\n`

      for (const hex of wasm) {
        moduleContainer.innerHTML += `<p class="hex-dump">${hex}</p>`
      }

      btn.removeAttribute('disabled')

      btn.addEventListener('click', () => {
        const n1 = Number(input1.value);
        const n2 = Number(input2.value)
        // Call Luna add function
        res.innerHTML = "Result: " 
        res.innerHTML += wasmer.instance.exports[fn](n1, n2)

      // ---------------------------------------------------------------------------------------
      // ---------------------------------------------------------------------------------------
      // ------------------------------RUN IT WITH AEON RUNTIME---------------------------------
      // ---------------------------------------------------------------------------------------
      // ---------------------------------------------------------------------------------------
      const funcName = fn.replace(/"/g, '');

      const result = startAeonRuntime(_wasm, funcName, n1, n2);
      aeonRes.innerHTML = `Ran with Aeon: ${result}`
      })
    } catch (err) {
      btn.setAttribute('disabled', true)
      moduleContainer.innerHTML = String(err);
    }

  })
  
};


runLunaAddition();