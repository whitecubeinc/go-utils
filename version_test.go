package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppVersion(t *testing.T) {
	appVersion := NewAppVersion("2.0.9")
	assert.Equal(t, false, appVersion.IsSupportedVersion("2.0.9.2"))

	appVersion = NewAppVersion("2.0.9.2")
	assert.Equal(t, true, appVersion.IsSupportedVersion("2.0.9"))

	appVersion = NewAppVersion("2.0.10")
	assert.Equal(t, true, appVersion.IsSupportedVersion("2.0.9.2"))

	appVersion = NewAppVersion("2.0.9.1")
	assert.Equal(t, false, appVersion.IsSupportedVersion("2.0.9.2"))

	appVersion = NewAppVersion("2.0.9.3")
	assert.Equal(t, true, appVersion.IsSupportedVersion("2.0.9.2"))

	appVersion = NewAppVersion("")
	assert.Equal(t, true, appVersion.IsSupportedVersion("2.0.9.2"))
}
