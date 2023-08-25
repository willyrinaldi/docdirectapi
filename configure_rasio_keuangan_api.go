// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"directApi/restapi/operations"
	"directApi/restapi/operations/balance_sheet_data"
	"directApi/restapi/operations/data_history"
	"directApi/restapi/operations/income_statement"
	"directApi/restapi/operations/index_data"
	"directApi/restapi/operations/informasi_pasar"
	"directApi/restapi/operations/rasio_keuangan"
)

//go:generate swagger generate server --target ../../APIDirect --name RasioKeuanganAPI --spec ../swagger.json --principal interface{}

func configureFlags(api *operations.RasioKeuanganAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.RasioKeuanganAPIAPI) http.Handler {
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

	if api.DataHistoryGetDirectV1SahamDpEqHandler == nil {
		api.DataHistoryGetDirectV1SahamDpEqHandler = data_history.GetDirectV1SahamDpEqHandlerFunc(func(params data_history.GetDirectV1SahamDpEqParams) middleware.Responder {
			return middleware.NotImplemented("operation data_history.GetDirectV1SahamDpEq has not yet been implemented")
		})
	}
	if api.IndexDataGetDirectV1SahamDpIxHandler == nil {
		api.IndexDataGetDirectV1SahamDpIxHandler = index_data.GetDirectV1SahamDpIxHandlerFunc(func(params index_data.GetDirectV1SahamDpIxParams) middleware.Responder {
			return middleware.NotImplemented("operation index_data.GetDirectV1SahamDpIx has not yet been implemented")
		})
	}
	if api.InformasiPasarGetDirectV1SahamFdrIPGetInformasiPasarDataHandler == nil {
		api.InformasiPasarGetDirectV1SahamFdrIPGetInformasiPasarDataHandler = informasi_pasar.GetDirectV1SahamFdrIPGetInformasiPasarDataHandlerFunc(func(params informasi_pasar.GetDirectV1SahamFdrIPGetInformasiPasarDataParams) middleware.Responder {
			return middleware.NotImplemented("operation informasi_pasar.GetDirectV1SahamFdrIPGetInformasiPasarData has not yet been implemented")
		})
	}
	if api.RasioKeuanganGetDirectV1SahamFdrRkGetRasioKeuanganDataHandler == nil {
		api.RasioKeuanganGetDirectV1SahamFdrRkGetRasioKeuanganDataHandler = rasio_keuangan.GetDirectV1SahamFdrRkGetRasioKeuanganDataHandlerFunc(func(params rasio_keuangan.GetDirectV1SahamFdrRkGetRasioKeuanganDataParams) middleware.Responder {
			return middleware.NotImplemented("operation rasio_keuangan.GetDirectV1SahamFdrRkGetRasioKeuanganData has not yet been implemented")
		})
	}
	if api.BalanceSheetDataGetDirectV1SahamLkBsHandler == nil {
		api.BalanceSheetDataGetDirectV1SahamLkBsHandler = balance_sheet_data.GetDirectV1SahamLkBsHandlerFunc(func(params balance_sheet_data.GetDirectV1SahamLkBsParams) middleware.Responder {
			return middleware.NotImplemented("operation balance_sheet_data.GetDirectV1SahamLkBs has not yet been implemented")
		})
	}
	if api.IncomeStatementGetDirectV1SahamLkIsHandler == nil {
		api.IncomeStatementGetDirectV1SahamLkIsHandler = income_statement.GetDirectV1SahamLkIsHandlerFunc(func(params income_statement.GetDirectV1SahamLkIsParams) middleware.Responder {
			return middleware.NotImplemented("operation income_statement.GetDirectV1SahamLkIs has not yet been implemented")
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
