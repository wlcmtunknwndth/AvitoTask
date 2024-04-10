package postgres

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/wlcmtunknwndth/AvitoTask/internal/config"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage/postgres"
	"log/slog"
	"strconv"
	"testing"
)

var testCasesBanners []storage.Banner

func InitStorage() (*postgres.Storage, error) {
	cfg := config.DbConfig{
		DbUser:  "postgres",
		DbName:  "Avito",
		DbPass:  "liza",
		SslMode: "disable",
	}

	pgsql, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}
	return pgsql, err
}

func InitData() {
	testCasesBanners = []storage.Banner{
		{
			FeatureId: gofakeit.UintRange(1, 324342321),
			Tag:       gofakeit.UintRange(1, 23432342),
			Title:     gofakeit.Username(),
			Text:      gofakeit.BuzzWord(),
			Url:       gofakeit.URL(),
		},
		{
			FeatureId: gofakeit.UintRange(1, 324342321),
			Tag:       gofakeit.UintRange(1, 23432342),
			Title:     gofakeit.Username(),
			Text:      gofakeit.BuzzWord(),
			Url:       gofakeit.URL(),
		},
		{
			FeatureId: gofakeit.UintRange(1, 324342321),
			Tag:       gofakeit.UintRange(1, 23432342),
			Title:     gofakeit.Username(),
			Text:      gofakeit.BuzzWord(),
			Url:       gofakeit.URL(),
		},
		{
			FeatureId: gofakeit.UintRange(1, 324342321),
			Tag:       gofakeit.UintRange(1, 23432342),
			Title:     gofakeit.Username(),
			Text:      gofakeit.BuzzWord(),
			Url:       gofakeit.URL(),
		},
	}
}

func TestSaveBanner(t *testing.T) {
	pgsql, err := InitStorage()
	InitData()
	if err != nil {
		slog.Info("Error:", err)
		return
	}
	defer func(pgsql *postgres.Storage) {
		err := pgsql.Close()
		if err != nil {
			t.Errorf("coudln't close connection: %s", err)
			return
		}
	}(pgsql)
	for i, testCase := range testCasesBanners {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err = pgsql.SaveBanner(&testCase)
			if err != nil {
				t.Errorf("couldn't save banner: %s", err)
			}
		})
	}
}

func TestGetBanner(t *testing.T) {
	pgsql, err := InitStorage()
	if err != nil {
		slog.Info("Error:", err)
		return
	}
	defer func(pgsql *postgres.Storage) {
		err := pgsql.Close()
		if err != nil {
			t.Errorf("coudln't close connection: %s", err)
			return
		}
	}(pgsql)
	args := [][2]uint{
		{267881467, 9072019},
		{173749647, 11053972},
		{2286002, 1648162},
		{259341820, 6710789},
	}
	//args := [][2]uint{
	//	{testCasesBanners[0].FeatureId, testCasesBanners[0].Tag},
	//	{testCasesBanners[1].FeatureId, testCasesBanners[1].Tag},
	//	{testCasesBanners[2].FeatureId, testCasesBanners[2].Tag},
	//	{testCasesBanners[3].FeatureId, testCasesBanners[3].Tag},
	//}

	for i, arg := range args {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			banner, err := pgsql.GetBanner(arg[0], arg[1])
			if err != nil {
				t.Errorf("couldn't get banner: %s", err)
				return
			}
			t.Logf("banner: %+v", banner)
		})
	}
}

func TestDeleteBanner(t *testing.T) {
	pgsql, err := InitStorage()
	if err != nil {
		slog.Info("Error:", err)
		return
	}
	defer func(pgsql *postgres.Storage) {
		err := pgsql.Close()
		if err != nil {
			t.Errorf("coudln't close connection: %s", err)
			return
		}
	}(pgsql)
	for i := 8; i < 9; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := pgsql.DeleteBanner(uint(i))
			if err != nil {
				t.Errorf("couldn't delete banner")
				return
			}
		})
	}
}
