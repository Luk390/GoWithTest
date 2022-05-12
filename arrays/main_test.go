package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1,2,-1}
		got := Sum(numbers)
		want := 2
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	slice1 := []int{1,2,3}
	slice2 := []int{4,5,6}
	slice3 := []int{4,5,6}
	got := SumAll(slice1, slice2, slice3)
	want := []int{6, 15, 15}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("sum populated slices", func(t0* testing.T) {
		slice1 := []int{1,2,3}
		slice2 := []int{4,5,6}
		slice3 := []int{4,5,6}
		got := SumAllTails(slice1, slice2, slice3)
		want := []int{5, 11, 11}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{}
		slice3 := []int{1,2,3}
		got := SumAllTails(slice1, slice2, slice3)
		want := []int{0, 0, 5}
		checkSums(t, got, want)

	})
}