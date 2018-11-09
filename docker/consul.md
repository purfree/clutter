#### 本机访问consul ui

启动consul时需要绑定client（ifconfig inet的值）

./consul agent -dev -ui -node=consul-dev -client=172.17.0.X

