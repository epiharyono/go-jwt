package get_jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type JWTHeader struct {
	Sub int64 `json:"sub"`
}

func GetID() int64 {
	// Example JWT
	jwt := "eyJhbGciOiJIUzI1NiIsImtpZCI6ImdseGJpSVVZTVhTM0ZPakxrMHNBZHRJWmZGYk9Zc3NaIiwidHlwIjoiSldUIn0.eyJleHAiOjE3MjA0MjIxMzcsInN1YiI6M30.PNLgPdkahPB649QTvCX_Fab9DewRA9xyx2Dv0RnMooY"

	// Split the JWT into its parts
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		fmt.Println("Invalid JWT")
		return 0
	}

	// Decode the header
	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Error decoding header:", err)
		return 0
	}

	// Unmarshal the JSON header into a struct
	var header JWTHeader
	err = json.Unmarshal(headerBytes, &header)
	if err != nil {
		fmt.Println("Error unmarshalling header:", err)
		return 0
	}

	// Print the decoded header
	fmt.Printf("Decoded JWT Header: %+v\n", header.Sub)
	return header.Sub
}
