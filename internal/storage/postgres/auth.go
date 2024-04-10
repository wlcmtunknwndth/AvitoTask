package postgres

import (
	"database/sql"
	"errors"
	"fmt"
)

func (s *Storage) GetPassword(login string) (string, error) {
	const op = "storage.postgres.GetPassword"
	if len(login) < minLenLogin {
		return "", fmt.Errorf("%s: login is too short", op)
	}
	stmt, err := s.db.Prepare(getPassword)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	var t string
	var pass string
	err = stmt.QueryRow(login).Scan(&t, &pass)
	if errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return pass, err
}

func (s *Storage) RegisterUser(login, pass string) error {
	const op = "storage.postgres.RegisterUser"
	if len(login) < minLenLogin {
		return fmt.Errorf("%s: Invalid login", login)
	}
	if len(pass) < minLenPass {
		return fmt.Errorf("%s: Invalid password", login)
	}

	_, err := s.db.Exec(registerUser, login, pass)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) DeleteUser(login string) error {
	const op = "storage.postgres.DeleteUser"

	stmt, err := s.db.Prepare(deleteUser)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(login)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
