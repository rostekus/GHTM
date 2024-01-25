package repository

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/rostekus/ghtm/internal/app/domain"
)

// InMemoryUserRepository is an in-memory implementation of the UserRepository.
type InMemoryUserRepository struct {
	users map[string]domain.User
	mu    sync.RWMutex
}

// NewInMemoryUserRepository creates a new instance of InMemoryUserRepository.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]domain.User),
	}
}

// CreateUser adds a new user to the repository.
func (r *InMemoryUserRepository) CreateUser(user domain.User) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.Email]; exists {
		return domain.User{}, errors.New("user already exists")
	}
	user.Id, _ = uuid.NewUUID()
	r.users[user.Email] = user
	return user, nil
}

// GetUserByID retrieves a user from the repository by ID.
func (r *InMemoryUserRepository) GetUser(user domain.User) (domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[user.Email]
	if !exists {
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}
