curl "https://api.m3o.com/v1/postcode/Validate" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "postcode": "SW1A 2AA"
}'