##### windows git clone 密码输入错误，被保存后，如何修改

打开控制面板->用户账户->凭据管理器，windows凭据。

根据git地址找到对应的内容，修改或删除密码。

##### git clone 带用户和密码

git clone http://username:password@127.0.0.1/res/res.git

##### 回滚到指定版本

git reset --hard 版本号



git 修改当前的project的用户名的命令为：**git config user.name 你的目标用户名;**

git 修改当前的project提交邮箱的命令为：**git config user.email 你的目标邮箱名;**

如果你要修改当前全局的用户名和邮箱时，需要在上面的两条命令中添加一个参数，--global，代表的是全局。

命令分别为：

git config  --global user.name 你的目标用户名；**

**git config  --global user.email 你的目标邮箱名;**

##### windows更新文件执行权限
git update-index --chmod=+x exe
