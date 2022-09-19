package config_test

import (
	"io"
	"os"
	"pb-dropbox-downloader/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateConfig(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "conf")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	conf, err := config.Load("", file.Name())
	require.NoError(t, err)

	conf.AccessToken = "accesstoken_value"

	err = conf.Save()
	require.NoError(t, err)

	b, err := io.ReadAll(file)
	require.NoError(t, err)

	assert.Contains(t, string(b), "accesstoken_value")
}
