package identity

type ProviderId string

const (
	GITHUB ProviderId = "github"
	GOOGLE ProviderId = "google"
)

var (
	Supported = []ProviderId{
		GITHUB,
		GOOGLE,
	}
)
