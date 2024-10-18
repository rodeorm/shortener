package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitialize(t *testing.T) {
	tests := []struct {
		name  string
		level string
	}{
		{
			name:  "проверка создания логгера на уровне",
			level: "info",
		},
		{
			name:  "проверка создания логгера на уровне",
			level: "error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Initialize(tt.level)
			require.NoError(t, err)
		})
	}
}
