package services

import (
	"RankEdge/internal/models"
	r "RankEdge/internal/repositories"
)

type LeaderboardService interface {
	CreateLeaderboard(boardName string) error
	GetLeaderboards() ([]string, error)
	GetLeaderboard(boardName string) ([]models.LeaderboardEntry, error)
	UpdateLeaderboard(boardName string, entry models.LeaderboardEntry) error
	GetTopNUsers(boardName string, n int) ([]models.LeaderboardEntry, error)
	GetUserRankAndScore(boardName, userID string) (int, float64, error)
	RemoveUser(boardName, userID string) error
}

type leaderboardService struct {
	leaderboardRepo r.LeaderboardRepository
}

// NewLeaderboardService creates a new instance of LeaderboardService
func NewLeaderboardService(leaderboardRepo r.LeaderboardRepository) LeaderboardService {
	return &leaderboardService{leaderboardRepo}
}

// CreateLeaderboard creates a new leaderboard in the cache
func (s *leaderboardService) CreateLeaderboard(boardName string) error {
	return s.leaderboardRepo.CreateLeaderboard(boardName)
}

// GetLeaderboards retrieves the names of all leaderboards in the cache
func (s *leaderboardService) GetLeaderboards() ([]string, error) {
	return s.leaderboardRepo.GetLeaderboards()
}

// GetLeaderboard retrieves the current leaderboard from the cache
func (s *leaderboardService) GetLeaderboard(boardName string) ([]models.LeaderboardEntry, error) {
	return s.leaderboardRepo.GetLeaderboard(boardName)
}

// UpdateLeaderboard updates the leaderboard in the cache
func (s *leaderboardService) UpdateLeaderboard(boardName string, entry models.LeaderboardEntry) error {
	return s.leaderboardRepo.UpdateLeaderboard(boardName, entry)
}

// GetTopNUsers retrieves the top N users from the leaderboard
func (s *leaderboardService) GetTopNUsers(boardName string, n int) ([]models.LeaderboardEntry, error) {
	return s.leaderboardRepo.GetTopNUsers(boardName, n)
}

// GetUserRankAndScore retrieves the rank and score of a user from the leaderboard
func (s *leaderboardService) GetUserRankAndScore(boardName, userID string) (int, float64, error) {
	return s.leaderboardRepo.GetUserRankAndScore(boardName, userID)
}

// RemoveUser removes a user from the leaderboard
func (s *leaderboardService) RemoveUser(boardName, userID string) error {
	return s.leaderboardRepo.RemoveUser(boardName, userID)
}
