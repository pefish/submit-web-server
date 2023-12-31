package command

import (
	"flag"
	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	"github.com/pefish/go-core/driver/logger"
	global_api_strategy "github.com/pefish/go-core/global-api-strategy"
	"github.com/pefish/go-core/service"
	go_logger "github.com/pefish/go-logger"
	task_driver "github.com/pefish/go-task-driver"
	"github.com/pefish/submit-web-server/pkg/constant"
	"github.com/pefish/submit-web-server/pkg/global"
	"github.com/pefish/submit-web-server/pkg/route"
	"github.com/pefish/submit-web-server/version"
)

type DefaultCommand struct {
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	flagSet.String("host", "127.0.0.1", "The host of web server.")
	flagSet.Int("port", 8000, "The port of web server.")
	return nil
}

func (dc *DefaultCommand) OnExited(data *commander.StartData) error {
	return nil
}

func (dc *DefaultCommand) Init(data *commander.StartData) error {
	service.Service.SetName(version.AppName)
	logger.LoggerDriverInstance.Register(go_logger.Logger)

	err := go_config.ConfigManagerInstance.Unmarshal(&global.GlobalConfig)
	if err != nil {
		return err
	}

	service.Service.SetHost(global.GlobalConfig.Host)
	service.Service.SetPort(global.GlobalConfig.Port)
	service.Service.SetPath(`/api`)
	global_api_strategy.ParamValidateStrategyInstance.SetErrorCode(constant.PARAM_ERROR)

	service.Service.SetRoutes(route.UserRoute)

	return nil
}

func (dc *DefaultCommand) Start(data *commander.StartData) error {

	taskDriver := task_driver.NewTaskDriver()
	taskDriver.Register(service.Service)

	taskDriver.RunWait(data.ExitCancelCtx)

	return nil
}
