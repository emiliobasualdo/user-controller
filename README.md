# Mas Simple Wallet Controller
##### Requisitos
- golang
- [eb](https://github.com/aws/aws-elastic-beanstalk-cli-setup)

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

## Deploy a current eb environment
```
$ git commit ..... # eb solo deploya lo que esté commiteado
$ eb deploy
```

## Para interesados
### Cómo levantar un beanstalk: 
https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/go-getstarted.html
1) meter código en application.go
2) eb init
3) eb deploy

- eb open // open the current env endpoint url