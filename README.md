## Beer App Backend

## Construir
docker-compose up --build

## Levantar aplicación
docker-compose up -d

## Cambiar el api key en app/config, si es que ya venció
"currency": {
    "url": "http://apilayer.net/api/live",
    "key": "3c6df49ec70f01acddc0ae0d11aa4f43"
  }

## Si no hay conexión a bd cambiar app/config/test-config.json, cambiar host:localhost o host:database

  "database": {
    "host": "database",
    "host_docker": "database",
    "port": "3306",
    "user": "test",
    "pass": "test",
    "name": "beerapp"
  },