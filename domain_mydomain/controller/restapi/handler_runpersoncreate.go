package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"your/path/project/domain_mydomain/usecase/runpersoncreate"
	"your/path/project/shared/infrastructure/logger"
	"your/path/project/shared/infrastructure/util"
	"your/path/project/shared/model/payload"
)

// runPersonCreateHandler ...
func (r *Controller) runPersonCreateHandler(inputPort runpersoncreate.Inport) gin.HandlerFunc {

	type request struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req runpersoncreate.InportRequest
		req.Name = jsonReq.Name
		req.Age = jsonReq.Age

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
