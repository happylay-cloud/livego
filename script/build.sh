#!/bin/sh
##########################################################################
#Author :       happylay 安徽理工大学
#Created Time : 2021-02-03 01:44
#Environment :  darwin
##########################################################################

# 打包版本号（修改）
VERSION=1.0.0

# 应用名称（修改）
APP_NAME=live

cd ../script

# linux_amd64环境
gfctl build ../main.go --name $APP_NAME --arch amd64 --system linux --version $VERSION -p ../bin

# windows_amd64环境
gfctl build ../main.go --name $APP_NAME --arch amd64 --system windows --version $VERSION -p ../bin

# mac_amd64环境
gfctl build ../main.go --name $APP_NAME --arch amd64 --system darwin --version $VERSION -p ../bin

# -----------------------------------新增配置--------------------------------------
# 创建配置文件夹
mkdir -p ../bin/$VERSION/config

# 复制配置文件
cp ../config/* ../bin/$VERSION/config

# 压缩应用
cd ../bin/$VERSION

# 创建各个平台配置文件
mkdir -p ./darwin_amd64/config

mkdir -p ./linux_amd64/config

mkdir -p ./windows_amd64/config

cp ./config/* ./darwin_amd64/config

cp ./config/* ./linux_amd64/config

cp ./config/* ./windows_amd64/config

#--------------------------------------------------------------------------------------

tar -zcvf $APP_NAME.$VERSION-darwin-amd64.tar.gz darwin_amd64

tar -zcvf $APP_NAME.$VERSION-linux-amd64.tar.gz linux_amd64

zip -r $APP_NAME.$VERSION-windows-amd64.zip windows_amd64
