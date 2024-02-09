package main

import "testing"

func TestReturnChannelName(t *testing.T) {
	actualOutput := ReturnChannelName()
	expectedOutput := "vanshika123"

	if actualOutput != expectedOutput {
		t.Errorf("error")
	}
}
