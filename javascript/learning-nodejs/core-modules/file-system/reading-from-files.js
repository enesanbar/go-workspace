const fs = require('fs');

// fs library reads the data as string
fs.readFile('./data.json', 'utf-8', (err, data) => {
    if (err) throw err;

    console.log("Data read with readFile is String.");
    console.log(data);
    const json = JSON.parse(data);
    console.log("Name: " + json.name);
});

// if we get json file with require, we read it as json
const data = require('./data');
console.log("Data imported with require is JSON");
console.log(data);
console.log("Name: " + data.name);
