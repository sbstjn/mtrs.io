package rating

import "testing"

func TestAverage(t *testing.T) {
	if !KnownFactor("type") {
		t.Error("Unknown factor:", "type")
	}

	if !AllowedRating("type", "crash") {
		t.Error("Not allowed rating for:", "type", " - value:", "crash")
	}

	if AllowedRating("type", "blocking") {
		t.Error("Not allowed rating for:", "type", " - value:", "blocking")
	}
}
