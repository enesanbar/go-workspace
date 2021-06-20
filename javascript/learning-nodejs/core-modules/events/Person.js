// EventEmitter objects provides us a way to create custom objects that raises custom events.
const EventEmitter = require('events').EventEmitter;
const util = require('util');

const Person = function (name) {
    this.name = name;
};

util.inherits(Person, EventEmitter);

// Module.exports is the object that is returned by the require statement.
// When we require this module, we will return anything that is on module.exports.
module.exports = Person;