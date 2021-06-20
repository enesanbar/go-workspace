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

    setInterval(() => {
        // send message to queue via channel
        message.time = new Date().getTime();
        channel.sendToQueue("jobsQueue", Buffer.from(JSON.stringify(message)));
        console.log(`Message sent: ${JSON.stringify(message)}`)
    }, 1000)

    } catch (e) {
        console.log(`Error: ${e}`)
    }
}
