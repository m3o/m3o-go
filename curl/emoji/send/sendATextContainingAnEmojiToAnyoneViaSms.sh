curl "https://api.m3o.com/v1/emoji/Send" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "from": "Alice",
  "message": "let's grab a :beer:",
  "to": "+44782669123"
}'