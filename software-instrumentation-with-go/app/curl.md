# CURL

## SEED

```bash
curl -X GET http://localhost:8080/api/v1/course/seed \
  -H "Content-Type: application/json" 
```

## CLEAN

```bash
curl -X GET http://localhost:8080/api/v1/course/clean \
  -H "Content-Type: application/json" 
```

## CREATE BOOKING

```bash
curl -X POST http://localhost:8080/api/v1/booking \
  -H "Content-Type: application/json" \
  -d '{
    "batch_code": "GIT-BATCH-01",
    "customer_name": "John Doe"
  }'
```