//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/kainonly/worker/app"
	"github.com/kainonly/worker/common"
)

func NewApp() (*app.App, error) {
	wire.Build(
		wire.Struct(new(common.Inject), "*"),
		LoadStaticValues,
		UseZap,
		UseNats,
		UseJetStream,
		app.Initialize,
	)
	return &app.App{}, nil
}
