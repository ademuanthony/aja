package endpoint

import (
	"context"
	service "github.com/ademuanthony/aja/bugs/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	S string `json:"s"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.BugsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		e0 := s.Create(ctx, req.S)
		return CreateResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, s string) (e0 error) {
	request := CreateRequest{S: s}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).E0
}