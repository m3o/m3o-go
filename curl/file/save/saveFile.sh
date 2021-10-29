curl "https://api.m3o.com/v1/file/Save" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "file": {
    "content": "file content example",
    "path": "/document/text-files/file.txt",
    "project": "examples"
  }
}'