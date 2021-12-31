package appConfig

var AuthCodePath = AppPathConfig["authCodeFile"]

var LoginAuthCodeConfig = map[string]interface{}{

	"chars":"ABCDEFGHIJKMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz0123456789",
	"width":340,
	"height":80,
	"length":6,
	"timeOut":300,
}

var AdminAuthCodeConfig = map[string]interface{}{

	"chars":"ABCDEFGHIJKMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz0123456789",
	"width":380,
	"height":80,
	"length":8,
	"timeOut":300,
}
