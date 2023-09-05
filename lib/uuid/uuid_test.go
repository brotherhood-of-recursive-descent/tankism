package uuid_test

import (
	"testing"

	"github.com/co0p/tankism/lib/uuid"
)

func Test_New_should_generate_random_uuids(t *testing.T) {

	a := uuid.New()
	b := uuid.New()

	if a == b {
		t.Errorf("expected '%s' not to equal '%s'\n", a, b)
	}

}
