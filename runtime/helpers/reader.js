
export default class WasmReader {
  constructor(wasm) {
      this.data = wasm;
      this.pos = 0;
  }

  dword() {
      let prev = this.pos;
      this.pos += 4;
      const arr = this.data.slice(prev, this.pos);
      return arr
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
}
