curl "https://api.m3o.com/v1/user/Read" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "email": "joe@example.com"
}'