package config

import (
	"os"
)

var Mode = os.Getenv("MODE")

const RELEASE_MODE = "release"

func GetSecret() string {
	if isRelease() {
		return os.Getenv("SECRET")
	}
	return "secret"
}

func GetUsername() string {
	if isRelease() {
		return os.Getenv("USERNAME")
	}
	return "name"
}

func GetPassword() string {
	if isRelease() {
		return os.Getenv("PASSWORD")
	}
	return "111111"
}

func GetDatabaseAddr() string {
	if isRelease() {
		return os.Getenv("DB_ADDR")
	}
	return "mongodb://localhost:27017"
}

func GetDatabaseName() string {
	return "test"
}

func isRelease() bool {
	return Mode == RELEASE_MODE
}
