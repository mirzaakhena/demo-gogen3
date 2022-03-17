package getallperson

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type getAllPersonInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllPersonInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllPersonInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	// code your usecase definition here ...

	personObjs, err := r.outport.FindAllPerson(ctx, "personID")
	if err != nil {
		return nil, err
	}

	for _, obj := range personObjs {
		fmt.Printf("%v\n", obj)
	}

	res.Items = personObjs

	return res, nil
}
