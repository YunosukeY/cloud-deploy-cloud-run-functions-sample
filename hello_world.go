package helloworld

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloGet", helloGet)
}

func helloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World2!")
}
