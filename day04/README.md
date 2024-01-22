## Ex00
Task is to generate Go code based on a given Swagger spec that implements 
a candy vending machine server. The server should handle candy orders through 
HTTP and JSON, validating input and calculating change.

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
Implement certificate authentication for the candy vending machine server. Use 
a self-signed certificate and a local certificate authority (minica) for secure 
communication. Modify the server code to support secure URLs. Create a test 
client that can query the API using the self-signed certificate and demonstrate 
the ability to verify it on both sides.

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
Integrate an existing C code that generates ASCII art of a cow saying funny things 
into the candy vending machine server response. The C code must be reused without 
modification. The goal is to return the ASCII art of the cow as part of the "thanks" 
field in the API response.

Run in two different terminal sessions <br>
`make run_ex02` <br>
`make test_ex02`
