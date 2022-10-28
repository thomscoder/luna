;; Inside this wasm module there are two main things:
;; - A function (identified with the keyword func)
;; - The export (identified with the keyword export)

;; More generally speaking this module consists of 3 parts
;; - Tokens: special keywords reserved by the language (e.g. func, param, module, local.get etc...)
;; - Identifiers: what can be set to arbitrary values (e.g. $firstNumber, $secondNumber)
;; - Value Types: defined by the Web Assembly specifications (e.g. i32)
(module
    (func $add (param $firstNumber i32) (param $secondNumber i32) (result i32)
        local.get $firstNumber
        local.get $secondNumber
        i32.add)
    (export "add" (func $add))
)