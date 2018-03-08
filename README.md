### Hyperledger-Fabric

#### environment
Ubuntu16.10

#### Preparation
Install docker (version 1.13.1)
```
sudo apt install docker.io  
```
Install docker-compose (version 1.19.0)
```
sudo curl -L https://github.com/docker/compose/releases/download/1.19.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose  
```
Modify the permission
```
sudo chmod +x /usr/local/bin/docker-compose 
```

Install go (version go1.8.3)<br>

Download url: https://golang.org/dl/ my version is go1.8.3.linux-amd64.tar.gz
```
tar -zxvf go1.8.3.linux-amd64.tar.gz -C /usr/local  
```
After that we can find /go at /usr/local，and my workspace is at $HOME/go<br>
```
sudo vi ~/.profile  
```
Then add the following instruction
```
export PATH=$PATH:/usr/local/go/bin   
export GOROOT=/usr/local/go   
export GOPATH=$HOME/go   
export PATH=$PATH:$HOME/go/bin  
```
Then we should check whether go is installed
```
go version
```
#### Build a fabric net-work
Enter the workplace $HOME/go

Create the following series of folders

$HOME/go/src/github.com/hyperledger

Enter the hyperledger folders,and download fabric source code
```
git clone https://github.com/hyperledger/fabric.git  
```
After that,enter fabric/examples/e2e_cli

There are some shell files for us to start up the fabric net-work

Firstly，start up download-dockerimages.sh to download images we need from docker hub:
```
chmod +x download-dockerimages.sh  
```
```
./download-dockerimages.sh  
```
Then,use the following instruction to check wether the  images is downloaded successfully
```
docker images  
```
Finally, start up the net-work:
```
./network_setup.sh up  
```
```
docker-compose -f docker-compose-cli.yaml up  
```
If everything goes well, the fabric net-work is builded successfully<br>
install my demo
```
git clone https://github.com/jsphLim/Hyperledger-Fabric.git
```
Start up startFabric.sh in Hyperledger-Fabric/app
```
sudo ./startFabric.sh
```
Then register a admin:
```
node registerAdmin.js
```
Then register a User:
```
node registerUser.js
```
If everything goes well,we can start the server:
```
node server.js
```
Then we can visit the app at:
```
localhost:7775 
```
This demo only has two function: 'get' and 'set',It's just for study Hyperledger-Fabric<br>
If something wrong, you can contact me by the gmail
