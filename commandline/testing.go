package poker

import "testing"

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.Scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, want string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatal("Expected a win call but didn't get any")
	}

	got := store.WinCalls[0]

	if got != want {
		t.Errorf("didin'T record correct winner, got %q, want %q", got, want)
	}

}
