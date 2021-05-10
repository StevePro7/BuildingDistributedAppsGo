# Building Distributed Apps in Go
Code from Building Distributed Applications in Go Pluralsight course
<br />
https://app.pluralsight.com/library/courses/building-distributed-applications-go

By Mike Van Sickle https://github.com/vansimke

```bash
cd app/cmd/registryservice
go build .
./registryservice
```

```bash
cd app/cmd/logservice
go build .
./logservice
```

```bash
cd app/cmd/gradingservice
go build .
./gradingservice
```

```bash
cd app/cmd/TeacherPortal
go build .
./TeacherPortal
```

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



REGISTRATION services
Two terminals

cd app/cmd/registryservice
go build .
./registryservice

cd app/cmd/logservice
go build .
./logservice



registryservice
De-register
cd app/cmd/registryservice
./registryservice

cd app/cmd/logservice
go build .
./logservice


http://localhost:4000
http://localhost:4000x


Windows VS Code
ST1005 go-staticcheck error strings should not be capitalized
https://staticcheck.io/docs/checks

problem
return fmt.Errorf("Service at Url %v not found", url)

solution
return fmt.Errorf("service at url %v not found", url)
