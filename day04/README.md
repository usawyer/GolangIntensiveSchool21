## Ex00
Run in two different terminal sessions <br>
`make run_ex00` <br>
`make test_ex00`

### Note
The candy server was generated using Swagger 2.0 with the following steps: <br>

1. Go-swagger installing <br>
   `brew tap go-swagger/go-swagger` <br>
   `brew install go-swagger`

2. Generate an API server, using the following spec: **./api/swagger.yml** <br>
   `swagger generate server -f ./api/swagger.yml  -A candies`

## Ex01
Run in two different terminal sessions <br>
`make run_ex01` <br>
`make test_ex01`

### Note
Certificates were generated using minica with the following steps: <br>

1. Installation <br>
   `go install github.com/jsha/minica@latest` <br>

2. Generate a root key and cert in minica-key.pem, and minica.pem,
   then generate and sign an end-entity key and cert for server and for client <br>

   `minica -ip-addresses 127.0.0.1 -domains server` <br>
   `minica  -domains client`

## Ex02
Run in two different terminal sessions <br>
`make run_ex02` <br>
`make test_ex02`