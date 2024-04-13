package postgres

import (
	"fmt"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage"
)

func (s *Storage) SaveBanner(banner *storage.Banner) error {
	const op = "storage.postgres.SaveBanner"

	_, err := s.db.Exec(saveBanner, &banner.FeatureId, &banner.Tag, &banner.Title, &banner.Text, &banner.Url)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) DeleteBanner(id uint) error {
	const op = "storage.postgres.DeleteBanner"
	_, err := s.db.Exec(deleteBanner, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) GetBanner(featureId, tag uint) (*storage.Banner, error) {
	const op = "storage.postgres.GetBanner"

	var banner storage.Banner
	err := s.db.QueryRow(getBanner, featureId, tag).Scan(&banner.Id, &banner.FeatureId, &banner.Tag, &banner.Title, &banner.Text, &banner.Url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &banner, nil
}

func (s *Storage) GetBannersByFeature(featureId uint) ([]storage.Banner, error) {
	const op = "storage.postgres.GetBannersByFeature"

	var banners []storage.Banner

	row, err := s.db.Query(getBannerByFeature, featureId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for row.Next() {
		var banner storage.Banner
		err = row.Scan(&banner.Id, &banner.FeatureId, &banner.Tag, &banner.Title, &banner.Text, &banner.Url)
		if err != nil {
			return banners, fmt.Errorf("%s: %w", op, err)
		}
		banners = append(banners, banner)
	}

	return banners, nil
}

func (s *Storage) GetBannersByTag(tag uint) ([]storage.Banner, error) {
	const op = "storage.postgres.GetBannersByFeature"

	var banners []storage.Banner

	row, err := s.db.Query(getBannerByTag, tag)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for row.Next() {
		var banner storage.Banner
		err = row.Scan(&banner.Id, &banner.FeatureId, &banner.Tag, &banner.Title, &banner.Text, &banner.Url)
		if err != nil {
			return banners, fmt.Errorf("%s: %w", op, err)
		}
		banners = append(banners, banner)
	}

	return banners, nil
}

func (s *Storage) GetBannerById(id uint) (*storage.Banner, error) {
	const op = "storage.postgres.GetBannerById"

	var banner storage.Banner

	row, err := s.db.Query(getBannerById, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = row.Scan(&banner.Id, &banner.FeatureId, &banner.Tag, &banner.Title, &banner.Text, &banner.Url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &banner, nil
}

func (s *Storage) UpdateBannerById(id uint, banner *storage.Banner) error {
	const op = "storage.postgres.UpdateBannerById"

	stmt, err := s.db.Prepare(updateBanner)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(banner.FeatureId, banner.Tag, banner.Title, banner.Text, banner.Url, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
