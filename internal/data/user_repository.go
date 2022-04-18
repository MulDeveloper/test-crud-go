package data

import (
	"context"
	"time"

	"github.com/MulDeveloper/go-test-crud/pkg/user"

	"github.com/google/uuid"
)

type UserRepository struct {
	Data *Data
}

func (ur *UserRepository) GetAll(ctx context.Context) ([]user.User, error) {
	q := `
	SELECT id, first_name, last_name, username, email,
		created_at, updated_at
		FROM users;
	`

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username,
			&u.Email, &u.CreatedAt, &u.UpdatedAt)
		users = append(users, u)
	}

	return users, nil
}

func (ur *UserRepository) GetOne(ctx context.Context, id uuid.UUID) (user.User, error) {
	q := `
    SELECT id, first_name, last_name, username, email,
        created_at, updated_at
        FROM users WHERE id = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var u user.User
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username, &u.Email,
		&u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (ur *UserRepository) GetByUsername(ctx context.Context, username string) (user.User, error) {
	q := `
    SELECT id, first_name, last_name, username, email,
        password, created_at, updated_at
        FROM users WHERE username = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, username)

	var u user.User
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username,
		&u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (ur *UserRepository) Create(ctx context.Context, u *user.User) error {
	q := `
    INSERT INTO users (id, first_name, last_name, username, email, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id;
    `

	if err := u.HashPassword(); err != nil {
		return err
	}

	row := ur.Data.DB.QueryRowContext(
		ctx, q, u.ID, u.FirstName, u.LastName, u.Username, u.Email, u.PasswordHash, time.Now(), time.Now(),
	)

	err := row.Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Update(ctx context.Context, id uuid.UUID, u user.User) error {
	q := `
    UPDATE users set first_name=$1, last_name=$2, email=$3, updated_at=$4
        WHERE id=$5;
    `

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, u.FirstName, u.LastName, u.Email, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	q := `DELETE FROM users WHERE id=$1;`

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
