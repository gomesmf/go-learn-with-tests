package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	poker "github.com/gomesmf/go-learn-with-tests/http-server"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules Alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(5)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		poker.CheckSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		poker.CheckSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}

		game := poker.NewTexasHoldem(dummySpyAlerter, playerStore)
		game.Finish("Cleo")

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}

var dummyBlindAlerter = &poker.SpyBlindAlerter{}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := poker.NewPlayerServer(&poker.StubPlayerStore{})

		request := NewGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
	})
}

func NewGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return req
}

func assertStatus(t testing.TB, response *httptest.ResponseRecorder, want int) {
	t.Helper()

	if response.Code != want {
		t.Errorf("did not get correct status, got %d, want %d", response.Code, want)
	}
}
