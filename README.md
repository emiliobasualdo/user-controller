# Mas Simple Wallet Controller
##### Requisitos
- golang
- [eb](https://github.com/aws/aws-elastic-beanstalk-cli-setup)

# Go Module Code structure
https://github.com/golang-standards/project-layout
# Ecosystem structure
The ecosystem is based on [this](https://www.usenix.org/legacy/publications/library/proceedings/ec98/full_papers/daswani/daswani.pdf) paper. 
This controller only represents the wallet-controller module.

# Controller structure
![Controller Structure](https://srv-file6.gofile.io/download/t8uEcz/Screen%20Shot%202020-07-04%20at%2014.50.25%20copy.png)
# Setup Wallet Controller
```
$ git clone git clone https://bitbucket.org/mas-simple/wallet-controller.git
$ cd wallet controller
$ go get ./...
```

## JWT
To generate a jwt HS256 secret key follow this steps:
```
openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -pubout -out public.pem
cat public.pem
```
Then copy the public key between //// wrappers and paste it in the environment config file 

## Gerneate go code from swagger.yaml
Install  swagger-codegen
```
brew install swagger-codegen
```
Generate the code
```
swagger-codegen generate -i path/to/swagger.json.yaml -l go -o /out/dir/
```
