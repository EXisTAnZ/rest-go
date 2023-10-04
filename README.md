# rest-go
Simple REST API Server on golang

### `/products` endpoint
product JSON
```json 
{
  "id": "string",
  "title": "string",
  "description": "string",
  "price": "number",
  "count": "number",
  "image": "string, URL"
}
```
supports methods:
- GET `/products` for get all products
- GET `/products/id` for get product by id
- POST `/products` with product in body for create new product
- PATCH `/products/id` with product in body for update product by id
- DELETE `/products/id` for delete product by id

### `/images` endpoint
static file endpoint
supports methods:
- GET  `/images/filename` to get file
- HEAD `/images/filename` to get file info