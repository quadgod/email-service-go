package config

type IConfig interface {
	GetAppPort() string
	GetDatabseName() string
	GetDatabaseUser() string
	GetDatabaseUserPassword() string
	GetDatabaseProtocol() string
	GetDatabaseHost() string
	GetDbUrl() string
	GetRateLimitIntervalMs() int64
	GetMaxEmailsPerInterval() int64
}
