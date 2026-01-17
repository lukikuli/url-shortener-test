package dotenv

func APPENV() string {
	return GetString("APP_ENV", "local")
}

func IsAppEnvProduction() bool {
	return APPENV() == "production"
}

func ISUSEHTTPS() bool {
	return GetBool("IS_USE_HTTPS", false)
}

func APPPORT() string {
	return GetString("APP_PORT", "8080")
}
