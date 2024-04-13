package postgres

import (
	"fmt"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage"
	"strconv"
	"strings"
)

func (s *Storage) RestoreCache() ([]storage.Banner, error) {
	const op = "storage.postgres.RestoreCache"
	rows, err := s.db.Query(restoreCache)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var banners []storage.Banner
	for rows.Next() {
		var banner storage.Banner

		err = rows.Scan(&banner.Id, &banner.FeatureId, &banner.Tag, &banner.Title, &banner.Text, &banner.Url)
		if err != nil {
			continue
		}
	}
	return banners, nil
}

func (s *Storage) SaveCache(uuid string) error {
	const op = "storage.postgres.SaveCache"
	stmt, err := s.db.Prepare(saveCache)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	featureId, tagId := SplitUuid(uuid)
	_, err = stmt.Exec(featureId, tagId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) DeleteCache(uuid string) error {
	const op = "storage.postgres.DeleteCache"
	stmt, err := s.db.Prepare(deleteCache)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	featureId, tagId := SplitUuid(uuid)
	_, err = stmt.Exec(featureId, tagId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) IsAlreadyCached(uuid string) bool {
	featureId, tagId := SplitUuid(uuid)
	if row := s.db.QueryRow(isAlreadyCached, featureId, tagId); row != nil {
		return false
	}
	return true
}

func SplitUuid(uuid string) (int, int) {
	args := strings.Split(uuid, "/")
	featureId, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, 0
	}

	tagId, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0
	}
	return featureId, tagId
}
