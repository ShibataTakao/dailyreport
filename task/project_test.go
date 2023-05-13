package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectEquals(t *testing.T) {
	tests := []struct {
		name   string
		p1     Project
		p2     Project
		equals bool
	}{
		{
			name:   "Equals",
			p1:     NewProject("Project A"),
			p2:     NewProject("Project A"),
			equals: true,
		},
		{
			name:   "NotEquals",
			p1:     NewProject("Project A"),
			p2:     NewProject("Project B"),
			equals: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.equals, tt.p1.Equals(tt.p2))
		})
	}
}
