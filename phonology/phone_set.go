package phonology

// A PSet is a set type that stores a rune to represent a phone
type PSet map[rune]bool

// And takes several sets and
func And(sets ...PSet) PSet {
	if len(sets) == 0 {
		return nil
	}

	out := make(PSet)

	first, rest := sets[0], sets[1:]

	for char := range first {
		inAll := true

		for _, set := range rest {
			if !(set[char]) {
				inAll = false
				break
			}
		}

		if inAll {
			out[char] = true
		}
	}

	return out
}

// Match checks if the given phone is included in the set. Even though the function itself is trivial for a set type, it allows the PSet to be stored in a Matchable array alongside Vowels and Consonants.
func (s PSet) Match(p Phone) bool {
	return s[p.Char()]
}
