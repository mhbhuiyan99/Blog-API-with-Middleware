
**Loads env variables from a .env file:**
`go get github.com/joho/godotenv`
<br>

**Extensions to golang's database/sql:**
`go get github.com/jmoiron/sqlx`
<br>

**Go PostgreSQL driver for database/sql:**
`go get github.com/lib/pq@latest`
<br>

**Generate slug:**
`go get -u github.com/gosimple/slug`

**Password hashing (bcrypt):**
`go get golang.org/x/crypto/bcrypt`

**Database (sql) Migration:**
`go get -v github.com/rubenv/sql-migrate/...`

**Sends some load to a web application:**
`go install github.com/rakyll/hey@latest`

**Rate Limiting:**
`go get golang.org/x/time/rate@latest`
<br>
Test the rate limiting (Send 8 rapid requests): <br>

`$ for i in {1..8}; do curl -s -w "Request $i: %{http_code}\n" http://localhost:4000/posts; done`

`$ for i in {1..8}; do curl -s -w "Request $i: %{http_code}\n" http://localhost:4000/posts & done; wait`