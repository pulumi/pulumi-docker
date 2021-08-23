FROM node:15 as builder

WORKDIR /app
COPY package.json .

RUN npm i

COPY app.js .

FROM node:15-alpine as runtime

WORKDIR /app

EXPOSE 3000/tcp

ENTRYPOINT [ "node", "app.js" ]

COPY --from=builder /app .