FROM node:alpine as builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci
COPY src tsconfig.json ./
RUN npm run build && sleep 30s

FROM node:alpine as runtime

WORKDIR /app

COPY --from=builder /app/bin ./bin
COPY . .

ENTRYPOINT [ "npm", "run", "start" ]
