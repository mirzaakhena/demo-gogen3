package getallperson

import (
	"context"
	"your/path/project/domain_mydomain/model/entity"
)

// mirza here

// Inport of Usecase
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase
type InportRequest struct {
	Page int
	Size int
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
	Count int
	Items []*entity.Person
}

func (r InportRequest) Validate() error {
	return nil
}
