package poker_test

import (
	"bytes"

	poker "github.com/gomesmf/go-learning/http-server"

	"strings"
	"testing"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		poker.AssertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		poker.AssertNumberOfPlayers(t, game.StartedWith, 7)
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		poker.AssertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
		poker.AssertStartNotCalled(t, game.StartCalled)
	})
}
