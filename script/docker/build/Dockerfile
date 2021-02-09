##########################################################################
#Author :       happylay 安徽理工大学
#Created Time : 2021-02-09 02:15
#Environment :  centos7.6
##########################################################################

FROM happylay/openjdk:8.0.0

ADD livego /

ADD config /config/

ADD statics /statics/

VOLUME /config

VOLUME /statics

COPY entrypoint.sh /usr/local/bin

RUN chmod 755 /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]
