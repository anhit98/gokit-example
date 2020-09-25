package todo

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetAllForUserEndPoint: MakeGetAllForUserEndpoint(s),
		GetByIDEndpoint:       MakeGetByIDEndpoint(s),
		CreateEndpoint:        MakeCreateEndpoint(s),
		UpdateEndpoint:        MakeUpdateEndpoint(s),
		DeleteEndpoint:        MakeDeleteEndpoint(s),
	}
}

func MakeGetAllForUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllForUserRequest)
		todos, err := s.GetAllForUser(req.Username)
		return GetAllForUserResponse{todos}, err
	}
}

func MakeGetByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		todo, err := s.GetByID(req.ID)
		return GetByIDResponse{todo}, err
	}
}

func MakeCreateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := s.Create(req)
		return CreateResponse{id}, err
	}
}

func MakeUpdateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		err := s.Update(req.ID, req.Todo)
		return UpdateResponse{}, err
	}
}

func MakeDeleteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(req.ID)
		return DeleteResponse{}, err
	}
}
