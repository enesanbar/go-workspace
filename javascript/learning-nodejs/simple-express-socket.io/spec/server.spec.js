const request = require('request');

describe('calc', () => {
    it('should multiply 2 and 2', () => {
        expect(2 * 2).toBe(4);
    });
});

describe('get messages', () => {
    it('should return 200 status', (done) => {
        request.get("http://localhost:3000/messages", (err, response) => {
            expect(response.statusCode).toEqual(200);
            done();
        });
    });

    it('should return a populated list', (done) => {
        request.get("http://localhost:3000/messages", (err, response) => {
            expect(JSON.parse(response.body).length).toBeGreaterThan(0);
            done();
        });
    });
});

describe('get messages from user', () => {

    const name = "bayram";
    const message = "hi there";

    beforeEach(function () {
        request.post('http://localhost:3000/messages', {
            form: {
                name: name,
                message: message
            }
        });
    });

    it('should return 200 status', (done) => {
        request.get("http://localhost:3000/messages/" + name, (err, response) => {
            expect(response.statusCode).toEqual(200);
            done();
        });
    });

    it('name should be Tim', function (done) {
        request.get("http://localhost:3000/messages/" + name, (err, response) => {
            expect(JSON.parse(response.body)[0].name).toEqual(name);
            done();
        });
    });
});