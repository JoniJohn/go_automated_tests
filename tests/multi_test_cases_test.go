package tests

import (
	discounter "automated_tests/discounter"
	"testing"
)

func TestDiscountedPrice(t *testing.T) {
	// arrange
	testCases := []struct {
		price           float64
		discountPercent float64
		expected        float64
		expectedError   bool
		desc            string
	}{
		{100.0, 0.0, 100.0, false, "no discount"},
		{100.0, 50.0, 50.0, false, "50% discount"},
		{100.0, 100.0, 0.0, false, "100% discount"},
		{100.0, 110.0, 0.0, true, "discount greater than 100%"},
		{100.0, -10.0, 0.0, true, "negative discount"},
	}

	// act + assert
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := discounter.DiscountedPrice(tc.price, tc.discountPercent)
			if tc.expectedError && err == nil {
				t.Errorf(
					"DiscountedPrice(%.2f, %.2f) "+"should return an error",
					tc.price,
					tc.discountPercent,
				)
			}

			if !tc.expectedError && err != nil {
				t.Errorf(
					"DiscountedPrice(%.2f, %.2f) "+"returned an error: %v",
					tc.price,
					tc.discountPercent,
					err,
				)
			}

			if !tc.expectedError && res != tc.expected {
				t.Errorf(
					"DiscountedPrice(%.2f, %.2f) = "+"%.2f; want %.2f",
					tc.price, tc.discountPercent,
					res, tc.expected,
				)
			}
		})
	}
}
