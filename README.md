# rest-go
Simple REST API Server on golang

## Start
You can specify server host and port by command line flags `-host` `-port`  
For example: `go run . -host=localhost -port=3232`  
Default values if omit them: `-host=localhost -port=5000`

## Endpoints
### `/products`
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

### `/images`
static file endpoint
supports methods:
- GET  `/images/filename` to get file
- HEAD `/images/filename` to get file info