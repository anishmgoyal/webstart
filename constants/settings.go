package constants

import (
	"os"
	"strconv"
	"strings"
)

// Domain contains a mapping of domains for redirection for SSL
var Domain = struct {
	Map map[string]string
}{
	Map: map[string]string{
	// TODO: Add url mappings here
	},
}

// OptionPrefix is the prefix for all environment variables
const OptionPrefix = "SERVEROPT"

// PortNum is the port on which the server shall run
var PortNum = 2646

// SSLPortNum is the port on which SSL should be handled if enabled
var SSLPortNum = 2443

// SSLEnable determines whether or not the server will use SSL
var SSLEnable = false

// SSLKeyFile is the path to a key file if SSL is used
var SSLKeyFile = ""

// SSLCertificate is the path to a certificate if SSL is used
var SSLCertificate = ""

// DatabaseName is the name of the database to use
var DatabaseName = "database"

// DatabaseUsername is the username for the database connection
var DatabaseUsername = "webstartuser"

// DatabasePassword is the password for the database connection
var DatabasePassword = "webstartpassword"

// DatabaseHost is the hostname and port number for the database connection
var DatabaseHost = "localhost:5432"

// DatabaseExtraArgs is a query string with any connection parameters
var DatabaseExtraArgs = "?sslmode=disable"

// LoadEnvironmentSettings loads settings from environment variables; note
// that all options will use the specified OptionPrefix constant
// (constants.OptionPrefix)
func LoadEnvironmentSettings() {
	loadIntSetting(&PortNum, formatOption("PORT_NUM"))
	loadIntSetting(&SSLPortNum, formatOption("SSL_PORT"))

	loadBooleanSetting(&SSLEnable, formatOption("SSL_ENABLE"))
	loadStringSetting(&SSLKeyFile, formatOption("SSL_KEYFILE"))
	loadStringSetting(&SSLCertificate, formatOption("SSL_CERTIFICATE"))

	loadStringSetting(&DatabaseUsername, formatOption("DB_UNAME"))
	loadStringSetting(&DatabasePassword, formatOption("DB_PWORD"))
	loadStringSetting(&DatabaseHost, formatOption("DB_HOST"))
	loadStringSetting(&DatabaseExtraArgs, formatOption("DB_ARGS"))
}

func formatOption(option string) string {
	return OptionPrefix + "_" + option
}

func loadBooleanSetting(setting *bool, envKey string) {
	envVal := os.Getenv(envKey)
	if strings.Compare(envVal, "true") == 0 {
		*setting = true
	} else if strings.Compare(envVal, "false") == 0 {
		*setting = false
	}
}

func loadStringSetting(setting *string, envKey string) {
	envVal := os.Getenv(envKey)
	if len(envVal) > 0 {
		*setting = envVal
	}
}

func loadIntSetting(setting *int, envKey string) {
	envVal := os.Getenv(envKey)
	if len(envVal) > 0 {
		val, err := strconv.Atoi(envVal)
		if err == nil {
			*setting = val
		}
	}
}
