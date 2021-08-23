## 🌀 Suscan 
<a href="https://github.com/tangxiaofeng7/Suscan"><img alt="Release" src="https://img.shields.io/badge/golang-1.16-9cf"></a>
<a href="hhttps://github.com/tangxiaofeng7/Suscan"><img alt="Release" src="https://img.shields.io/badge/Suscan-1.0-ff69b4"></a>
<a href="https://github.com/tangxiaofeng7/Suscan"><img alt="Release" src="https://img.shields.io/badge/LICENSE-GPL-important"></a>
![GitHub Repo stars](https://img.shields.io/github/stars/tangxiaofeng7/Suscan?color=success)
![GitHub forks](https://img.shields.io/github/forks/tangxiaofeng7/Suscan)  


## 😊 Run
安装docker与docker-compose
```
curl -fsSL [https://get.docker.com/](https://get.docker.com/) | sh  

systemctl restart docker  

systemctl enable docker  

wget [https://github.com/docker/compose/releases/download/1.29.2/docker-compose-Linux-x86_64](https://github.com/docker/compose/releases/download/1.29.2/docker-compose-Linux-x86_64)

sudo cp docker-compose-Linux-x86_64 /usr/bin/docker-compose

sudo chmod 777 /usr/bin/docker-compose
```

docker-compose启动
```
docker-compose up -d
```

## 😏 Api
添加扫描资产：
curl --location --request POST 'http://ip:18000/api/v1/assets' \
--form 'assets="txf.com"'

修改扫描配置：
curl --location --request POST 'http://ip:18000/api/v1/scansetting' \
--form 'thread="5000"' \
--form 'port="1-65535"' \
--form 'cmd="7"'

开启扫描：
curl --location --request GET 'http://ip:18000/api/v1/scan'

获取高危端口：
curl --location --request GET 'http://ip:18000/api/v1/getVulPort'

获取高危协议：
curl --location --request GET 'http://ip:18000/api/v1/getVulPro'