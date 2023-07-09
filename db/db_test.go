package db

import "testing"

func TestConn(t *testing.T) {
	t.Run("initial", func(t *testing.T) {
		ConnectionDB()
	})
}
