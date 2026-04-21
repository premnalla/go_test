package opertest

import (
	"fmt"
	"os"

  "github.com/premnalla/go_test/pkg/openapi"
)

func listTemplates() {
	config := *openapi.NewConfig("listTemplates")
	resp, err := gApiClient.ConfigAPI.ConfigPost(gNewCtx).Config(config).Execute()
	if err != nil {
		printError(err, resp)
		return
	}
	fmt.Fprintf(os.Stderr, "Resp `ConfigPost.listTemplates`: %v\n", resp)
}
