package main

import (
	"crypto/tls"
	"github.com/premnalla/cmd/clustercontrol-api-example/opertest"
	"net/http"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	opertest.AuthenticateWithCmon()

	opertest.Discovery()
}
