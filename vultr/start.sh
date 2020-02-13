#!/bin/bash

#yum update -y
#yum install -y vim

echo "为了使用BBR，系统内核必须 >= 4.9.0"
echo "当前系统内核: `uname -r`"

read -p "是否要升级内核? Y/N: "
typeset -u reply
reply=$REPLY
if [ $reply == "Y" ]
then
    echo "安装软件源 ELRepo repo"
    sudo rpm --import https://www.elrepo.org/RPM-GPG-KEY-elrepo.org
    sudo rpm -Uvh http://www.elrepo.org/elrepo-release-7.0-2.el7.elrepo.noarch.rpm
    echo "安装系统内核"
    sudo yum --enablerepo=elrepo-kernel install kernel-ml -y
    echo "确认安装结果，安装成功可以看见kernel-ml-x.x.x-1.el7.elrepo.x86_64中的版本好大于4.9.0"
    rpm -qa | grep kernel
    echo "系统内核列表"
    sudo egrep ^menuentry /etc/grub2.cfg | cut -f 2 -d \'
    echo "启用新内核"
    sudo grub2-set-default 0
    echo "重启系统"
    sudo shutdown -r now
fi

read -p "是否启用BBR? Y/N: "
typeset -u reply
reply=$REPLY
if [ $reply == "Y" ]
then
    echo "修改sysctl配置"
    echo 'net.core.default_qdisc=fq' | sudo tee -a /etc/sysctl.conf
    echo 'net.ipv4.tcp_congestion_control=bbr' | sudo tee -a /etc/sysctl.conf
    sudo sysctl -p
    echo "启用BBR"
    sudo sysctl net.ipv4.tcp_available_congestion_control
    echo "验证BBR是否启动成功, 成功会输出bbr"
    sudo sysctl -n net.ipv4.tcp_congestion_control
    echo "检查内核是否加载BBR, 成功会包含bbr"
    lsmod | grep bbr
fi

read -p "安装ss, 请选择版本(go): "
typeset -u reply
reply=$REPLY
if [ $reply == "GO" ]
then
    yum install -y go
    go get github.com/shadowsocks/shadowsocks-go/cmd/shadowsocks-server
fi

echo "ss启动中"
read -p "请输入服务器监听端口: " port
read -p "请输入数据加密格式: " method
read -p "请输入服务器访问密码: " password
# read -p "请输入是否启用UDP回复: " udp

if [ -z $port ]
then
    echo "必须输入服务器监听端口"
    exit
fi

if [ -z $method ]
then
    method="aes-256-cfb"
fi

if [ -z $password ]
then
    echo "必须输入访问密码"
    exit
fi

# if [ -z $udp ]
# then
#     udp=false
# fi

echo "添加端口"
firewall-cmd --zone=public --add-port=$port/tcp --permanent 
echo "重新加载"
firewall-cmd --reload
echo "查看所有打开的端口"
firewall-cmd --zone=public --list-ports

nohup /root/go/bin/shadowsocks-server -p $port -k $password -m $method > /root/go/bin/log 2>&1 &

