package todo

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetAllForUserEndPoint endpoint.Endpoint
	GetByIDEndpoint       endpoint.Endpoint
	CreateEndpoint        endpoint.Endpoint
	UpdateEndpoint        endpoint.Endpoint
	DeleteEndpoint        endpoint.Endpoint
}

type GetAllForUserRequest struct {
	Username string
}

type GetAllForUserResponse struct {
	Todos []Todo `json:"todos"`
}

type GetByIDRequest struct {
	ID string
}

type GetByIDResponse struct {
	Todo Todo `json:"todo"`
}

type CreateRequest struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

type UpdateRequest struct {
	ID   string
	Todo Todo
}

type UpdateResponse struct {
}

type DeleteRequest struct {
	ID string
}

type DeleteResponse struct {
}
