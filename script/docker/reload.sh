#!/bin/sh
##########################################################################
#Author :       happylay 安徽理工大学
#Created Time : 2021-02-09 02:15
#Environment :  centos7.6
##########################################################################
docker rm -f livego
docker rmi -f livego
docker-compose up -d
