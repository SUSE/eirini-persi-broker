package log

import "code.cloudfoundry.org/lager"

var logger lager.Logger

func Logger() lager.Logger {
	if logger == nil {
		logger = lager.NewLogger("eirini-persi-broker")
	}
	return logger
}
