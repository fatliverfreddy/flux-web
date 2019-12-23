FROM node:lts-alpine3.11 as builder

WORKDIR /app

COPY package*.json ./

RUN npm i

COPY . .

RUN npm run build

CMD npm start
