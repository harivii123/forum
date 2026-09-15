
### Table relationship
Users → Posts                   1:N (creates)
Posts → Comments                1:N
Users → Comments                1:N
Users → Sessions                1:N

Users → Post votes              1:N
Posts → Post votes              1:N
Users → Comment votes           1:N
Comments → Comment votes        1:N
Posts ↔ Categories              N:N
Posts → Post categories         1:N
Categories → Post categories    1:N


### Routes
GET  /                     → get all posts

GET  /register             → register page
POST /register             → register

GET  /login                → login page
POST /login                → login
POST /logout               → logout

GET  /posts/create         → create form
POST /posts/create         → create post

GET  /posts/{id}           → get post + its comments
POST /posts/{id}/comments  → comment

POST /posts/{id}/vote      → like/dislike post
POST /comments/{id}/vote   → like/dislike comment

If you want to seed database, use this:
```
go run -tags "fts5" . -seed 
```

### Login 
LOGIN  

email + password  
       ↓  
SELECT user by email  
       ↓  
bcrypt compare  
       ↓  
password correct?  
       ↓ yes  
generate UUID  
       ↓  
INSERT session into DB  
       ↓  
put UUID in cookie  
       ↓  
browser stores cookie  
       ↓  
redirect /  

For the next request:

    BROWSER
    Cookie: session_id=UUID
            ↓
    SERVER
    read cookie
            ↓
    SELECT session WHERE id = UUID
            ↓
    get user_id
            ↓
    "Ah, this is user 3"