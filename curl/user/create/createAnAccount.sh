curl "https://api.m3o.com/v1/user/Create" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "email": "joe@example.com",
  "id": "usrid-1",
  "password": "mySecretPass123",
  "username": "usrname-1"
}'