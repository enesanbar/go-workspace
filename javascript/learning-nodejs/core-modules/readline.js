const readline = require('readline');

// readline is a wrapper around stdin and stdout objects
// it allows us to ask questions to the terminal users.
const rl = readline.createInterface(process.stdin, process.stdout);

const realPerson = {
   name: '',
   sayings: []
};

rl.question("What is the name of a real person?\n > ", answer => {
    realPerson.name = answer;

    rl.setPrompt(`What would ${realPerson.name} say?\n > `);
    rl.prompt();

    rl.on('line', answer => {
        if (answer.toLowerCase().trim() === 'exit') {
            rl.close();
        } else {
            realPerson.sayings.push(answer.trim());
            rl.setPrompt(`What else would ${realPerson.name} say? ('exit' to leave)\n > `);
            rl.prompt();
        }
    });
});

rl.on('close', () => {
    console.log("%s is a real person that says %j", realPerson.name, realPerson.sayings);
    process.exit();
});