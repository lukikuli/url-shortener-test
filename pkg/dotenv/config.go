package dotenv

// APPENV get app env
func APPENV() string {
	return GetString("APP_ENV", "local")
}

// ISNFT nft checker
func ISNFT() bool {
	return GetBool("IS_NFT", false)
}

// IsAppEnvProduction app env checker
func IsAppEnvProduction() bool {
	return APPENV() == "production"
}

// ISUSEHTTPS check use https
func ISUSEHTTPS() bool {
	return GetBool("IS_USE_HTTPS", false)
}

// APPPORT get app port
func APPPORT() string {
	return GetString("APP_PORT", "8080")
}

// APPTLSCERTFILENAME get app tls cert filename
func APPTLSCERTFILENAME() string {
	return GetString("APP_TLS_CERT_FILENAME", "")
}

// APPTLSKEYFILENAME get app tls key filename
func APPTLSKEYFILENAME() string {
	return GetString("APP_TLS_KEY_FILENAME", "")
}

// PREFIX get endpoint prefix
func PREFIX() string {
	return GetString("PREFIX", "/v1")
}

// APPTIMEZONE get app timezone
func APPTIMEZONE() string {
	return GetString("APP_TIMEZONE", "Asia/Jakarta")
}

// FORGEROCKCONFIGFRALG get forgerock config fralg
func FORGEROCKCONFIGFRALG() string {
	return GetString("FORGEROCK_CONFIG_FRALG", "")
}

// FORGEROCKCONFIGTIMEOUT get forgerock config timeout
func FORGEROCKCONFIGTIMEOUT() string {
	return GetString("FORGEROCK_CONFIG_TIMEOUT", "")
}

// FORGEROCKCONFIGDOMAIN get forgerock config domain
func FORGEROCKCONFIGDOMAIN() string {
	return GetString("FORGEROCK_CONFIG_DOMAIN", "")
}

// FORGEROCKCONFIGCLIENTID get forgerock config client id
func FORGEROCKCONFIGCLIENTID() string {
	return GetString("FORGEROCK_CONFIG_CLIENT_ID", "")
}

// APPNAMENEWRELIC get name newrelic
func APPNAMENEWRELIC() string {
	return GetString("APP_NAME_NEWRELIC", "")
}

// APPKEYNEWRELIC get app key newrelic
func APPKEYNEWRELIC() string {
	return GetString("APP_KEY_NEWRELIC", "")
}

// REDISURL get redis url
func REDISURL() string {
	return GetString("REDIS_URL", "")
}

// REDISPASS get redis pass
func REDISPASS() string {
	return GetString("REDIS_PASS", "")
}

// ISREDISTLS check redis tls or not
func ISREDISTLS() bool {
	return GetBool("IS_REDIS_TLS", false)
}

// REDISDB get redis db
func REDISDB() int {
	return GetInt("REDIS_DB", 0)
}

// ISUSEPROXY check use proxy
func ISUSEPROXY() bool {
	return GetBool("IS_USE_PROXY", false)
}

// ISRABBITMQTLS get rabbit mq tls
func ISRABBITMQTLS() bool {
	return GetBool("IS_RABBIT_MQ_TLS", false)
}

// MAILGUNDOMAIN get mailgun domain config
func MAILGUNDOMAIN() string {
	return GetString("MAILGUN_DOMAIN", "")
}

// MAILGUNAPIKEY get mailgun api key config
func MAILGUNAPIKEY() string {
	return GetString("MAILGUN_API_KEY", "")
}

func MAILGUNSENDEREMAIL() string {
	return GetString("MAILGUN_SENDER_EMAIL", "")
}

func REDISEXPIREDLONG() string {
	return GetString("REDIS_EXPIRED_LONG", "1h")
}

func REDISEXPIREDMIDDLE() string {
	return GetString("REDIS_EXPIRED_MIDDLE", "30m")
}

func REDISEXPIREDSHORT() string {
	return GetString("REDIS_EXPIRED_SHORT", "1m")
}

func MONGODBURL() string {
	return GetString("MONGODB_URL", "")
}

func MONGODBDBNAME() string {
	return GetString("MONGODB_DB_NAME", "")
}

func ELASTICADDRESS() string {
	return GetString("ELASTIC_ADDRESS", "https://localhost:9200")
}

func ELASTICUSERNAME() string {
	return GetString("ELASTIC_USERNAME", "elastic")
}

func ELASTICPASSWORD() string {
	return GetString("ELASTIC_PASSWORD", "")
}

// GBURL get growthbook url
func GBURL() string {
	return GetString("GB_URL", "")
}
