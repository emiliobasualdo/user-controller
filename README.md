
## levantar un beanstalk: 
https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/go-getstarted.html
1) meter código en application.go
2) eb init
3) eb deploy

- eb open // open the current env endpoint url

## asociar el beanstalk a un api-gateway
1) levantanar un Load Balancer para el app : https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-nlb-for-vpclink-using-console.html

## deploy
eb solo deploya código que haya sido commiteado
1) git commit ...
2) eb deploy
