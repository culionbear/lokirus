<h1 align="center">
  Logrus + Loki = ❤️
</h1>

<h3 align="center">
  A <a href="https://grafana.com/oss/loki/">Loki</a> hook for <a href="https://github.com/Sirupsen/logrus">Logrus</a>

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/culionbear/lokirus/go.yml?branch=main)](https://github.com/culionbear/lokirus/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/culionbear/lokirus)](https://goreportcard.com/report/github.com/culionbear/lokirus)
[![License](https://img.shields.io/github/license/culionbear/lokirus)](https://github.com/culionbear/lokirus/blob/main/LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/culionbear/lokirus?include_prereleases)](https://github.com/culionbear/lokirus/releases)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/culionbear/lokirus)](https://pkg.go.dev/mod/github.com/culionbear/lokirus)

  <img src="demo.gif" />
</h3>

# Installation

```sh
go get github.com/culionbear/lokirus
```

# Usage

```go
package main

func main() {

	// Configure the Loki hook
	opts := lokirus.NewLokiHookOptions().
		// Grafana doesn't have a "panic" level, but it does have a "critical" level
 	 // https://grafana.com/docs/grafana/latest/explore/logs-integration/
		WithLevelMap(lokirus.LevelMap{logrus.PanicLevel: "critical"}).
   		WithFormatter(&logrus.JSONFormatter{}).
		WithStaticLabels(lokirus.Labels{
			"app":         "example",
			"environment": "development",
		}).
		WithBasicAuth("admin", "secretpassword") // Optional

	hook := lokirus.NewLokiHookWithOpts(
		"http://localhost:3100",
		opts,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel)

	// Configure the logger
	logger := logrus.New()
	logger.AddHook(hook)

	// Log all the things!
  logger.WithField("fizz", "buzz").Warnln("warning")
}
```

# Contributing

Contributions are what make the open source community such an amazing place to be, learn, inspire, and create.
Any contributions you make are **greatly appreciated**!

