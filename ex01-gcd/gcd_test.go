package gcd

import "testing"

var testCases = []struct {
    input    [2]int
    expected int
}{
    {
        input:    [2]int{1, 3},
        expected: 1,
    },
    {
        input:    [2]int{1, 1},
        expected: 1,
    },
    {
        input:    [2]int{4, 4},
        expected: 4,
    },
    {
        input:    [2]int{100, 10},
        expected: 10,
    },
    {
        input:    [2]int{45, 48},
        expected: 3,
    },
    {
        input:    [2]int{7*6*5*4*3, 7*5*3},
        expected: 7*5*3,
    },
    {
        input:    [2]int{121*121, 10*11*12},
        expected: 11,
    },
}

func TestGcd(t *testing.T) {
    for _, tt := range testCases {
        actual, err := Gcd(tt.input[0], tt.input[1])
        // We don't expect errors for any of the test cases
        if err != nil {
            t.Fatalf("Downcase(%q) returned error %q.  Error not expected.", tt.input, err)
        }
        if actual != tt.expected {
            t.Fatalf("Downcase(%q) was expected to return %v but returned %v.",
                tt.input, tt.expected, actual)
        }
    }
}

func BenchmarkGcd(b *testing.B) {
    b.StopTimer()
    for _, tt := range testCases {
        b.StartTimer()
        for i := 0; i < b.N; i++ {
            Gcd(tt.input[0], tt.input[1])
        }
        b.StopTimer()
    }
}

