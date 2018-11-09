#### 安装

docker pull Name

#### 运行

交换式运行

docker run -it Name:Version /bin/bash 

后台运行

docker run -d Name:Version /bin/bash

-v 挂载目录

-p 端口映射

#### 进入运行中的容器

docker attach ImageID    

docker exec -it ImageID /bin/bash  

#### 打印容器name-ip  

docker inspect -f "{{.Name}} - {{.NetworkSettings.IPAddress }}" ImageID  

#### 保存镜像  

docker commit imageID new_name:new_tag  

#### 注意事项

docker挂载,	如果挂载的时候需要授权，在修改系统密码后需要重置授权密码，Shared Drives -> Reset credentials。