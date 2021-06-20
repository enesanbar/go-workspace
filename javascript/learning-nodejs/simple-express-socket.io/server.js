const express = require('express');
const bodyParser = require('body-parser');
const mongoose = require('mongoose');
const app = express();
const http = require('http').Server(app);
const io = require('socket.io')(http);

app.use(express.static(__dirname));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: false}));

mongoose.Promise = Promise;

const Message = mongoose.model('Message', {
    name: String,
    message: String
});

app.get('/messages', (req, res) => {
    Message.find({}, (err, messages) => {
        if (err) throw err;
        res.send(messages);
    });
});

app.get('/messages/:user', (req, res) => {
    const user = req.params.user;
    Message.find({name: user}, (err, messages) => {
        if (err) throw err;
        res.send(messages);
    });
});

io.on('connection', (socket) => {
    console.log('user connected');
});

app.post('/messages', async (req, res) => {

    try {
        const message = Message(req.body);

        const savedMessage = await message.save();
        console.log('saved');

        const censored = await Message.findOne({message: 'badword'});
        if (censored) {
            console.log("censored words found!");
            await Message.deleteOne({_id: censored.id})
        } else {
            io.emit('message', req.body);
        }

        res.sendStatus(200);
    } catch (e) {
        res.sendStatus(500);
        console.error(e);
    } finally {
        console.log("/message post called");
    }
});

mongoose.connect("mongodb://mongo/chat", (err) => {
    console.log("Connected to mongo\nerr: " + err);
});

const server = http.listen(3000, (err) => {
    if (err) throw err;

    console.log(`Listening on port ${server.address().port}`);
});