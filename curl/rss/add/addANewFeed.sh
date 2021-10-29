curl "https://api.m3o.com/v1/rss/Add" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "category": "news",
  "name": "bbc",
  "url": "http://feeds.bbci.co.uk/news/rss.xml"
}'