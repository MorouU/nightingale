package appConfig

var SessionConfig = map[string]interface{}{
	"serectKey":"nightingale",
	"path":"/",
	"domain":CookieDomain,
	"maxAge":60 * 60 * 12,
	"secure": false,
	"httpOnly": true,
	"cookieName":"niSession",
}
