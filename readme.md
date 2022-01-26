## Create a project that calls service created in project1, pass text and prints JSON output returned from the service

# SETUP

Copy `config.sample.env` to `config.env`.

# RUN

`go run .`

# API Test
```
curl --location --request GET 'http://localhost:8081/api/v1/word/occurrence' \
--header 'Content-Type: text/plain' \
--data-raw '10 10 10 10 10 10 10 10 10 10 9 9 9 9 9 9 9 9 9 8 8 8 8 8 8 8 8 7 7 7 7 7 7 7 6 6 6 6 6 6 6 5 5 5 5 5 4 4 4 4 3 3 3 2 2 1 11 11 11 11 11 11 11 11 11 11 11 0'
```