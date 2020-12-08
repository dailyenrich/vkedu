FROM node:12.20.0-alpine3.12 as builder

COPY . /app

RUN apk add python3 && cd /app && rm -rf node_modules && npm install && npm run generate

# FROM node:12.20.0-alpine3.12

# COPY --from=builder /app/.nuxt /app
# COPY --from=builder /app/package.js /app

WORKDIR /app

EXPOSE 80

CMD [ "npm", "run","start" ]