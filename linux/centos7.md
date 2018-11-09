#### 列出所以服务  

systemctl list-unit-files

#### nohup

0、1和2分别表示标准输入、标准输出和标准错误信息输出，可以用来指定需要重定向的标准输入或输出。

只输出标准错误信息

nohup ./program >/dev/null 2>log &

不输出任何信息，标准错误信息重定向到标准输出

nohup ./program >/dev/null 2>&1 &