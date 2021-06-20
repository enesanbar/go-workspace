const fs = require('fs');
const path = require('path');

// fs.readdirSync -> use sync functions to load conf file
fs.readdir('./', (err, files) => {
    if (err) throw err;

    console.log(`/ directory contains ${files.length} entities`);
    files.forEach((file) => {

        let filePath = path.join(__dirname, file);
        let stats = fs.statSync(filePath);

        if (stats.isFile()) {
            fs.readFile(file, 'UTF-8', (err, content) => {
                if (err) throw err;

                console.log(file);
                console.log(content);
            });
        }

    });

});
