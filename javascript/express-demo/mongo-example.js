const mongoose = require('mongoose');

mongoose.connect('mongodb://localhost/playground', { useNewUrlParser: true, useUnifiedTopology: true })
        .then(() => console.log('Connected to MongoDB'))
        .catch(() => console.log('Could not connect to MongoDB'));

const courseSchema = new mongoose.Schema({
        name: String,
        tags: [String],
        date: {
                type: Date, default: Date.now
        }
});

const Course = mongoose.model('Course', courseSchema);

async function createCourse() {
        const course = new Course({
                name: "NodeJS Course",
                tags: ['node', 'backend']
        });

        const result = await course.save();
        console.log(result);
}

createCourse();