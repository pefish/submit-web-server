package controller

import (
	"fmt"
	go_config "github.com/pefish/go-config"
	_type "github.com/pefish/go-core-type/api-session"
	go_error "github.com/pefish/go-error"
	go_file "github.com/pefish/go-file"
	go_time "github.com/pefish/go-time"
	"path"
	"time"
)

type DataControllerClass struct {
}

var DataController = DataControllerClass{}

type SubmitParams struct {
	Desc string `json:"desc" validate:"required"`
	Data string `json:"data" validate:"required" desc:""`
}

func (lc *DataControllerClass) Submit(apiSession _type.IApiSession) (interface{}, *go_error.ErrorInfo) {
	params := SubmitParams{}
	apiSession.MustScanParams(&params)

	filename := fmt.Sprintf("%s_%s.txt", params.Desc, go_time.TimeInstance.TimeToStr(time.Now(), "0000-00-00 00:00:00"))
	err := go_file.FileInstance.WriteFile(path.Join(go_config.ConfigManagerInstance.MustGetString("data-dir"), filename), []byte(params.Data))
	if err != nil {
		return nil, go_error.Wrap(err)
	}
	return true, nil
}
