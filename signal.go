package autonomous

import (
	"os"
	"os/signal"
)

func HandleIntercept(end func()) {
	// Intercept sigterm
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	// Block on this channel
	/*sig*/ _ = <-signalChannel

	end()
}
