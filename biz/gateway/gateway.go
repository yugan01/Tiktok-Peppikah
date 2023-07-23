package gateway

import (
	"context"
	"fmt"
	"net/http"
	"orbital_demo/biz/errors"
	"orbital_demo/kitex_gen/common"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

var SvcMap = make(map[string]genericclient.Client)

func GenericCall(service string, method string, ctx context.Context, c *app.RequestContext) {

	cli := SvcMap[service]

	fmt.Println("Generic Call execution")

	jsonBody := string(c.Request.BodyBytes())

	//Creating the generic call, this function will be used by router.go in hertz_gateway
	resp, err := cli.GenericCall(ctx, method, jsonBody)

	if err != nil {
		hlog.Errorf("GenericCall err:%v", err)
		bizErr, ok := kerrors.FromBizStatusError(err)
		if !ok {
			c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
			return
		}
		resp = errors.Err{ErrCode: int64(bizErr.BizStatusCode()), ErrMsg: bizErr.BizMessage()}
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
