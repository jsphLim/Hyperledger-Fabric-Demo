### Hyperledger-Fabric

#### 开发环境
ubuntu16.10

#### 前期准备
docker安装 (我使用的是1.13.1)
```
sudo apt install docker.io  
```
docker-compose安装 (我使用1.19.0)
```
sudo curl -L https://github.com/docker/compose/releases/download/1.19.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose  
```
修改下权限
```
sudo chmod +x /usr/local/bin/docker-compose 
```

go语言 (我使用的是go1.8.3)<br>

下载地址: https://golang.org/dl/ 选择合适的版本 我的版本是go1.8.3.linux-amd64.tar.gz
```
tar -zxvf go1.8.3.linux-amd64.tar.gz -C /usr/local  
```
之后就可以在/usr/local下看到go文件夹了,接着配置环境变量，我的workspace是$HOME/go<br>
```
sudo vi ~/.profile  
```
在末尾加上
```
export PATH=$PATH:/usr/local/go/bin   
export GOROOT=/usr/local/go   
export GOPATH=$HOME/go   
export PATH=$PATH:$HOME/go/bin  
```
查看go是否安装成功
```
go version
```
#### Fabric网络搭建
进入工作目录$HOME/go

创建以下一系列文件夹:

$HOME/go/src/github.com/hyperledger

创建后进入hyperledger文件夹下，执行以下指令获取Fabric源码:


