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
  const codeContainer = document.getElementById('code');
  const moduleContainer = document.getElementById('module');

  const compile = document.getElementById('compile');
  const btn = document.getElementById('btn')

  const res = document.getElementById('result');


  btn.setAttribute('disabled', true)
  const input = `(module
  (func (export "addNumbers") (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.add)
)
`


    const editor = CodeMirror(codeContainer, {
      value: input,
      mode:  "wast",
      lineNumbers: true,
    });

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
        console.log("n", input1.value)
        console.log("n", input2.value)
        // Call Luna add function
        res.innerHTML = "Result: " 
        res.innerHTML += wasmer.instance.exports[fn](n1, n2)

      })
    } catch (err) {
      btn.setAttribute('disabled', true)
      moduleContainer.innerHTML = String(err);
    }

    
  })
  
};


runLunaAddition();