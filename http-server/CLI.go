package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
const BadWinnerInputMsg = "Bad value received for winner input, please try again with the pattern '{name} wins'"

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

// func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers, cli.out)

	winnerInput := cli.readLine()
	winner, err := extractWinner(winnerInput)

	if err != nil {
		fmt.Fprint(cli.out, BadWinnerInputMsg)
		return
	}

	cli.game.Finish(winner)
}

const playerWinsSuffix = " wins"
const playerWinsExample = "{name}" + playerWinsSuffix
const ErrBadWinnerSuffix = "bad winner pattern, try again with " + playerWinsExample + "pattern"

func hasWinsSuffix(userInput string) bool {
	return strings.HasSuffix(userInput, playerWinsSuffix)
}

func extractWinner(userInput string) (string, error) {
	if hasWinsSuffix(userInput) {
		return strings.Replace(userInput, playerWinsSuffix, "", 1), nil
	}

	return "", errors.New(ErrBadWinnerSuffix)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
