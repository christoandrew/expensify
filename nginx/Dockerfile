FROM nginx:latest

VOLUME log/access.log /var/log/nginx/access.log
VOLUME log/error.log /var/log/nginx/error.log
VOLUME log/expensify.log /var/log/nginx/expensify.log
COPY nginx.conf /etc/nginx/nginx.conf
COPY mime.types /etc/nginx/mime.types