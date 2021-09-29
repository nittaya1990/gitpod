// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package config

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/gitpod-io/gitpod/image-builder/pkg/orchestrator"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

type Service struct {
	Addr string    `json:"address"`
	TLS  TLSConfig `json:"tls"`
}

type Config struct {
	Orchestrator *orchestrator.Configuration `json:"orchestrator"`
	RefCache     RefCacheConfig              `json:"refCache,omitempty"`
	Service      Service                     `json:"service"`
	Prometheus   Service                     `json:"prometheus"`
	PProf        Service                     `json:"pprof"`
}

type RefCacheConfig struct {
	Interval string   `json:"interval"`
	Refs     []string `json:"refs"`
}

type TLSConfig struct {
	Authority   string `json:"ca"`
	Certificate string `json:"crt"`
	PrivateKey  string `json:"key"`
}

// ServerOption produces the GRPC option that configures a server to use this TLS configuration
func (c *TLSConfig) ServerOption() (grpc.ServerOption, error) {
	if c.Authority == "" || c.Certificate == "" || c.PrivateKey == "" {
		return nil, nil
	}

	// Load certs
	certificate, err := tls.LoadX509KeyPair(c.Certificate, c.PrivateKey)
	if err != nil {
		return nil, xerrors.Errorf("cannot load TLS certificate: %w", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(c.Authority)
	if err != nil {
		return nil, xerrors.Errorf("cannot not read ca certificate: %w", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, xerrors.Errorf("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})

	return grpc.Creds(creds), nil
}
