package poker_test

import (
	"bytes"
	"io"
	poker "learn-go-with-tests"
	"strings"
	"testing"
)

var dummyStdOut = &bytes.Buffer{}
var dummySpyAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}

type GameSpy struct {
	StartCalled     bool
	StartCalledWith int

	FinishCalled     bool
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartCalledWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalled = true
	g.FinishCalledWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("3", "Chris wins")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Cleo wins")

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("Pie")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("3", "Lloyd is a killer")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertNotFinished(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
	})
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	got := stdout.String()
	want := strings.Join(messages, "")
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, numberOfPlayersWanted int) {
	t.Helper()
	if game.StartCalledWith != numberOfPlayersWanted {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayersWanted, game.StartCalledWith)
	}
}

func assertFinishCalledWith(t testing.TB, game *GameSpy, winner string) {
	t.Helper()
	if game.FinishCalledWith != winner {
		t.Errorf("wanted Finish called with %s but got %s", winner, game.FinishCalledWith)
	}
}

func assertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Error("game should not have started")
	}
}

func assertNotFinished(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.FinishCalled {
		t.Error("game should not have finished")
	}
}
