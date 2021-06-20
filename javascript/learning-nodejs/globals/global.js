// console object
const helloWorld = "Hello World";
console.log(`${helloWorld} from Node.js`);

// __dirname and __filename
console.log(`__dirname: ${__dirname}`);
console.log(`__filename: ${__filename}`);

// require
const path = require('path');
const baseName = path.basename(__filename);
console.log(`${helloWorld} from ${baseName}`);

console.log(module);