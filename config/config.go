package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var TCP_ADDRESS string
var LOG_PATH string

var HOST_ADDRESS string
var HOST_PORT string
var EMAIL_VERIFICATION_DURATION int

var SMTP_SENDER_EMAIL string
var SMTP_SENDER_PASSWORD string
var SMTP_HOST string
var SMTP_PORT string

func init() {
	godotenv.Load()
	HOST_ADDRESS = os.Getenv("HOST_ADDRESS")
	HOST_PORT = os.Getenv("HOST_PORT")
	TCP_ADDRESS = HOST_ADDRESS + ":" + HOST_PORT
	LOG_PATH = os.Getenv("LOG_PATH")
	EMAIL_VERIFICATION_DURATION, _ = strconv.Atoi(os.Getenv("EMAIL_VERIFICATION_DURATION"))
	SMTP_SENDER_EMAIL = os.Getenv("SMTP_SENDER_EMAIL")
	SMTP_SENDER_PASSWORD = os.Getenv("SMTP_SENDER_PASSWORD")
	SMTP_HOST = os.Getenv("SMTP_HOST")
	SMTP_PORT = os.Getenv("SMTP_PORT")
	// Menampilkan nilai variabel lingkungan
}
