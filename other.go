package main

import (
	_ "golang.org/x/net/http2"

	_ "github.com/BurntSushi/toml"
	_ "github.com/Sirupsen/logrus"
	_ "github.com/etcd-io/etcd"
	_ "github.com/ethereum/go-ethereum"
	_ "github.com/go-playground/validator"
	_ "github.com/go-yaml/yaml"
	_ "github.com/golang/oauth2"
	_ "github.com/golang/protobuf"
	_ "github.com/google/uuid"
	_ "github.com/hashicorp/consul"
	_ "github.com/hashicorp/terraform"
	_ "github.com/kubernetes/kubernetes"
	_ "github.com/mitchellh/mapstructure"
	_ "github.com/open-telemetry/opentelemetry-go"
	_ "github.com/plutov/paypal"
	_ "github.com/prometheus/prometheus"
	_ "github.com/uber-go/zap"
)
