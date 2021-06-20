const Person = require('./Person');
const Logger = require('./logger');
const logger = new Logger();

const ben = new Person("Ben Franklin");
const george = new Person("George Washington");

ben.on('speak', function (said) {
    console.log(`${this.name}: ${said}`);
});

george.on('speak', function (said) {
    console.log(`${this.name} -> ${said}`);
});

ben.emit('speak', 'You may delay, but time will not');
george.emit('speak', 'It is far better to be alone than to be in bad company.');

logger.on('messageLogged', function (data) {
    console.log(`george responded to event: ${data}`);
});

logger.info("some message");