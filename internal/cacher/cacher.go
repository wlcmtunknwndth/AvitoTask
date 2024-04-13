package cacher

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage"
	"log/slog"
	"time"
)

var cached = make(map[string]struct{})

// uuid is string(feature_id + "/" + tag)

type Storage interface {
	RestoreCache() ([]storage.Banner, error)
	SaveCache(uuid string) error
	DeleteCache(uuid string) error
	IsAlreadyCached(uuid string) bool
}

type Cacher struct {
	handler *cache.Cache
	db      Storage
}

// New -- creates new instance of Cacher with Storage interface and cache.Cache vars. expTime -- is the standard expiration time of cached item.
// purgeTime -- is the time the cacher cleans up itself
func New(db Storage, expTime time.Duration, purgeTime time.Duration) *Cacher {
	return &Cacher{
		handler: cache.New(expTime, purgeTime),
		db:      db,
	}
}

// CacheOrder -- caches the order given as an arg and maps order's uuid to cache map.
func (c *Cacher) CacheOrder(banner storage.Banner) {
	c.handler.OnEvicted(c.onEvicted)
	c.handler.Set(fmt.Sprintf("%d/%d", banner.FeatureId, banner.Tag), banner, cache.DefaultExpiration)
	//err := c.db.SaveCache(order.OrderID)
	//if err != nil {
	//	slog.Error("couldn't save backup: ", order.OrderID, err)
	//}
	cached[fmt.Sprintf("%d/%d", banner.FeatureId, banner.Tag)] = struct{}{}
}

// onEvicted -- is a custom func, handling cached item after expiration. It deletes item from cache map and deletes uuid from storage Cache backup.
func (c *Cacher) onEvicted(uuid string, data any) {
	delete(cached, uuid)
	err := c.db.DeleteCache(uuid)
	if err != nil {
		slog.Error("couldn't delete order from cache")
	}
}

// GetOrder -- gets order from cache if found
func (c *Cacher) GetOrder(uuid string) (*storage.Banner, bool) {
	data, found := c.handler.Get(uuid)
	if found {
		order := data.(storage.Banner)
		return &order, true
	}
	return nil, false
}

// Restore -- restores cached item from backup copy in storage. Must be used at the start of ur application.
func (c *Cacher) Restore() error {
	banners, err := c.db.RestoreCache()
	//fmt.Println(orders)
	if err != nil {
		slog.Error("couldn't restore cache: ", err)
		return err
	}

	for i := range banners {
		c.CacheOrder(banners[i])
	}
	return nil
}

// SaveCache -- backups cache to the storage
func (c *Cacher) SaveCache() error {
	var err error
	for key := range cached {
		if c.db.IsAlreadyCached(key) {
			continue
		}
		err = c.db.SaveCache(key)
		if err != nil {
			slog.Error("couldn't save uuid to cache zone: ", key, err)
			continue
		}
	}
	return nil
}
