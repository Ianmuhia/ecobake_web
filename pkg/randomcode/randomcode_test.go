package randomcode_test

import (
	"fainda/pkg/randomcode"
	"testing"
)

func FuzzBusinessCode(f *testing.F) {
	tests := make([]int, 10)
	tests = append(tests, 10, 20, 240, 40, 50, 7, 6, 2345, 45)
	for _, tc := range tests {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, n int) {
		got := randomcode.Code(n)
		t.Log(len(got))
		t.Logf("ExpiryCode() %v", got)

	})

}

//func Sum(vals []int64) int64 {
//	var total int64  for _, val := range vals {
//		if val%1e5 ! = 0 {
//			total += val
//		}
//	}  return total
//}
//func FuzzSum(f *testing.F) {
//	f.Add(10)
//	f.Fuzz(func(t *testing.T, n int) {
//		n %= 20
//		...
//	})
//}

//func FuzzSum(f *testing.F) {
//	rand.Seed(time.Now().UnixNano())  f.Add(10)
//	f.Fuzz(func(t *testing.T, n int) {
//		n %= 20
//		var vals []int64
//		var expect int64
//		for i := 0; i < n; i++ {
//			val := rand.Int63() % 1e6
//			vals = append(vals, val)
//			expect += val
//		}    assert.Equal(t, expect, Sum(vals))
//	})
//}

//func FuzzReverse(f *testing.F) {
//	testcases := []string{"Hello, world", " ", "!12345"}
//	for _, tc := range testcases {
//		f.Add(tc) // Use f.Add to provide a seed corpus
//	}
//	f.Fuzz(func(t *testing.T, orig string) {
//		rev, err1 := Reverse(orig)
//		if err1 != nil {
//			return
//		}
//		doubleRev, err2 := Reverse(rev)
//		if err2 != nil {
//			return
//		}
//		if orig != doubleRev {
//			t.Errorf("Before: %q, after: %q", orig, doubleRev)
//		}
//		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
//			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
//		}
//	})
//}
