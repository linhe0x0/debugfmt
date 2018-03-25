package debug_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/apex/log"
	"github.com/sqrthree/debug"
)

func Test(t *testing.T) {
	var buf bytes.Buffer

	log.SetHandler(debug.New(&buf))
	log.WithField("address", "http://localhost:3000").WithField("foo", "bar").Info("hello")
	log.WithField("foo", "bar").Info("hello")
	log.WithField("foo", "bar").Warn("holy guacamole")
	log.WithField("foo", "bar").Error("boom")

	io.Copy(os.Stdout, &buf)
}
