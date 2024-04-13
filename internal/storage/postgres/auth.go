package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/wlcmtunknwndth/AvitoTask/internal/auth"
)

const minLenUsername = 5
const minLenPassword = 5

func (s *Storage) GetPassword(username string) (string, error) {
	const op = "storage.postgres.GetPassword"
	if len(username) < minLenUsername {
		return "", fmt.Errorf("%s: username is too short", op)
	}
	stmt, err := s.db.Prepare(getPassword)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	var pass string
	err = stmt.QueryRow(username).Scan(&pass)
	if errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return pass, err
}

func (s *Storage) RegisterUser(usr *auth.User) error {
	const op = "storage.postgres.RegisterUser"
	if len(usr.Username) < minLenUsername {
		return fmt.Errorf("%s: Invalid login", usr.Username)
	}
	if len(usr.Password) < minLenPassword {
		return fmt.Errorf("%s: Invalid password", usr.Password)
	}

	_, err := s.db.Exec(registerUser, usr.Username, usr.Password)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) DeleteUser(username string) error {
	const op = "storage.postgres.DeleteUser"

	stmt, err := s.db.Prepare(deleteUser)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(username)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) IsAdmin(username string) bool {
	if len(username) < minLenUsername {
		return false
	}
	var ans bool
	stmt, err := s.db.Prepare(isAdmin)
	if err != nil {
		return false
	}

	err = stmt.QueryRow(username).Scan(&ans)
	if err != nil {
		return false
	}
	return ans
}
