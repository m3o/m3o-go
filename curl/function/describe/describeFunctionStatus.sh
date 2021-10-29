curl "https://api.m3o.com/v1/function/Describe" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "name": "my-first-func",
  "project": "tests"
}'