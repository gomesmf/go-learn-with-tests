package poker_test

import (
	"bytes"

	poker "github.com/gomesmf/go-learning/http-server"

	"testing"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}

		in := poker.UserSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		poker.AssertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		poker.AssertNumberOfPlayers(t, game.StartedWith, 3)
		poker.AssertFinishCalledWith(t, game, "Chris")
	})

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := poker.UserSends("7", "Bob wins")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		poker.AssertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		poker.AssertNumberOfPlayers(t, game.StartedWith, 7)
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := poker.UserSends("Pies")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		poker.AssertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
		poker.AssertStartNotCalled(t, game.StartCalled)
	})

	t.Run("it prints an error when the pattern '{name} wins' is not entered after game started", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := poker.UserSends("7", "Lloyd is a killer")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		poker.AssertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputMsg)

		if game.FinishCalled {
			t.Errorf("game should not have finished")
		}
	})
}
