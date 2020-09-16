FROM node:10

RUN npm i -g express compress
COPY . /app

WORKDIR /app
RUN npm install
RUN npm run build:prod

EXPOSE 8080
CMD ["node", "server.js"]