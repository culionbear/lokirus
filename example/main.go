package main

import (
	"github.com/culionbear/lokirus"
	"github.com/sirupsen/logrus"
)

func main() {

	// Configure the Loki hook
	opts := lokirus.NewLokiHookOptions().
		// Grafana doesn't have a "panic" level, but it does have a "critical" level
		// https://grafana.com/docs/grafana/latest/explore/logs-integration/
		WithLevelMap(lokirus.LevelMap{logrus.PanicLevel: "critical"}).
		WithStaticLabels(lokirus.Labels{
			"app":         "example",
			"environment": "development",
		}).WithFormatter(&logrus.JSONFormatter{})

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
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("foo", "bar").Log(logrus.InfoLevel, "Road work ahead? Uh yea, I sure hope it does.")
}
