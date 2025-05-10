package repository

import "go.mongodb.org/mongo-driver/mongo"

type (
	// declare UserRepository interface
	UserRepository interface {
		// DefaultRepository
	}

	// declare userRepository struct
	userRepository struct {
		db *mongo.Database
	}
)

// declare NewUserRepository function
func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db: db}
}

