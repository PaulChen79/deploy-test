FROM golang:1.17.1

ENV TZ=UTC

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update \
    && apt-get -y upgrade

RUN apt-get install -y apt-utils \
    supervisor \
    net-tools \
    vim

RUN mkdir -p /var/www/html/microService \
    && mkdir -p /var/log/supervisor \
    && mkdir -p /var/log/ServiceLog

COPY ./supervisor/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

WORKDIR /var/www/html/microService

EXPOSE 80

CMD ["/usr/bin/supervisord"]