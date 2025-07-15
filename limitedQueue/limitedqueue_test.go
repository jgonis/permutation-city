package limitedqueue

import (
	"slices"
	"testing"
)

func TestLimitedQueue(t *testing.T) {
	testCases := map[string]struct {
		setupFunc func() LimitedQueue[int]
		expected  []int
	}{
		"insert lower end doesn't add": {
			setupFunc: func() LimitedQueue[int] {
				lq := CreateLimitedQueue[int](3, func(a, b int) int {
					return a - b
				})
				lq.Insert(2)
				lq.Insert(3)
				lq.Insert(4)
				lq.Insert(5)
				return lq
			},
			expected: []int{2, 3, 4},
		},
		"insert higher end does add": {
			setupFunc: func() LimitedQueue[int] {
				lq := CreateLimitedQueue[int](3, func(a, b int) int {
					return a - b
				})
				lq.Insert(3)
				lq.Insert(4)
				lq.Insert(5)
				lq.Insert(2)
				lq.Insert(1)
				return lq
			},
			expected: []int{1, 2, 3},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			lq := testCase.setupFunc()
			if !slices.Equal(lq.items, testCase.expected) {
				t.Errorf("test failed, expected: %v, got: %v", testCase.expected, lq.items)
			}
		})
	}
}
