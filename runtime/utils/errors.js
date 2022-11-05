const RuntimeErrors = {
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
  ExportNotFound: "Export not found",
}

const _RuntimeErrors = RuntimeErrors;
export { _RuntimeErrors as RuntimeErrors };