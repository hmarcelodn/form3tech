package test

import "testing"

func TestFetchExistingAccount(t *testing.T) {
	t.Cleanup(func() {
		Truncate()
	})
}
