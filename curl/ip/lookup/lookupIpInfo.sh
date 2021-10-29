curl "https://api.m3o.com/v1/ip/Lookup" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "ip": "93.148.214.31"
}'