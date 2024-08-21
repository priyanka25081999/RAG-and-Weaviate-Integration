package configs

import (
	"os"
)

func LoadConfig() {
	// Load environment variables or other configurations
	os.Setenv("WEAVIATE_API_KEY", "your-api-key")
	os.Setenv("SHAREPOINT_CLIENT_ID", "your-client-id")
	os.Setenv("SHAREPOINT_CLIENT_SECRET", "your-client-secret")
}
