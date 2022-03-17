package runpersoncreate

import (
	"context"
	"your/path/project/domain_mydomain/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type runPersonCreateInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runPersonCreateInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runPersonCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	// code your usecase definition here ...

	personObj, err := entity.NewPerson(req.Name, req.Age)
	if err != nil {
		return nil, err
	}

	err = r.outport.SavePerson(ctx, personObj)
	if err != nil {
		return nil, err
	}

	//!

	return res, nil
}
