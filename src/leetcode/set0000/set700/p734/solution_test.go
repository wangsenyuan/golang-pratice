package p734

import "testing"

func TestSample1(t *testing.T) {
	words1 := []string{"great", "acting", "skills"}
	words2 := []string{"fine", "drama", "talent"}
	pairs := [][]string{
		{"great", "fine"}, {"acting", "drama"}, {"skills", "talent"},
	}

	res := areSentencesSimilar(words1, words2, pairs)
	if !res {
		t.Errorf("the sample %v and %v with give pairs %v, are simliar, but got not simliar", words1, words2, pairs)
	}
}
