package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ip2location "github.com/ip2location/ip2location-go/v9"
)

func CountryLookupMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		country := getRequestCountry(c.Request)

		c.Set("country", country)

		c.Next()
	}

}

func getRequestCountry(r *http.Request) string {
	// Check for the Google App Engine country header
	if country := r.Header.Get("X-Appengine-Country"); country != "" {
		return country
	}

	// If the header is not present or the country is unspecified, fall back to IP-based lookup
	return getCountryByIP(r.RemoteAddr)
}

func getCountryByIP(ip string) string {
	db, err := ip2location.OpenDB(os.Getenv("IP2LOCATION_DB"))
	if err != nil {
		return "ZZ"
	}

	results, err := db.Get_country_short(ip)
	if err != nil {
		return "ZZ"
	}

	return results.Country_short
}
