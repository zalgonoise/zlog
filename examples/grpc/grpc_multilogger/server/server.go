package main

import (
	"os"

	"github.com/zalgonoise/zlog/log"

	"github.com/zalgonoise/zlog/grpc/server"
)

func getEnv(env string) (val string, ok bool) {
	v := os.Getenv(env)

	if v == "" {
		return v, false
	}

	return v, true
}

func getTLSConf() server.LogServerConfig {
	var tlsConf server.LogServerConfig

	var withCert bool
	var withKey bool
	var withCA bool

	certPath, ok := getEnv("TLS_SERVER_CERT")

	if ok {
		withCert = true
	}

	keyPath, ok := getEnv("TLS_SERVER_KEY")

	if ok {
		withKey = true
	}

	caPath, ok := getEnv("TLS_CA_CERT")

	if ok {
		withCA = true
	}

	if withCert && withKey {
		if withCA {
			tlsConf = server.WithTLS(certPath, keyPath, caPath)
		} else {
			tlsConf = server.WithTLS(certPath, keyPath)
		}
	}

	return tlsConf
}

func main() {
	// set up two servers for demo purposes

	grpcLoggerAlpha := server.New(
		server.WithLogger(
			log.New(
				log.WithFormat(log.TextColorLevelFirstSpaced),
			),
		),
		server.WithServiceLogger(
			log.New(
				log.WithFormat(log.TextColorLevelFirstSpaced),
			),
		),
		server.WithAddr("127.0.0.1:9299"),
		server.WithGRPCOpts(),
		getTLSConf(), // loaded from provided env variables
	)

	grpcLoggerBeta := server.New(
		server.WithLogger(
			log.New(
				log.WithFormat(log.TextColorLevelFirstSpaced),
			),
		),
		server.WithServiceLogger(
			log.New(
				log.WithFormat(log.TextColorLevelFirstSpaced),
			),
		),
		server.WithAddr("127.0.0.1:9399"),
		server.WithGRPCOpts(),
		getTLSConf(), // loaded from provided env variables
	)

	go grpcLoggerAlpha.Serve()
	grpcLoggerBeta.Serve()

}
