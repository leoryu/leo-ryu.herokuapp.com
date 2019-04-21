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
	return "abc"
}

func GetPassword() string {
	if isRelease() {
		return os.Getenv("PASSWORD")
	}
	return "123"
}

func GetDatabaseAddr() string {
	if isRelease() {
		return os.Getenv("DB_ADDR")
	}
	return "mongodb://localhost:27017"
}

func GetDatabaseName() string {
	if isRelease() {
		return os.Getenv("DB_NAME")
	}
	return "test"
}

func GetStaticFilesPath() string {
	return os.Getenv("BLOG_STATIC_FILES_PATH")
}

func GetListenAddr() string {
	if isRelease() {
		return ":" + os.Getenv("PORT")
	}
	return ":7777"
}

func isRelease() bool {
	return Mode == RELEASE_MODE
}
