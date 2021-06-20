FROM node:alpine
WORKDIR /usr/src/app
COPY package.json yarn.lock ./
COPY . ./
RUN npm install

CMD ["npm", "start"]
