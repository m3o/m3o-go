curl "https://api.m3o.com/v1/twitter/Timeline" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "limit": 1,
  "username": "m3oservices"
}'