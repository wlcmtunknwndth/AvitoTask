package postgres

import (
	"github.com/stretchr/testify/require"
	"github.com/wlcmtunknwndth/AvitoTask/internal/config"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage/postgres"
	"testing"
)

func TestPostgresInit(t *testing.T) {
	cfg := config.DbConfig{
		DbUser:  "postgres",
		DbName:  "Avito",
		DbPass:  "liza",
		SslMode: "disable",
	}
	storage, err := postgres.New(cfg)
	if err != nil {
		t.Errorf("couldn't init postgres storage: %s", err)
	}
	defer func(storage *postgres.Storage) {
		err := storage.Close()
		if err != nil {
			t.Errorf("couldn't close DB: %s", err)
		}
	}(storage)
}

var testCases = []struct {
	TestName string
	Login    string
	Pass     string
}{
	{
		TestName: "Valid claims",
		Login:    "1233242",
		Pass:     "12121323",
	},
	{
		TestName: "Invalid claims",
		Login:    "",
		Pass:     "",
	},
}

func TestRegisterUser(t *testing.T) {
	cfg := config.DbConfig{
		DbUser:  "postgres",
		DbName:  "Avito",
		DbPass:  "liza",
		SslMode: "disable",
	}
	storage, err := postgres.New(cfg)
	if err != nil {
		t.Errorf("couldn't init postgres storage: %s", err)
		return
	}
	defer func(storage *postgres.Storage) {
		err := storage.Close()
		if err != nil {
			t.Errorf("couldn't close DB: %s", err)
		}
	}(storage)

	for _, tCase := range testCases {
		t.Run(tCase.TestName, func(t *testing.T) {
			err := storage.RegisterUser(tCase.Login, tCase.Pass)
			if err != nil {
				t.Errorf("couldn't register user: %s", err)
			}
		})
	}
}

func TestGetPassword(t *testing.T) {
	cfg := config.DbConfig{
		DbUser:  "postgres",
		DbName:  "Avito",
		DbPass:  "liza",
		SslMode: "disable",
	}
	storage, err := postgres.New(cfg)
	if err != nil {
		t.Errorf("couldn't init postgres storage: %s", err)
		return
	}
	defer func(storage *postgres.Storage) {
		err := storage.Close()
		if err != nil {
			t.Errorf("couldn't close DB: %s", err)
		}
	}(storage)

	for _, tCase := range testCases {
		t.Run(tCase.TestName, func(t *testing.T) {
			pass, err := storage.GetPassword(tCase.Login)
			if err != nil {
				t.Errorf("couldn't get user's password: %s", err)
				return
			}
			t.Logf("For login %s the password is %s and got: %s", tCase.Login, tCase.Pass, pass)
			require.Equal(t, tCase.Pass, pass)
		})
	}
}

func TestDeleteUser(t *testing.T) {

}
