	# BuildingDistributedAppsGo

10/04/2021

Gitkraken
Alt + T
git status
git add .
git commit -m "My message"

GitKraken | push


Weird not working
go version 
which go

Command 'go' not found, but can be installed with:

sudo snap install go         # version 1.16.2, or
sudo apt  install golang-go  # version 2:1.13~1ubuntu2
sudo apt  install gccgo-go   # version 2:1.13~1ubuntu2


sudo snap install go --classic
sudo apt  install golang-go  

sudo apt  install gccgo-go 
Unable to locate gcc-go


go version
go build .
Now works OK		


IMPORTANT
go mod init app


LogService TEST
cd app/cmd/logservice
go build .

NB:
I am getting some error about GOROOT not set because I installed snap
weird as this was all working the other day


POSTMAN Canary
https://www.postman.com/downloads/canary

cmd prompt
cd ~
cd Steven
cd Installation
cd PostmanCanary
./PostmanCanary
