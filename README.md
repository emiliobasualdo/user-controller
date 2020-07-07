# Mas Simple Wallet Controller
##### Requisitos
- golang
- [eb](https://github.com/aws/aws-elastic-beanstalk-cli-setup)

# Go Module Code structure
https://github.com/golang-standards/project-layout
# Ecosystem structure
The ecosystem is based on [this](https://www.usenix.org/legacy/publications/library/proceedings/ec98/full_papers/daswani/daswani.pdf) paper. 
This controller only representes the wallet-controller module.

# Controller structure
![Controller Structure](https://srv-file6.gofile.io/download/t8uEcz/Screen%20Shot%202020-07-04%20at%2014.50.25%20copy.png)
# Setup Wallet Controller
```
$ git clone git clone https://bitbucket.org/mas-simple/wallet-controller.git
$ cd wallet controller
$ go get ./...
$ eb init // selecionar el enviroment en el que se quiera trabajar
```

## Run locally
```
$ go run application.go
```
### Run with auto-rebuild
```
$ go get github.com/cespare/reflex
```
and
```
$ sh scripts/runDev.sh
```
or
```
$ sh scripts/run.sh
```

## Deploy a current eb environment
```
$ git commit ..... # eb only deploys commited code
$ eb deploy
```

### For interested people; How to start a beanstalk?:
https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/go-getstarted.html
1) Create an app in application.go // beanstalk only start code in application.go 
2) `eb init`
3) `eb deploy`

- eb open // open the current env endpoint url