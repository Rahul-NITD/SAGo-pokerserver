package pokerhttpserver

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		store: map[string]int{},
	}
}

type InMemoryStorage struct {
	store map[string]int
}

func (ts *InMemoryStorage) GetPlayerScore(player string) (int, error) {
	return ts.store[player], nil
}

func (ts *InMemoryStorage) RecordWin(player string) error {
	ts.store[player]++
	return nil
}
