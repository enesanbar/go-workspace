FROM node:10.13.0-alpine

ENV APP_PATH="/usr/src/app"

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH

COPY package*.json $APP_PATH/
RUN npm install

COPY . $APP_PATH
RUN npm install -g nodemon
CMD ["nodemon", "server.js"]
