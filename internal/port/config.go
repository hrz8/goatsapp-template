package port

type AppConfigor interface {
	GetAppPort() string
	GetBasePath() string
	GetDatabaseURL() string
}
