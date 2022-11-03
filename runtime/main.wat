;; Example for the Aeon runtime and Luna compilation
(module
  (func (export "addTwo") (param i32) (param i32) (result i32)
    local.get 0
    local.get 1
    i32.add)
)