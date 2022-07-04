There are two ways to start the service you can use any on of them
## 1. Starting the service from code build
### starting key-value service
``make up``
### stopping key-value service
``make down``

## 2. Starting the service using docker compose file
``docker-compose up ``


### API Endpoints:
1. Getting value of key
```azure
    Request Endpoint: http://localhost:8080/api/get/<key>
    Request Type: GET
```
2. Set the value of key
```azure
    Request Endpoint: http://localhost:8080/api/set
    Request Type: POST
    Request Body: {
        "Key": "abc-1",
        "Value":"1-abc"
    }
```
3. Search for Keys
- http://localhost:8080/api/search  → Search for keys using the following filters 
  - Assume you have keys: abc-1, abc-2, xyz-1, xyz-2 
  - /search?prefix=abc would return abc-1 and abc-2 
  - /search?suffix=-1 would return abc-1 and xyz-1
```azure
    Request Endpoint: http://localhost:8080/api/search?prefix=abc
    Request Type: GET
```


### Prometheus Metrics Endpoint
- http://localhost:8080/metrics