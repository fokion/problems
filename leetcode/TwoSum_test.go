package leetcode

import "testing"

func TestTwoSumWithFistLast(t *testing.T) {
	got := TwoSum([]int{0, 4, 3, 0}, 0)
	want := []int{0, 3}

	for i := 0; i < len(want); i++ {
		if !contains(got, want[i]) {
			t.Errorf("missing %d ", want[i])
		}
	}
}

func TestTwoSumWithFistSecond(t *testing.T) {
	got := TwoSum([]int{3, 2, 5, 0}, 5)
	want := []int{0, 1}

	for i := 0; i < len(want); i++ {
		if !contains(got, want[i]) {
			t.Errorf("missing %d ", want[i])
		}
	}
}

func TestTwoSumWithNegatives(t *testing.T) {
	got := TwoSum([]int{-3, -2, -5, 20}, -5)
	want := []int{0, 1}

	for i := 0; i < len(want); i++ {
		if !contains(got, want[i]) {
			t.Errorf("missing %d ", want[i])
		}
	}
}

func contains(elems []int, key int) bool {
	for _, s := range elems {
		if key == s {
			return true
		}
	}
	return false
}
