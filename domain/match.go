package domain

type Match struct {
	GameID       string              `json:"game_id"`
	TotalKills   int                 `json:"total_kills"`
	Players      map[string]struct{} `json:"-"`
	PlayerList   []string            `json:"players"`
	Kills        map[string]int      `json:"kills"`
	KillsByMeans map[string]int      `json:"kills_by_means"`
}

func (m *Match) GeneratePlayerList() {
	m.PlayerList = []string{}
	for player := range m.Players {
		m.PlayerList = append(m.PlayerList, player)
	}
}
