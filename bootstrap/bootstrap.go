package bootstrap

import (
	"os"
	"strings"

	"github.com/caarlos0/env/v10"
	"github.com/kainonly/worker/common"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

func LoadStaticValues() (values *common.Values, err error) {
	values = new(common.Values)
	if err = env.Parse(values); err != nil {
		return
	}
	return
}

func UseZap() (log *zap.Logger, err error) {
	if os.Getenv("MODE") != "release" {
		if log, err = zap.NewDevelopment(); err != nil {
			return
		}
	} else {
		if log, err = zap.NewProduction(); err != nil {
			return
		}
	}
	return
}

func UseNats(values *common.Values) (nc *nats.Conn, err error) {
	if nc, err = nats.Connect(
		strings.Join(values.Nats.Hosts, ","),
		nats.Token(values.Nats.Token),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(-1),
	); err != nil {
		return
	}
	return
}

func UseJetStream(nc *nats.Conn) (jetstream.JetStream, error) {
	return jetstream.New(nc)
}
