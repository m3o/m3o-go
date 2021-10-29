curl "https://api.m3o.com/v1/answer/Question" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "query": "microsoft"
}'