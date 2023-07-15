package config

import "testing"

func TestConnection(t *testing.T) {
	t.Run("connection", func(t *testing.T) {
		ConnectionDB()
	})
}
