package bootstrap

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/anishmgoyal/webstart/constants"
	"github.com/anishmgoyal/webstart/controllers"
)

// GlobalStart begins initialization for the application,
// and notifies main() if an error occurrs.
func GlobalStart() bool {

	log.Println("[STARTUP] Loading Environment Settings")
	constants.LoadEnvironmentSettings()

	// Seed for random generators
	rand.Seed(time.Now().UTC().UnixNano())

	log.Println("[STARTUP] Loading Templates")
	templates := GetTemplates()

	log.Println("[STARTUP] Connecting to DB")
	db := GetDatabaseConnection()

	log.Println("[STARTUP] Initializing Controllers / Services")
	controllers.BaseInitialization(templates, db)

	log.Println("[STARTUP] Creating Routes")
	CreateRoutes()

	var sslRedirect = func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		if index := strings.Index(host, ":"); index > -1 {
			host = host[0:index]
		}

		nextHost := host
		if replaceHost, ok := constants.Domain.Map[nextHost]; ok {
			nextHost = replaceHost
		} else {
			nextHost = "https://" + nextHost
		}
		redirectURL := nextHost

		if constants.SSLPortNum != 443 {
			redirectURL += ":" + strconv.Itoa(constants.SSLPortNum)
		}
		http.Redirect(w, r, redirectURL+r.URL.RequestURI(),
			http.StatusMovedPermanently)
	}

	if constants.SSLEnable {
		log.Println("[STARTUP] Starting server on port " +
			strconv.Itoa(constants.SSLPortNum))
		log.Println("[STARTUP] Using redirect from port " +
			strconv.Itoa(constants.PortNum))

		go http.ListenAndServe(":"+strconv.Itoa(constants.PortNum),
			http.HandlerFunc(sslRedirect))
		http.ListenAndServeTLS(":"+strconv.Itoa(constants.SSLPortNum),
			constants.SSLCertificate, constants.SSLKeyFile, nil)
	} else {
		log.Println("[STARTUP] Starting server on port " +
			strconv.Itoa(constants.PortNum))

		http.ListenAndServe(":"+strconv.Itoa(constants.PortNum), nil)
	}

	// Listen and serve should not return
	log.Panic("[STARTUP] Startup failed.")
	return false
}
