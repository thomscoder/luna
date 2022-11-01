const fs = require('fs');

try {
  const data = fs.readFileSync('./dist/main.wasm', 'binary');
  console.log(data)

  const main = async () => {
    const instance = await WebAssembly.instantiate(wasm);
    return instance
}

main().then(i => console.log("mama", i)).catch(console.log)
} catch (err) {
  console.error(err);
}
