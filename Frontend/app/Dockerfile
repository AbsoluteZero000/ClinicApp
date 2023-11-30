FROM node:18.12.0 as build-stage

WORKDIR /app

COPY package*.json ./

RUN npm i -g @angular/cli@16.2.10


RUN npm i

COPY . .
EXPOSE 4200

CMD ["ng", "serve", "--host", "0.0.0.0"]
