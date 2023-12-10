ARG HOSTNAME=api.ecomm.com
FROM nginx:1.24-bullseye
ARG HOSTNAME
RUN apt-get update \
    && apt-get install -y nginx-extras apache2-utils vim openssl \
    &&   openssl req -x509 -nodes -days 365 -newkey rsa:2048  -subj "/C=IT/IT=CT/L=Catania/O=UNICT/CN=$HOSTNAME" -keyout /etc/ssl/private/"$HOSTNAME".key -out /etc/ssl/certs/"$HOSTNAME".crt
    #&& sed -i '28s/^/\tinclude \/etc\/nginx\/myconf.conf;\n/' /etc/nginx/nginx.conf