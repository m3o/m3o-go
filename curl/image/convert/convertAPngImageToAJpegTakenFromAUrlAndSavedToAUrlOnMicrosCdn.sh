curl "https://api.m3o.com/v1/image/Convert" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "name": "cat.jpeg",
  "outputURL": true,
  "url": "somewebsite.com/cat.png"
}'