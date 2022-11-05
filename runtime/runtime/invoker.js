import { RuntimeErrors } from '../utils/errors.js';
import Processor from './processor.js';

// we need to call the exported function(s)
export function invokeFunction(ast, funcName, params) {
  // we search if the export name does exists
  // if not we throw an error
  let exportIndex = ast.exportSection.findIndex(exp => exp.name === funcName);
  if (exportIndex === -1) throw new Error(RuntimeErrors.ExportNotFound);

  // Get the exported function's body
  let func = ast.codeSection[exportIndex];
  // Check if the number of parameters of the exported function
  // corresponds to the number of provided parameters
  if (ast.typeSection[exportIndex][0].length !== params.length) {
      throw new Error(RuntimeErrors.InvalidArgumentsNumber);
  }

  // execute the function
  const processor = new Processor(func, params);
  processor.executeFunc();

  return processor.getResult();
}