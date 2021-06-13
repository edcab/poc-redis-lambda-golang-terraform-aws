# POC implementación IoC para administrar Api Serverless AWS

Servicios AWS que implementaremos:
* AWS Lambda
* ElastiCache REDIS
* Rest API Gateway
* CloudWatch Logging

## Prerrequisitos
* Terraform
* Cuenta AWS creada
  * Credenciales AWS configuradas localmente
* Golang

## Construcción de Binarios de las lambdas
El objetivo sera crear en la carpeta /terraform/dist los ejecutables o binarios creados a partir de las lambdas creadas:
$ GOOS=linux go build -v -o ../terraform/dist/poc_<name>/main ..functions/<name_lambda>/main.go

## Instalación
* Comandos basicos
```
$ terraform init
$ terraform plan
$ terraform apply
```




