// process.argv is an array
console.log(process.argv);

function grab(flag) {
    const index = process.argv.indexOf(flag);
    return (index === -1) ? null : process.argv[index + 1]
}

const greeting = grab("--greeting");
const user = grab("--user");

if (!greeting || !user) {
    console.log("Where's the parameters!!!");
} else {
    console.log(`Welcome ${user}, ${greeting}`);
}

