package utils

import (
	"testing"
)

func TestListContainers(t *testing.T) {
	containers, err := ListContainers()
	if err != nil {
		t.Error(err)
	}

	// output all the continers
	for _, c := range containers {
		t.Logf("container: %s", c.ID)
	}
}
