FROM node:10.13.0-alpine

ENV APP_PATH="/usr/src/app"

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH

COPY package.json $APP_PATH/package.json
RUN npm install

COPY . $APP_PATH

CMD ["npm", "run", "start"]
