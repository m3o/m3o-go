curl "https://api.m3o.com/v1/user/UpdatePassword" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "confirmPassword": "myEvenMoreSecretPass123",
  "id": "usrid-1",
  "newPassword": "myEvenMoreSecretPass123",
  "oldPassword": "mySecretPass123"
}'