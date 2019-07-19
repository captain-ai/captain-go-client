package captain

import (
	"context"
	"encoding/json"
	"github.com/go-test/deep"
	"os"
	"testing"
	"time"
)

func newClientFromEnv(t *testing.T) *Client {
	client := NewClient()
	client.IntegrationKey = mustGetenv(t, "CAPTAIN_INTEGRATION_KEY")
	client.DeveloperKey = mustGetenv(t, "CAPTAIN_DEVELOPER_KEY")
	return client
}

func mustGetenv(t *testing.T, name string) string {
	value := os.Getenv(name)
	if value == "" {
		t.Fatalf("missing environment variable: %q", name)
	}
	return value
}

func withTimeout(timeout time.Duration) context.Context {
	ctxWithTimeout, _ := context.WithTimeout(context.Background(), timeout)
	return ctxWithTimeout
}

func testExactJSON(t *testing.T, v interface{}, data []byte) {
	var want interface{}
	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Fatal(err)
	}
	buf, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	var have interface{}
	err = json.Unmarshal(buf, &have)
	if err != nil {
		t.Fatal(err)
	}
	if diff := deep.Equal(want, have); diff != nil {
		t.Fatal(diff)
	}
}