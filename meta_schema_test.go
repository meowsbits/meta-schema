package meta_schema

import (
	"encoding/json"
	"testing"
)

func TestOpenrpc(t *testing.T) {
	original := OpenrpcDocument{}
	b, err := json.MarshalIndent(original, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}
