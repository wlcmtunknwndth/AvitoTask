package storage

//type Storage interface {
//	GetPassword(login string) (string, error)
//	RegisterUser(login, pass string) error
//	DeleteUser(login string) error
//
//	SaveBanner(banner *Banner) error
//	DeleteBanner(id uint) error
//	GetBanner(featureId, tag uint) (*Banner, error)
//	GetBannersByFeature(featureId uint) ([]Banner, error)
//	GetBannersByTag(tag uint) ([]Banner, error)
//}

//type User struct {
//	login string
//	pass  string
//}

type Banner struct {
	Id        uint   `json:"id"`
	FeatureId uint   `json:"feature_id"`
	Tag       uint   `json:"tag_id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Url       string `json:"url"`
}
