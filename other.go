package main

import (
	_ "golang.org/x/net/http2"
	_ "golang.org/x/oauth2"

	_ "github.com/BurntSushi/toml"
	_ "github.com/envoyproxy/go-control-plane/pkg/util"
	_ "github.com/etcd-io/etcd/client"
	_ "github.com/ethereum/go-ethereum"
	_ "github.com/go-playground/validator"
	_ "github.com/go-yaml/yaml"
	_ "github.com/golang/protobuf/proto"
	_ "github.com/google/uuid"

	// _ "github.com/hashicorp/terraform"
	_ "github.com/mitchellh/mapstructure"
	_ "github.com/plutov/paypal"
	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/sirupsen/logrus"
	_ "go.opentelemetry.io/otel"
	_ "go.uber.org/zap"
	_ "k8s.io/api/core/v1"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/kubernetes"
)
