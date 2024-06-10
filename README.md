# Synapsis

## Features
- Authentication for register and login
- show product list by category
- add, update & delete product from cart
- checkout cart to order
- confirm payment

## How to Run Locally
```bash
# clone repository
git clone git@github.com:IbnAnjung/sinapsys.git

cd sinapsys

## run with docker
docker-compose up

## run development mode
cp .env.example .env

set -o allexport; source .env; go run main.go

## also you can pull from docker image
docker pull anggasaputra/synapsis:latest
#running container 
docker run --network local -p 8000:8000 --env HTTP_PORT=8000 \ 
  --env DB_USER=root --env DB_PASSWORD=secret --env DB_HOST=mysql \
  --env DB_PORT=3306 --env DB_SCHEMA=synapsis --env DB_TIMEOUT=60 \
  --env DB_MAX_IDDLE_CONNECTION=5 --env DB_MAX_IDDLE_LIFETTIME=10 \
  --env DB_MAX_OPEN_CONNECTION=20 --env DB_MAX_LIFETIME=600 \
  --env JWT_SECRET=secret --env JWT_SELLER_SECRET=seller_secret \
  --env JWT_ACCESS_TOKEN_LIFETIME=24 --env JWT_REFRESH_TOKEN_LIFETIME=72 \
  --env REDIS_ADDR=redis:6379 --env REDIS_USERNAME= --env REDIS_PASSWORD= --env \
  REDIS_DB=0 --env REDIS_MIN_IDLE_CONNECTION=5 --env REDIS_MAX_IDLE_CONNECTION=10 \
  --env REDIS_MAX_ACTIVE_CONNECTION=20 \ 
  --rm --name synapsis synapsis

```
## List Of Endpoint
-  ``[GET] /product`` 
-  ``[POST] /auth/register``
-  ``[POST] /auth/login``
-  ``[GET] /cart``
-  ``[POST] /cart``
-  ``[PUT] /cart/:id``
-  ``[DEL] /cart/:id``
-  ``[POST] /cart/checkout``
-  ``[POST] /payment/manual-transfer/confirm``

## Project Structure
![](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

      .
      ├── build                         # build & ci/cd content 
      ├   ├── http.Dockerfile 
      ├── cmd                           # main / bootstrapping services 
      ├   ├── http                      # http service 
      ├   ├   ├── config                 
      ├   ├   ├── handler                      
      ├   ├   ├── router                 
      ├   ├   ├── server.go             # load all dependency & start http 
      ├── database                      # initiate db migration & seeder
      ├── entity                        # entity layer
      ├   ├── dto                       # for usecase data transfer object                
      ├── pkg                           # helper, driver, library & etc
      ├   ├── cache                
      ├   ├── crypt                
      ├   ├── error                
      ├   ├── http                
      ├   ├── jwt                
      ├   ├── orm                
      ├   ├── redis                
      ├   ├── sql                
      ├   ├── string                
      ├   ├── structvalidator                
      ├   ├── time                
      ├── repository                      # adapter to get data from source
      ├   ├── gorm                        # source SQL with gorm              
      ├   ├   ├── model                            
      ├── usecase                         # bussiness logic
      ├   ├── auth
      ├   ├── cart
      ├   ├── payment
      ├   ├── product
      ├── .env.example                  
      ├── .gitignore                   
      ├── .golangci.yaml                   
      ├── docker-compose.yaml
      ├── go.mod
      ├── go.sum
      ├── main.go
      └── README.md
## Tech Stack
- Golang
- Mysql 8.0
- Redis

## Entity Relationship Diagram
![ERD](./resource/synapsis.erd.png)