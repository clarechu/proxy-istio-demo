package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendMailByServer(t *testing.T) {
	err := SendMailByServer()
	assert.Equal(t, nil, err)
}
