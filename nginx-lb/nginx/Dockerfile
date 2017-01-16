FROM debian:jessie

ARG CONFD_VERSION=0.12.0-alpha3
ARG NGINX_VERSION=1.11.8-1

RUN \
     apt-key adv --keyserver hkp://pgp.mit.edu:80 --recv-keys 573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62 \
	&& echo "deb http://nginx.org/packages/mainline/debian/ jessie nginx" >> /etc/apt/sources.list \
  && apt-get update \
  && apt-get install --no-install-recommends -y \
    curl \
    ca-certificates \
    nginx=${NGINX_VERSION}~jessie

RUN \
    curl -sL -o /bin/confd https://github.com/kelseyhightower/confd/releases/download/v${CONFD_VERSION}/confd-${CONFD_VERSION}-linux-amd64 \
    && chmod +x /bin/confd

COPY init.sh /
COPY etc/nginx/ /etc/nginx/
COPY etc/confd/ /etc/confd/

ENV \
  LISTEN_IPV4=0.0.0.0 \
  LISTEN_IPV6=::

CMD ["/init.sh"]

EXPOSE 80
