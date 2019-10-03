# hoproxy

## Usage

## Example

- config.yml

`````
server:
    port: 8088
exchange:
  - method: POST
    path: /api/post
    call_cli: /usr/local/bin/post.sh
  - method: GET
    path: /search
    call_uri: https://www.google.com/search
  - method: DELETE
    path: /api/delete
    call_cli: /usr/local/bin/delete.sh
  - method: PUT
    path: /api/put
    call_uri: https://example.com/api/put

`````

start server

`````
$ hoproxy -conf config.yml
`````

