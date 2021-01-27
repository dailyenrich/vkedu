FROM node:14.11.0-alpine3.10 as builder

WORKDIR /var/www

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk update && apk add git python3 && \
    npm install -g cnpm --registry=https://registry.npm.taobao.org

COPY ./ /var/www/vadmin

RUN cd /var/www/vadmin && rm -r node_modules && cnpm install && cnpm run build

FROM nginx:1.17-alpine

COPY --from=builder /var/www/vadmin/dist /var/www/vadmin
COPY ./nginx.conf /etc/nginx/conf.d/default.conf