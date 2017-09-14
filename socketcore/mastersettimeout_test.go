package gluasocket_socketcore_test

import (
	"testing"

	"github.com/BixData/gluasocket"
	"github.com/stretchr/testify/assert"
	"github.com/yuin/gopher-lua"
)

func TestMasterSetTimeout(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	assert.NoError(L.DoString(`return require 'socket'.tcp():settimeout(.25)`))
	retval := L.Get(-1)
	assert.Equal(lua.LTNil, retval.Type())
}