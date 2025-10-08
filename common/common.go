package common

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

type Inject struct {
	V    *Values
	Log  *zap.Logger
	Nats *nats.Conn
	Js   jetstream.JetStream
}

type Values struct {
	Mode      string `yaml:"mode"`
	Namespace string `yaml:"namespace"`
	Nats      Nats   `yaml:"nats"`
}

type Nats struct {
	Hosts []string `yaml:"hosts"`
	Token string   `yaml:"token"`
}

type Job struct {
	Key    string      `msgpack:"key"`
	Index  int         `msgpack:"index"`
	Mode   string      `msgpack:"mode"`
	Option interface{} `msgpack:"option"`
}

type HttpOption struct {
	Method  string                 `json:"method" msgpack:"method"`
	Url     string                 `json:"url" msgpack:"url"`
	Headers map[string]string      `json:"headers" msgpack:"headers"`
	Body    map[string]interface{} `json:"body" msgpack:"body"`
}
