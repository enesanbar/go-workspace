const amqp = require("amqplib");

const message = {
    description: "This is a test message"
}

connect_rabbitmq()

async function connect_rabbitmq() {
    try {

        const connection = await amqp.connect("amqp://localhost:5672");
        const channel = await connection.createChannel()
        const assertion = await channel.assertQueue("jobsQueue");

        console.log("Awaiting message..")
        channel.consume("jobsQueue", receivedMessage => {
            console.log(`Message received: ${receivedMessage.content.toString()}`);
            channel.ack(receivedMessage);
        });
    } catch (e) {
        console.log(`Error: ${e}`)
    }
}
