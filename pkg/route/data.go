package route

import (
	"github.com/pefish/go-core/api"
	"github.com/pefish/go-http/gorequest"
	"github.com/pefish/submit-web-server/pkg/controller"
)

var UserRoute = []*api.Api{
	{
		Description: "",
		Path:        "/v1/submit",
		Method:      gorequest.POST,
		Params:      controller.SubmitParams{},
		Controller:  controller.DataController.Submit,
	},
}
