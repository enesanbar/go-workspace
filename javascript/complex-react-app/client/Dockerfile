FROM node:10.13.0-alpine as build

ENV APP_PATH="/usr/src/app"

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH

COPY package*.json $APP_PATH/
RUN npm install

COPY . $APP_PATH
RUN npm run build

FROM nginx
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
COPY --from=build /usr/src/app/build /usr/share/nginx/html

EXPOSE 80
