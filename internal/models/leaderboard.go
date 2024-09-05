package models

type LeaderboardEntry struct {
	UserID string  `json:"user_id"`
	Score  float64 `json:"score"`
}

type Leaderboard struct {
	Name string `json:"name"`
}
