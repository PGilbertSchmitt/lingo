package phonology

import (
	"fmt"
	"math/rand"
)

// Phone covers both vowels and consonants
type Phone interface {
	Char() rune
}

// Matchable should cover a Consonant, Vowel, or Set of phones
type Matchable interface {
	Match(Phone) bool
}

// WeightedPhones wraps any held phones with their respective weights
type WeightedPhones struct {
	phones      []Phone
	weights     []int
	totalWeight int
}

// AddPhone takes a single phone and weight and adds it to the WeightedPhones struct
func (w *WeightedPhones) AddPhone(phone Phone, weight int) {
	w.phones = append(w.phones, phone)
	w.weights = append(w.weights, weight)
	w.totalWeight += weight
}

// RandomPhone uses the weights to pull an appropriate phone from the WeightedPhones struct
func (w WeightedPhones) RandomPhone() Phone {
	var idx, rollingSum int
	randVal := int(rand.Float64() * float64(w.totalWeight))
	for idx = 0; idx < len(w.weights); idx++ {
		rollingSum += w.weights[idx]
		if randVal < rollingSum {
			return w.phones[idx]
		}
	}

	fmt.Println("Should never get here")
	return w.phones[idx-1]
}
