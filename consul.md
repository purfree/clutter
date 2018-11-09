#### 注册服务

curl -X PUT -d '{"id": "mtail","name": "mtail","address": "172.17.0.3","port": 3903,"tags": ["dev"],"checks": [{"http": "http://172.17.0.3:3903/","interval": "5s"}]}'     http://172.17.0.2:8500/v1/agent/service/register