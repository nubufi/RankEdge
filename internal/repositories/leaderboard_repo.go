package repositories

import (
	"context"

	"RankEdge/internal/models"
	"RankEdge/internal/utils"

	"github.com/redis/go-redis/v9"
)

type LeaderboardRepository interface {
	CreateLeaderboard(boardName string) error
	GetLeaderboards() ([]string, error)
	GetLeaderboard(boardName string) ([]models.LeaderboardEntry, error)
	UpdateLeaderboard(boardName string, entry models.LeaderboardEntry) error
	GetTopNUsers(boardName string, n int) ([]models.LeaderboardEntry, error)
	GetUserRankAndScore(boardName, userID string) (int, float64, error)
	RemoveUser(boardName, userID string) error
}

type leaderboardRepository struct {
	redisClient *redis.Client
}

// NewLeaderboardRepository creates a new instance of LeaderboardRepository
func NewLeaderboardRepository(redisClient *redis.Client) LeaderboardRepository {
	return &leaderboardRepository{redisClient}
}

// CreateLeaderboard creates a new leaderboard in the cache
func (r *leaderboardRepository) CreateLeaderboard(boardName string) error {
	// Create a new leaderboard in the cache
	_, err := r.redisClient.ZAdd(context.Background(), boardName, redis.Z{}).Result()
	if err != nil {
		return err
	}

	return nil
}

// GetLeaderboards retrieves the names of all leaderboards in the cache
func (r *leaderboardRepository) GetLeaderboards() ([]string, error) {
	// Get the names of all leaderboards in the cache
	val, err := r.redisClient.Keys(context.Background(), "*").Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

// GetLeaderboard retrieves the current leaderboard from the cache
func (r *leaderboardRepository) GetLeaderboard(boardName string) ([]models.LeaderboardEntry, error) {
	// Get the leaderboard from the cache
	val, err := r.redisClient.ZRevRangeWithScores(context.Background(), boardName, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	leaderboard := utils.ConvertRedisZToLeaderboardEntry(val)

	return leaderboard, nil
}

// UpdateLeaderboard updates the leaderboard with a new entry
func (r *leaderboardRepository) UpdateLeaderboard(boardName string, entry models.LeaderboardEntry) error {
	// Add the new entry to the leaderboard
	_, err := r.redisClient.ZAdd(context.Background(), boardName, redis.Z{
		Score:  entry.Score,
		Member: entry.UserID,
	}).Result()
	if err != nil {
		return err
	}

	return nil
}

// GetTopNUsers retrieves the top N users from the leaderboard
func (r *leaderboardRepository) GetTopNUsers(boardName string, n int) ([]models.LeaderboardEntry, error) {
	// Get the top N users from the leaderboard
	val, err := r.redisClient.ZRevRangeWithScores(context.Background(), boardName, 0, int64(n-1)).Result()
	if err != nil {
		return nil, err
	}

	leaderboard := utils.ConvertRedisZToLeaderboardEntry(val)

	return leaderboard, nil
}

// GetUserRankAndScore retrieves the rank and score of a user from the leaderboard
func (r *leaderboardRepository) GetUserRankAndScore(boardName, userID string) (int, float64, error) {
	// Get the rank and score of the user
	rank, err := r.redisClient.ZRevRank(context.Background(), boardName, userID).Result()
	if err != nil {
		return 0, 0, err
	}

	score, err := r.redisClient.ZScore(context.Background(), boardName, userID).Result()
	if err != nil {
		return 0, 0, err
	}

	return int(rank) + 1, score, nil
}

// RemoveUser removes a user from the leaderboard
func (r *leaderboardRepository) RemoveUser(boardName, userID string) error {
	// Remove the user from the leaderboard
	_, err := r.redisClient.ZRem(context.Background(), boardName, userID).Result()
	if err != nil {
		return err
	}

	return nil
}
