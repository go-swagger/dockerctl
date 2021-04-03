package tests

import (
	"testing"

	"github.com/go-openapi/dockerctl/client"
	"github.com/go-openapi/dockerctl/client/image"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func TestCreateImage(t *testing.T) {
	r := httptransport.New("localhost", client.DefaultBasePath, []string{"http"})

	appCli := client.New(r, strfmt.Default)

	_, err := appCli.Image.ImageCreate(image.NewImageCreateParams())

	if err != nil {
		// shows: Unexpected error:  (string) is not supported by the ByteStreamProducer, can be resolved by supporting Reader/BinaryMarshaler interface
		// but expect connection refused.
		t.Errorf("Unexpected error: %+v", err)
	}
}
