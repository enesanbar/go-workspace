const express = require('express');
const Joi = require('joi');

const router = express.Router();
router.use((req, res, next) => {
    res.setHeader('Content-Type', 'application/json');
    console.log(req.body);
    next();
});

const courses = [
    { id: 1, name: 'Microprocessors' },
    { id: 2, name: 'Operating Systems' },
    { id: 3, name: 'Database Systems' },
];

router.get('/', (req, res) => {
    const sortBy = req.query.sortBy;
    if (sortBy && sortBy == 'name') {
        courses.sort((a, b) => a[sortBy] > b[sortBy]);
    }
    res.send(JSON.stringify(courses));
});

router.get('/:id', (req, res) => {
    const course = courses.find(course => course.id === parseInt(req.params.id));

    if (!course) {
        return res.status(404).send(JSON.stringify({
            "status": 404,
            "message": "The course with a given id cannot be found"
        }));
    }
    res.send(course);
});

router.post('/', (req, res) => {
    // Validate
    const { error } = validateCourse(req.body);
    if (error) {
        return res.status(400).send({
            "status": 400,
            "message": error.details[0].message
        });
    }

    const course = {
        id: courses.length + 1,
        name: req.body.name
    }

    courses.push(course);
    res.status(201).send(course);
});

router.put('/:id', (req, res) => {
    // Look up the course
    let course = courses.find(course => course.id === parseInt(req.params.id));
    if (!course) {
        return res.status(404).send(JSON.stringify({
            "status": 404,
            "message": "The course with a given id cannot be found"
        }));
    }

    // Validate
    const { error } = validateCourse(req.body);
    if (error) {
        return res.status(400).send({
            "status": 400,
            "message": error.details[0].message
        });
    }

    // Update course
    course['name'] = req.body.name;

    // Send the updated result
    res.status(200).send(course);
});

router.delete('/:id', (req, res) => {
    // Look up the course
    let course = courses.find(course => course.id === parseInt(req.params.id));
    if (!course) {
        return res.status(404).send({
            "status": 404,
            "message": "The course with a given id cannot be found"
        });
    }

    const index = courses.indexOf(course);
    courses.splice(index, 1);

    res.send(course);
});

function validateCourse(course) {
    const schema = {
        name: Joi.string().min(3).required()
    };
    return Joi.validate(course, schema);
}

module.exports = router;
