curl "https://api.m3o.com/v1/user/Update" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "email": "joeotheremail@example.com",
  "id": "usrid-1"
}'