const EventEmitter = require('events');

class Logger extends EventEmitter{

    info(message) {
        console.log(`[INFO] ${message}`);
        this.emit('messageLogged', {message: message});
    }

    debug(message) {
        console.log(`[DEBUG] ${message}`);
    }

}

module.exports = Logger;
