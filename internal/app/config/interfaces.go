package config

type Config interface {
	GetAppPort() string
	GetDatabaseName() string
	GetDatabaseUser() string
	GetDatabaseUserPassword() string
	GetDatabaseProtocol() string
	GetDatabaseHost() string
	GetDbUrl() string
	GetSendSleepIntervalSec() int
	GetUnlockEmailsAfterSec() int
	GetLogLevel() string
}
