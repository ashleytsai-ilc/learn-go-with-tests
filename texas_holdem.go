package poker

import (
	"io"
	"time"
)

type TexasHoldem struct {
	store   PlayerStore
	alerter BlindAlerter
}

func NewTexasHoldem(store PlayerStore, alerter BlindAlerter) Game {
	return &TexasHoldem{
		store:   store,
		alerter: alerter,
	}
}

func (t *TexasHoldem) Start(numberOfPlayers int, alertsDestination io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		t.alerter.ScheduledAlertAt(blindTime, blind, alertsDestination)
		blindTime += blindIncrement
	}
}

func (t *TexasHoldem) Finish(winner string) {
	t.store.RecordWin(winner)
}
