package xpath

import (
	"io/ioutil"
	"log"
	"os"
)

var l *log.Logger

func init() {
	if os.Getenv("XPATH_LOG") == "debug" {
		l = log.New(os.Stderr, "", log.LstdFlags)
	} else {
		l = log.New(ioutil.Discard, "", log.LstdFlags)
	}
}
