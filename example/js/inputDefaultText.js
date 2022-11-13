import { ADDITION_MODE, CONST_MODE, DIVISION_MODE, MULTIPLICATION_MODE, SUBTRACTION_MODE } from "./texts/texts.js"

export const defaultText = (select) => {
    const instruction = select.value
    switch (instruction) {
        case ADDITION_MODE:
            return `(module
  (func (export "addNumbers") (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.add)
)
`
        case SUBTRACTION_MODE:
            return `(module
  (func (export "subtractNumbers") (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.sub)
)
`
        case MULTIPLICATION_MODE:
            return `(module
  (func (export "multiplyNumbers") (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.mul)
)
`
        case DIVISION_MODE:
            return `(module
  (func (export "divideNumbers") (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.div)
)
`
        case CONST_MODE:
            return `(module
  (func (export "operationWithInternalVariable") (param i32 i32) (result i32)
    local.get 0
    i32.const 10
    i32.add)
)
`
    }
}