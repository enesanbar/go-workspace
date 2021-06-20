const express = require('express');
const config = require('config');
const helmet = require('helmet');
const morgan = require('morgan');
const debug = require('debug')('app:startup');
const logger = require('./middlewares/logger');
const auth = require('./authentication');
const courses = require('./routes/courses');
const home = require('./routes/home');

const app = express();

// Set the templating engine
app.set('view engine', 'pug');
app.set('views', './views'); // default

debug(`NODE_ENV: ${process.env.NODE_ENV}`);
debug(`app.get('env'): ${app.get('env')}`);
debug(`Application Name: ${config.get('name')}`);
debug(`Mail Server: ${config.get('mail.host')}`);
// debug(`Mail Password: ${config.get('mail.password')}`);

if (app.get('env') === 'development') {
    app.use(morgan('tiny'));
    debug('Morgan enabled.');
}

app.use(helmet());
app.use(express.json());
// app.use(express.urlencoded({extended: true}));
app.use(express.static('public'));
app.use(logger);
app.use(auth);

app.use('/api/courses', courses);
app.use('/', home);

const port = process.env.PORT || 3000;
app.listen(port, () => {
    console.log(`Listening on port ${port}`);
});
