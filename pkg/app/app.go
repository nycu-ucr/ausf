package app

import (
	ausf_context "github.com/nycu-ucr/ausf/internal/context"
	"github.com/nycu-ucr/ausf/pkg/factory"
)

type App interface {
	SetLogEnable(enable bool)
	SetLogLevel(level string)
	SetReportCaller(reportCaller bool)

	Start()
	Terminate()

	Context() *ausf_context.AUSFContext
	Config() *factory.Config
}
