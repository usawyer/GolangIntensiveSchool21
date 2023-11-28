// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"candies/restapi/operations"
)

//go:generate swagger generate server --target ../../day04 --name Candies --spec ../api/swagger.yml --principal interface{}

func implement(params operations.BuyCandyParams) middleware.Responder {
	var candies = map[string]int64{"CE": 10, "AA": 15, "NT": 17, "DE": 21, "YR": 23}

	kind := *params.Order.CandyType
	count := *params.Order.CandyCount
	money := *params.Order.Money
	if _, ok := candies[kind]; !ok {
		payload := operations.BuyCandyBadRequestBody{
			Error: "invalid candy type, should be CE/AA/NT/DE/YR",
		}
		return operations.NewBuyCandyBadRequest().WithPayload(&payload)
	}
	need := candies[kind] * count

	if need <= money {
		payload := operations.BuyCandyCreatedBody{
			Change: money - need,
			Thanks: "Thank you!",
		}
		return operations.NewBuyCandyCreated().WithPayload(&payload)

	} else {
		payload := operations.BuyCandyPaymentRequiredBody{
			Error: fmt.Sprintf("You need %d more money!", need-money)}
		return operations.NewBuyCandyPaymentRequired().WithPayload(&payload)
	}
}

func configureFlags(api *operations.CandiesAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CandiesAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BuyCandyHandler = operations.BuyCandyHandlerFunc(implement)
	if api.BuyCandyHandler == nil {
		api.BuyCandyHandler = operations.BuyCandyHandlerFunc(func(params operations.BuyCandyParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.BuyCandy has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
