package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(rdr io.Reader) (League, error) {
	var player League

	err := json.NewDecoder(rdr).Decode(&player)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return player, err
}

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
