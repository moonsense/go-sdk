package sdk

type Config struct {
	// API secret token generated from the Moonsense Cloud web console
	SecretToken string
	// Root API domain (defaults to moonsense.cloud)
	RootDomain string
	// Protocol to use when connecting to the API (defaults to https)
	Protocol string
	// Default Moonsense Cloud Data Plane region to connect to
	DefaultRegion string
}
