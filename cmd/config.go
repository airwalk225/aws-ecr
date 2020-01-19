package cmd

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

type AWSConfig struct {
	ID         int
	Region     string
	Repository string
	Session    *session.Session
	Context    aws.Context
}

type LogConfig struct {
	Level log.Level
}

type Config struct {
	AWS       AWSConfig
	Log       LogConfig
}

var (
	cnf Config
)

func NewConfig() {
	logLevel, _ := log.ParseLevel(getEnv("DEBUG_MODE", "info"))
	cnf = Config{
		AWS: AWSConfig{
			ID:         getEnvAsInt("ECR_ACCOUNT_ID", 123123123),
			Region:     getEnv("ECR_AWS_REGION", "eu-west-2"),
			Repository: getEnv("ECR_REPOSITORY", ""),
			Session: session.Must(session.NewSessionWithOptions(session.Options{
				SharedConfigState: session.SharedConfigEnable,
			})),
			Context: aws.BackgroundContext(),
		},
		Log: LogConfig{
			Level: logLevel,
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
