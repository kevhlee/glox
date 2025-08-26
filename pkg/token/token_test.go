package token

import "testing"

// TestTypeHasName checks to make sure each lexical type has a
// non-empty, unique name.
func Test(t *testing.T) {
	seen := make(map[string]Type)

	for curr := __START__ + 1; curr < __END__; curr++ {
		name := curr.String()

		if len(name) == 0 {
			t.Errorf("Lexical type %d has no name", curr)
			continue
		}

		if prev, ok := seen[name]; ok {
			t.Errorf("Lexical type %d has the same name as type %d", curr, prev)
		}
		seen[name] = curr
	}
}
