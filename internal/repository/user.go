package repository

import (
	"context"
	"database/sql"
	"fmt"
	"frappuccino/internal/helpers"
	"frappuccino/internal/models"
	"frappuccino/pkg/jtoken"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserRepository struct {
	Db  *sql.DB
	Rdb *redis.Client
}

const (
	expirationJWT = time.Hour * 5
)

func NewUserRepository(db *sql.DB, rdb *redis.Client) *UserRepository {
	return &UserRepository{
		Db:  db,
		Rdb: rdb,
	}
}

func (r *UserRepository) Register(ctx context.Context) {}
func (r *UserRepository) GetToken(ctx context.Context, username, pass string) (string, error) {
	query := `SELECT * FROM users WHERE username = $1`

	user := &models.User{}
	err := r.Db.QueryRowContext(ctx, query, username).Scan(
		&user.UserID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Age,
		&user.Sex,
		&user.FirstOrder,
		&user.Allergens,
	)
	if err != nil {
		return "", fmt.Errorf("error fetching user: %v", err)
	}

	hashedPass := helpers.CreateMd5Hash(pass)
	if hashedPass != user.Password {
		return "", fmt.Errorf("invalid password")
	}

	payload := make(map[string]interface{})
	payload["role"] = user.Role
	payload["username"] = user.Username
	payload["age"] = user.Age
	payload["sex"] = user.Sex
	payload["first_order"] = user.FirstOrder
	payload["allergens"] = user.Allergens
	payload["expires_at"] = time.Now().Add(expirationJWT).Unix()

	token, err := jtoken.GenerateAccessToken(ctx, r.Rdb, payload)
	if err != nil {
		return "", fmt.Errorf("eror generating jwt token: %v", err)
	}

	return token, nil
}
