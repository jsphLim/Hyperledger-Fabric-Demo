### Hyperledger-Fabric

#### environment
ubuntu16.10

#### Preparation
install docker (version 1.13.1)
```
sudo apt install docker.io  
```
install docker-compose (version 1.19.0)
```
sudo curl -L https://github.com/docker/compose/releases/download/1.19.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose  
```
modify the permission
```
sudo chmod +x /usr/local/bin/docker-compose 
```

install go (version go1.8.3)<br>

download url: https://golang.org/dl/ my version is go1.8.3.linux-amd64.tar.gz
```
tar -zxvf go1.8.3.linux-amd64.tar.gz -C /usr/local  
```
after that we can find /go at /usr/local，and my workspace is at $HOME/go<br>
```
sudo vi ~/.profile  
```
then add the following instruction
```
export PATH=$PATH:/usr/local/go/bin   
export GOROOT=/usr/local/go   
export GOPATH=$HOME/go   
export PATH=$PATH:$HOME/go/bin  
```
then we should check whether go is installed
```
go version
```
#### Build a fabric net-work
enter the workplace $HOME/go

Create the following series of folders

$HOME/go/src/github.com/hyperledger

Enter the hyperledger folders,and download fabric source code
```
git clone https://github.com/hyperledger/fabric.git  
```
after that,Enter fabric/examples/e2e_cli

There are some shell files for us to start up the fabric net-work

firstly，start up download-dockerimages.sh to download images we need from docker hub:
```
chmod +x download-dockerimages.sh  
```
```
./download-dockerimages.sh  
```
then,use the following instruction to check wether the  images is downloaded successfully
```
docker images  
```
finally, start up the net-work:
```
./network_setup.sh up  
```
```
docker-compose -f docker-compose-cli.yaml up  
```
if everything goes well, the fabric net-work is builded successfully<br>
install my demo
```
https://github.com/jsphLim/Hyperledger-Fabric.git
```
start up startFabric.sh in Hyperledger-Fabric/app
```
sudo ./startFabric.sh
```
then register a admin:
```
node registerAdmin.js
```
then register a User:
```
node registerUser.js
```
if everything goes well,we can start the server:
```
node server.js
```
then we can visit the app at:
```
localhost:7775 
```
this demo only has two function: 'get' and 'set',It's just for study Hyperledger-Fabric<br>
If something wrong, you can contact me by the gmail
