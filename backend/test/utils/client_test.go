package utils

import "testing"

func Test_ConnectDB(t *testing.T) {
	_ = NewTestClient(t)
}
