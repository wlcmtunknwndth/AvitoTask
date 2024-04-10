package postgres

const (
	getPassword  = "SELECT * FROM Auth WHERE login = $1"
	registerUser = "INSERT INTO Auth(login, pass) VALUES($1, $2)"
	deleteUser   = "DELETE FROM Auth WHERE login = $1"

	getBanner          = "SELECT * FROM Banners WHERE feature_id = $1 AND tag = $2"
	getBannerByFeature = "SELECT * FROM Banners WHERE feature_id = $1"
	getBannerByTag     = "SELECT * FROM Banners WHERE tag = $1"
	saveBanner         = "INSERT INTO Banners(feature_id, tag, title, text, url) VALUES($1, $2, $3, $4, $5)"
	deleteBanner       = "DELETE FROM Banners WHERE id = $1"
)
