//docker运行centos7  
docker run --privileged -t -i -v /e/docker/centos/data:/data -p 80:80 centos-ssh:latest /usr/sbin/init  
docker run --privileged -d -v /e/docker/centos/data:/data -p 80:80 centos-ssh:latest /usr/sbin/init  

//进入运行的容器  
docker attach ImageID  
docker exec -it ImageID /bin/bash  

打印容器name-ip  
docker inspect -f "{{.Name}} - {{.NetworkSettings.IPAddress }}" ImageID  

//新下载的容器需要更新  
yum update -y  
  
yum install -y net-tools //ifconfig  

//保存镜像  
docker commit imageID new_name:new_tag  

docker挂载,	如果挂载的时候需要授权，在更新系统后需要重置，Shared Drives -> Reset credentials，例如系统密码更新后需要更新，docker不会主动提示密码过期。
