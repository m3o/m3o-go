curl "https://api.m3o.com/v1/cache/Set" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "key": "foo",
  "value": "bar"
}'