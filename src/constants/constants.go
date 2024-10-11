package constants

func Environment() map[string]string {
	return map[string]string{
		"development": "development",
		"production":  "production",
		"test":        "test",
	}
}
