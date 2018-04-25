package cache

import (
	"bytes"
	"context"
	"testing"

	"github.com/golang/groupcache"
)

func TestGetter_Get(t *testing.T) {
	content := "testing"

	client := newMockClient()
	client.PutObject("foo", "bar", bytes.NewReader([]byte(content)))

	g := NewGetter(client, "foo", nil)

	var output string
	if err := g.Get(context.Background(), "bar", groupcache.StringSink(&output)); err != nil {
		t.Error(err)
		return
	}

	if output != content {
		t.Errorf("Expect %s, get %s", content, output)
	}
}
