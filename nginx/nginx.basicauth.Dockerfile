FROM nginx
RUN apt-get update \
    && apt-get install -y nginx-extras apache2-utils vim \
    && htpasswd -b -c /etc/nginx/.htpasswd admin admin
    #&& sed -i '28s/^/\tinclude \/etc\/nginx\/myconf.conf;\n/' /etc/nginx/nginx.conf