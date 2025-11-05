## Create Posts

POST http://localhost:3002/posts
Content-Type: application/json

{
  "title": "First Post",
  "content": "Hello World",
  "author": "Eddy",
  "tags": ["go", "fiber"]
}

