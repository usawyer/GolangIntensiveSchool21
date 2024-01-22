## Ex00
Utilize Go Elasticsearch bindings to create an index, define mappings, 
and upload a dataset of Moscow restaurants into Elasticsearch. Implement 
a Store interface to abstract database operations for fetching paginated 
restaurant entries. Run an HTTP server on port 8888 to display a simple 
HTML UI providing a paginated list of restaurant names, addresses, and 
phones.

## Ex01
Develop an HTML UI for the restaurant database, abstracting the database 
behind a Store interface. The interface should support pagination, and 
the HTML UI should render a list of restaurants along with pagination 
links. Ensure proper error handling for invalid page values.

## Ex02
Implement another HTTP handler that responds with JSON data for the restaurant 
database. This API should also support pagination, and the JSON response should 
include information about the total number of entries, the current page, and 
links for pagination. Properly handle errors when an invalid page is specified.

## Ex03
Implement functionality to search for the three closest restaurants based on 
given coordinates. Configure sorting for the query, and the API should respond 
with JSON containing information about the recommended restaurants.

## Ex04
Implement a JWT-based authentication system. Create an API endpoint to generate 
a token and protect the /api/recommend endpoint with JWT middleware. The middleware 
should check the validity of the token, allowing access only when a valid JWT is 
provided.