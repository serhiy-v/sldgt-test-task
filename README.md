# REST API Service for Test Task

Service for credit cards validation

## Starting Up

1.**Create docker image and run it**

```bash
   docker build --tag sldgt-test . 

   docker run --name name -p 8080:8080 sldgt-test:latest
 
```

## Request example

```bash
curl -X POST -H 'Content-Type: application/json' -d '{"Card number": 4111111111111111,"Expiration month": 12,"Expiration year": 2021}' http://localhost:8080/verify
```

## JSON data

```json
{
    "Card number": 4111111111111111,
    "Expiration month": 12,
    "Expiration year": 2024
}
```