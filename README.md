# tenant-api


## Build

### Dev
#### Swagger API Documentation setup
Install go swagger (library here). 
Navigate to `./api`, if the folder `swaggerdocs` is present make sure to remove it. Then run `swag init --dir . --output ./swaggerdocs`.

#### Docker

`docker build -f Dockerfile.dev tenantapi:dev .`
`docker-compose -f docker-compose.dev.yml build api`
`docker-compose -f docker-compose.dev.yml up -d`
`docker-compose -f docker-compose.dev.yml down`

### Prod
`docker build -t tenantapi:prod .`
`docker-compose up -d`

`docker-compose down`
#### Linux
`rm -rf db-data` Removes all data from the database
