package env

import "github.com/spf13/viper"

func AppHost() string {
	return viper.GetString("APP_HOST")
}

func AppPort() string {
	return viper.GetString("APP_PORT")
}

func AppName() string {
	return viper.GetString("APP_NAME")
}

func DBHost() string {
	return viper.GetString("DB_HOST")
}

func DBPort() string {
	return viper.GetString("DB_PORT")
}

func DBUser() string {
	return viper.GetString("DB_USER")
}

func DBPassword() string {
	return viper.GetString("DB_PASSWORD")
}

func DBName() string {
	return viper.GetString("DB_NAME")
}

func DBSslMode() string {
	return viper.GetString("DB_SSL_MODE")
}

func DBDebug() int {
	return viper.GetInt("DB_DEBUG")
}

// GinMode indicates environment mode
// possible value is one of (debug|release|test)
func GinMode() string {
	return viper.GetString("GIN_MODE")
}

func JWTSecret() string {
	return viper.GetString("JWT_SECRET")
}

func LogEncoder() string {
	return viper.GetString("LOG_ENCODER")
}

func RedisPort() string {
	return viper.GetString("REDIS_PORT")
}

func RedisHost() string {
	return viper.GetString("REDIS_HOST")
}

func RedisDB() int {
	return viper.GetInt("REDIS_DB")
}

func RedisPassword() string {
	return viper.GetString("REDIS_PASSWORD")
}

func SmtpHost() string {
	return viper.GetString("SMTP_HOST")
}

func SmtpPort() int {
	return viper.GetInt("SMTP_PORT")
}

func SmtpUsername() string {
	return viper.GetString("SMTP_USERNAME")
}

func SmtpPassword() string {
	return viper.GetString("SMTP_PASSWORD")
}

func ResetPasswordTokenSecret() string {
	return viper.GetString("RESET_PASSWORD_TOKEN_SECRET")
}

func AWSRegion() string {
	return viper.GetString("AWS_REGION")
}

func AWSAccessKeyId() string {
	return viper.GetString("AWS_ACCESS_KEY_ID")
}

func AWSSecretAccessKey() string {
	return viper.GetString("AWS_SECRET_ACCESS_KEY")
}

func AWSBucketName() string {
	return viper.GetString("AWS_BUCKET_NAME")
}
