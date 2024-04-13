package postgres

const (
	getPassword  = "SELECT password FROM Auth WHERE username = $1"
	registerUser = "INSERT INTO Auth(username, password) VALUES($1, $2)"
	deleteUser   = "DELETE FROM Auth WHERE username = $1"
	isAdmin      = "SELECT isAdmin FROM Auth WHERE username = $1"

	getBanner          = "SELECT * FROM Banners WHERE feature_id = $1 AND tag = $2"
	getBannerByFeature = "SELECT * FROM Banners WHERE feature_id = $1"
	getBannerByTag     = "SELECT * FROM Banners WHERE tag = $1"
	getBannerById      = "SELECT * FROM Banners WHERE id = $1"
	saveBanner         = "INSERT INTO Banners(feature_id, tag, title, text, url) VALUES($1, $2, $3, $4, $5)"
	updateBanner       = "UPDATE Banners SET feature_id = $1, tag = $2, title = $3, text = $4, url = $5 WHERE id = $6"
	deleteBanner       = "DELETE FROM Banners WHERE id = $1"

	restoreCache    = "SELECT * FROM Cache"
	deleteCache     = "DELETE * FROM Cache WHERE feature_id = $1 AND tag = $2"
	saveCache       = "INSERT INTO Cache(feature_id, tag) VALUES($1, $2)"
	isAlreadyCached = "SELECT * FROM Cache WHERE feature_id = $1 AND tag = $2"
)
