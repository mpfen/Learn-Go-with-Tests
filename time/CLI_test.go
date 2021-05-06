package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	poker "github.com/mpfen/Learn-Go-with-Tests/time"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}

type GameSpy struct {
	StartCalledWith  int
	FinishCalledWith string
	StartCalled      bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalledWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalledWith = winner
}

func TestCLI(t *testing.T) {

	t.Run("Start a game with 3 players and finish game with 'Chris' as winer", func(t *testing.T) {
		game := &GameSpy{}
		out := &bytes.Buffer{}
		in := userSends("3", "Chris")
		cli := poker.NewCLI(in, out, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, out, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}
		out := &bytes.Buffer{}
		in := userSends("8", "Cleo")
		cli := poker.NewCLI(in, out, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, out, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func assertScheduledAlert(t *testing.T, want, got poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got amount %v, want %v", got, want)
	}
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q, expected %q", got, want)
	}
}

func assertGameStartedWith(t *testing.T, game *GameSpy, numberOfPlayersWanted int) {
	t.Helper()

	if game.StartCalledWith != numberOfPlayersWanted {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayersWanted, game.StartCalledWith)
	}
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()

	if game.FinishCalledWith != winner {
		t.Errorf("wanted Finish calles with %q but got %q", winner, game.FinishCalledWith)
	}
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game should not have been started")
	}
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}
