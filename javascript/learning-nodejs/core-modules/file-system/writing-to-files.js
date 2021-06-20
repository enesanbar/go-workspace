const fs = require('fs');

const data = {
    name: "John"
};

fs.writeFile('.newfile.json', JSON.stringify(data), (err) => {
    if (err) throw err;

    console.log("File has been saved");
});
