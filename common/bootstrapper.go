package common

func init() {
	// init AppConfig
	initConfig()
	// init keys
	initKeys()
	// init log level

	// start a MongoDB session
	createDBSession()
	// add indexes
	addIndexes()
}
