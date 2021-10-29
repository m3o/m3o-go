curl "https://api.m3o.com/v1/qr/Generate" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "size": 300,
  "text": "https://m3o.com/qr"
}'