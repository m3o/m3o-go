curl "https://api.m3o.com/v1/currency/Convert" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "from": "USD",
  "to": "GBP"
}'