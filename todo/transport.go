package todo

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

var ErrMissingParam = errors.New("Missing parameter")

func MakeHTTPHandler(endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/todo").Handler(httptransport.NewServer(
		endpoints.GetAllForUserEndPoint,
		decodeGetRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/todo/{id}").Handler(httptransport.NewServer(
		endpoints.GetByIDEndpoint,
		decodeGetByIDRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/todo").Handler(httptransport.NewServer(
		endpoints.CreateEndpoint,
		decodeCreateRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/todo/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateEndpoint,
		decodeUpdateRequest,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/todo/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteEndpoint,
		decodeDeleteRequest,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func decodeGetRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req GetAllForUserRequest
	username := r.URL.Query().Get("username")

	req = GetAllForUserRequest{
		Username: username,
	}
	return req, nil
}

func decodeGetByIDRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req GetByIDRequest
	vars := mux.Vars(r)

	req = GetByIDRequest{
		ID: vars["id"],
	}
	return req, err
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, err
}

func decodeUpdateRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, ErrMissingParam
	}
	var todo Todo
	err = render.Decode(r, &todo)
	if err != nil {
		return nil, err
	}
	return UpdateRequest{id, todo}, err
}

func decodeDeleteRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, ErrMissingParam
	}
	return DeleteRequest{id}, err
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok && err != nil {
		encodeError(ctx, err, w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrInconsistentIDs, ErrMissingParam:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
