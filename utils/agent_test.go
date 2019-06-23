package utils

import (
	"testing"
)

func TestListContainers(t *testing.T) {
	containers, err := ListContainers()
	if err != nil {
		t.Error(err)
	}

	// output all the containers
	for _, c := range containers {
		t.Logf("container: %s", c.ID)
		t.Log("container name:", c.Names)
	}
}
