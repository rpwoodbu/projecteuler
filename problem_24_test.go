package projecteuler

import "testing"

func TestPermutations(t *testing.T) {
	permutationsOf3 := [][]int{
		[]int{0, 1, 2},
		[]int{0, 2, 1},
		[]int{1, 0, 2},
		[]int{1, 2, 0},
		[]int{2, 0, 1},
		[]int{2, 1, 0},
	}

	testPermutations := Permutations(3)
	for _, p := range permutationsOf3 {
		tp, ok := <-testPermutations
		if !ok {
			t.Error("Ran out of permutations.")
			break
		}
		if len(p) != len(tp) {
			t.Errorf("Expected %v, got %v", p, tp)
		} else {
			for i := range p {
				if p[i] != tp[i] {
					t.Errorf("Expected %v, got %v", p, tp)
					break
				}
			}
		}
	}

	_, ok := <-testPermutations
	if ok {
		t.Error("More permutations than expected.")
	}
}

func TestProblem24(t *testing.T) {
	t.Parallel()
	expected := "2783915460"
	value, ok := Problem24()
	if !ok {
		t.Error("Did not get an answer.")
	} else if value != expected {
		t.Errorf("Expected %v, got %v", expected, value)
	}
}
