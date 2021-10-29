curl "https://api.m3o.com/v1/sms/Send" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "from": "Alice",
  "message": "Hi there!",
  "to": "+447681129"
}'