# Stage 1 - the build process
FROM node:alpine as build
ARG REACT_APP_API_URL=http://localhost:4000
ENV REACT_APP_API_URL ${REACT_APP_API_URL}
WORKDIR /usr/src/app
COPY package.json yarn.lock ./
RUN yarn
COPY . ./

RUN yarn build

# Stage 2 - the production environment
FROM nginx:alpine
COPY --from=build /usr/src/app/build /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY docker/nginx.conf /etc/nginx/conf.d
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
