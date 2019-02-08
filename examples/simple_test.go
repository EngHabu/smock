package examples

import (
	"github.com/stretchr/testify/assert"
	"smock/examples/mocks"
	"testing"
)

func TestSample_GetAbc(t *testing.T) {
	s := mocks.Simple{}
	s.OnGetAbc().Return(5, nil)

	i, err := s.GetAbc()
	assert.Equal(t, 5, i)
	assert.NoError(t, err)
}
