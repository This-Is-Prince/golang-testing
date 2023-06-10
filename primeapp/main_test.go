package main

import "testing"

// func Test_isPrime(t *testing.T) {
// 	n := 0
// 	result, msg := isPrime(n)

// 	if result {
// 		t.Errorf("with %d as test parameter, got true, but expected false", n)
// 	}

// 	if msg != "0 is not prime, by definition!" {
// 		t.Error("wrong message returned:", msg)
// 	}

// 	n = 7
// 	result, msg = isPrime(n)

// 	if !result {
// 		t.Errorf("with %d as test parameter, got false, but expected true", n)
// 	}

// 	if msg != "7 is a prime number!" {
// 		t.Error("wrong message returned:", msg)
// 	}
// }

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"7Prime", 7, true, "7 is a prime number!"},
		{"1Prime", 1, false, "1 is not prime, by definition!"},
		{"-1Prime", -1, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
