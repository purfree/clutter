#### 安装

docker pull centos

#### 运行

交换式运行

docker run --privileged -t -i centos:latest /bin/bash

后台运行

docker run --privileged -d centos:latest /usr/sbin/init 

#### 进入运行中的容器

docker attach ImageID    

docker exec -it ImageID /bin/bash  

#### 推荐安装的软件

yum update -y  //先执行update

yun install -y zip	//zip

yum install -y unzip	//unzip

yum install -y net-tools	//ifconfig

yum install -y vim		//vim