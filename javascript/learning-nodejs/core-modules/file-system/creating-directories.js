const fs = require('fs');

if (fs.existsSync('lib')) {
    console.log('Directory already exists');
} else {
    fs.mkdir('lib', err => {
        if (err) throw err;

        console.log("Directory created");
    });
}
