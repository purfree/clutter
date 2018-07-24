# nodejs使用web3

1. 安装git
2. 安装make yum -y install gcc automake autoconf libtool make
3. 安装g++ yum install -y gcc gcc-c++


nodejs下载  
https://nodejs.org/en/download/

配置NODE_HOME，进入profile编辑环境变量  
vim /etc/profile  
设置nodejs环境变量，在 export PATH USER LOGNAME MAIL HOSTNAME HISTSIZE HISTCONTROL 一行的上面添加如下内容:  
export NODE_HOME=/usr/local/node/0.10.24  
export PATH=$NODE_HOME/bin:$PATH  
:wq保存并退出，编译/etc/profile 使配置生效  
source /etc/profile  
验证是否安装配置成功  

node -v

we3js安装  
npm install -y web3  

扩展web3js未提供的以太坊接口函数，参考下面的地址  
go-ethereum/internal/web3ext/web3ext.go
