const waitTime = 5000;
let currentTime = 0;
const waitInterval = 10;
let percentWaited = 0;

function writeWaitingPercentage(p) {
    process.stdout.clearLine();
    process.stdout.cursorTo(0);
    process.stdout.write(`waiting ... ${p}%`);
}

let interval = setInterval(() => {
    currentTime += waitInterval;
    percentWaited = Math.floor((currentTime / waitTime) * 100);
    writeWaitingPercentage(percentWaited);
}, waitInterval);

setTimeout(() => {
    clearInterval(interval);
    writeWaitingPercentage(100);
    console.log("\ndone\n");
}, waitTime);

process.stdout.write("\n");
writeWaitingPercentage(percentWaited);