const path = require('path');

console.log(path.basename(__filename));

const uploadDirectory = path.join(__dirname, 'www', 'files', 'uploads');
console.log(uploadDirectory);

const pathObj = path.parse(__filename);
console.log(pathObj);
