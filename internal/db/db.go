package db

type Config struct {
	Pkg           string `bson:"package"`
	CountryCode   string `bson:"country_code"`
	PercentileMin int    `bson:"percentile_min"`
	PercentileMax int    `bson:"percentile_max"`
	MainSKU       string `bson:"main_sku"`
}

// ConfigDB defines the interface for database operations
type ConfigDB interface {
	GetConfigs(pkg string, country string, randomPercentile int) ([]*Config, error)
}
