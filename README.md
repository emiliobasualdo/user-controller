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
$ eb init // selecionar el enviroment en el que se quiera trabajar
```

# DEV environment
## Config file
The dev config file must be placed in the root folder of this project.  
The prod config file must be placed in `$HOME/config.yaml` in the EC2 instance.  
## Connect to db
We can connect to a local MySQL db or to the Aws MySQL instance.  
For either, we must specify the `connection string` in the config file.
### Connect to local db
Set the mysql connection string in `config-dev.yaml`
### Connect to aws db
[AWS/EB provides with environment variables](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features.managing.db.html) for the db connection string in production.
If you want to connect to the production db consider that the AWS instance is inside a VPC and has no public ip, 
so we can only connect from within the vpc. 
For that we need to create a ssh tunnel to the EC2. For this, you must first follow the `eb init` and `eb ssh` instructions.  
Then run:
```
$ eb ssh --custom 'ssh -N  -L 3307:aa1pre92f7a8m7u.cbbexwsgvaxw.sa-east-1.rds.amazonaws.com:3306' // {local-port}:{mysql-host}:{remote-port}
``` 
Now the local application just needs to access localhost:3307
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
The current EB environment is deployed into a EC2 inside a VPC where the RDS is.  
Use the deployment bash script to deploy as it uploads the config file to the EC2 instance.  
The EB environment must have the environment variable `ENV` set to `PROD`. 
This can be done with `eb setenv ENV=PROD`, this has only to be done once, just posting here for future reference.  
To deploy just:  
```
$ git commit ..... # eb only deploys commited code
$ sh scripts/deploy.sh
```

### For interested people; How to start a beanstalk?:
https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/go-getstarted.html
1) Create an app in application.go // beanstalk only start code in application.go 
2) `eb init`
3) `eb deploy`

- eb open // open the current env endpoint url