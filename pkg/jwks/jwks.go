package jwks

import (
	"encoding/json"
	"github.com/MicahParks/keyfunc"
	"github.com/dgrijalva/jwt-go"
	"log"
)

func ParseJwt() {
	// Get the JWKS as JSON.
	var jwksJSON json.RawMessage = []byte(`{"keys":[{"kty":"RSA","e":"AQAB","use":"sig","kid":"MjhhMDk2N2M2NGEwMzgzYjk2OTI3YzdmMGVhOGYxNjI2OTc5Y2Y2MQ","alg":"RS256","n":"zZU9xSgK77PbtkjJgD2Vmmv6_QNe8B54eyOV0k5K2UwuSnhv9RyRA3aL7gDN-qkANemHw3H_4Tc5SKIMltVIYdWlOMW_2m3gDBOODjc1bE-WXEWX6nQkLAOkoFrGW3bgW8TFxfuwgZVTlb6cYkSyiwc5ueFV2xNqo96Qf7nm5E7KZ2QDTkSlNMdW-jIVHMKjuEsy_gtYMaEYrwk5N7VoiYwePaF3I0_g4G2tIrKTLb8DvHApsN1h-s7jMCQFBrY4vCf3RBlYULr4Nz7u8G2NL_L9vURSCU2V2A8rYRkoZoZwk3a3AyJiqeC4T_1rmb8XdrgeFHB5bzXZ7EI0TObhlw"}]}`)

	// Create the JWKS from the resource at the given URL.
	jwks, err := keyfunc.New(jwksJSON)
	if err != nil {
		log.Fatalf("Failed to create JWKS from resource at the given URL.\nError: %s", err.Error())
	}

	// Get a JWT to parse.
	jwtB64 := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6ImI2TnozUDJwMHg1QWpfWENsUmhrVDFzNlNIQSJ9.eyJodHRwOlwvXC93c28yLm9yZ1wvY2xhaW1zXC9hcHBsaWNhdGlvbnRpZXIiOiJVbmxpbWl0ZWQiLCJodHRwOlwvXC93c28yLm9yZ1wvY2xhaW1zXC9rZXl0eXBlIjoiUFJPRFVDVElPTiIsImh0dHA6XC9cL3dzbzIub3JnXC9jbGFpbXNcL3ZlcnNpb24iOiIxLjAiLCJpc3MiOiJ3c28yLm9yZ1wvcHJvZHVjdHNcL2FtIiwiaHR0cDpcL1wvd3NvMi5vcmdcL2NsYWltc1wvYXBwbGljYXRpb25uYW1lIjoiVGFseW9uIiwiaHR0cDpcL1wvd3NvMi5vcmdcL2NsYWltc1wvZW5kdXNlciI6IkZEQkBjYXJib24uc3VwZXIiLCJodHRwOlwvXC93c28yLm9yZ1wvY2xhaW1zXC9lbmR1c2VyVGVuYW50SWQiOiItMTIzNCIsImh0dHA6XC9cL3dzbzIub3JnXC9jbGFpbXNcL3N1YnNjcmliZXIiOiJGREIiLCJodHRwOlwvXC93c28yLm9yZ1wvY2xhaW1zXC90aWVyIjoiR29sZCIsImh0dHA6XC9cL3dzbzIub3JnXC9jbGFpbXNcL2FwcGxpY2F0aW9uaWQiOiIxNDU2IiwiaHR0cDpcL1wvd3NvMi5vcmdcL2NsYWltc1wvdXNlcnR5cGUiOiJBUFBMSUNBVElPTiIsImV4cCI6MTU4OTQ2NjI0MSwiaHR0cDpcL1wvd3NvMi5vcmdcL2NsYWltc1wvYXBpY29udGV4dCI6IlwvY3VycmVudC1hY2NvdW50XC9jaGVxdWVzXC9hdXRvbWF0aWMtZGVwb3NpdHNcL2F0bVwvMS4wIn0=.K1iPtdXiuicuDPaLC6Exw/7UpJVW6Uy1tPpJlfZ29Vqs9M1zR00JpKxvymQMAzbD0GHlXPPsZmhDxOn0WMAPfr1Xi8tiruTLXNbwUPJ/SOovt+zK4JGtrydhc4iv2EROhMUk2uwJUb4DFjqKZRhBvtCW7fRtdtI9yJL4W4OK8Ld90yOb97usPjEPz8S4E4uNrb5lE2rLzIp+EaPwA232lDkhS8gGPIKdlLG1IdEfQ4cFU1VIplvWoHzprF9mGR0ahT2QGgmGE3AcBfkURk8VzIKDG/UcBA9eHu3XGg28j3OvIXWwJhd7Hi+jTqvggi0hplao8ElvjNBw/wNy2UO9WA=="

	// Parse the JWT.
	token, err := jwt.Parse(jwtB64, jwks.KeyFunc)
	if err != nil {
		log.Fatalf("Failed to parse the JWT.\nError: %s", err.Error())
	}

	// Check if the token is valid.
	if !token.Valid {
		log.Fatalf("The token is not valid.")
	}

	log.Println("The token is valid.")
}
