const url = 'http://mylogger.io/log';

console.log(__dirname);
console.log(__filename);

function info(message) {
    console.log(`[INFO] ${message}`);
}

function debug(message) {
    console.log(`[DEBUG] ${message}`);
}

module.exports = {
    "info": info,
    "debug": debug
};
