// Code generated by goa v3.4.3, DO NOT EDIT.
//
// chatapi endpoints
//
// Command:
// $ goa gen chat-api/design

package chatapi

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "chatapi" service endpoints.
type Endpoints struct {
	Getchat goa.Endpoint
}

// NewEndpoints wraps the methods of the "chatapi" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Getchat: NewGetchatEndpoint(s),
	}
}

// Use applies the given middleware to all the "chatapi" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Getchat = m(e.Getchat)
}

// NewGetchatEndpoint returns an endpoint function that calls the method
// "getchat" of service "chatapi".
func NewGetchatEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetchatPayload)
		res, err := s.Getchat(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGoaChatCollection(res, "default")
		return vres, nil
	}
}
