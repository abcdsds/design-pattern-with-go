package singleton

import "testing"

func TestScreenConfig(t *testing.T) {
	setUpScreenConfig(10, 10)
	setUpScreenConfig(10, 10)
	setUpScreenConfig(10, 10)

	// Output:
	// setup screen config
}
