package main

import (
	"encoding/json"
	"errors"
	"fmt"
	stdhttp "net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

var (
	errBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func makeHandler(ctx context.Context, s CategoryService, logger kitlog.Logger) stdhttp.Handler {
	e := makeEndpoints(s)
	r := mux.NewRouter()

	commonOptions := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/").HandlerFunc(index)

	// GET     /categories/load
	r.Methods("GET").Path("/categories/load").Handler(kithttp.NewServer(
		ctx,
		e.loadCategoriesFileEndpoint,
		decodeLoadCategoriesFileRequest,
		encodeResponse,
		commonOptions...,
	))

	// GET     /categories/:id/subcategories             retrieve subcategories associated with the category
	r.Methods("GET").Path("/categories").Handler(kithttp.NewServer(
		ctx,
		e.getCategoriesEndpoint,
		decodeGetCategoriesRequest,
		encodeResponse,
		commonOptions...,
	))

	// GET     /categories/:id/subcategories             retrieve subcategories associated with the category
	r.Methods("GET").Path("/categories/{id}/subcategories").Handler(kithttp.NewServer(
		ctx,
		e.getSubCategoriesEndpoint,
		decodeGetSubCategoriesRequest,
		encodeResponse,
		commonOptions...,
	))

	return r
}

func index(w stdhttp.ResponseWriter, req *stdhttp.Request) {
	fmt.Fprintf(w, "OK!")
}

func decodeLoadCategoriesFileRequest(r *stdhttp.Request) (request interface{}, err error) {
	return loadCategoriesFileRequest{}, nil
}

func decodeGetSubCategoriesRequest(r *stdhttp.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRouting
	}
	return getSubCategoriesRequest{CategoryID: id}, nil
}

func decodeGetCategoriesRequest(r *stdhttp.Request) (request interface{}, err error) {
	return getCategoriesRequest{}, nil
}

// errorer is implemented by all concrete response types. It allows us to
// change the HTTP response code without needing to trigger an endpoint
// (transport-level) error. For more information, read the big comment in
// endpoint.go.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because I didn't know if something more
// specific was necessary. It's certainly possible to specialize on a
// per-response (per-method) basis.
func encodeResponse(w stdhttp.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(w, e.error())
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(w stdhttp.ResponseWriter, err error) {
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case nil:
		return stdhttp.StatusOK
	case errNotFound:
		return stdhttp.StatusNotFound
	case errAlreadyExists, errInconsistentIDs:
		return stdhttp.StatusBadRequest
	default:
		if _, ok := err.(kithttp.BadRequestError); ok {
			return stdhttp.StatusBadRequest
		}
		return stdhttp.StatusInternalServerError
	}
}
