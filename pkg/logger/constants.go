package logger

const (
	APP_PREFIX = " [ APP ] "

	INFO_PREFIX     = " INFO "
	STATUS_PREFIX   = " STATUS "
	RESPONSE_PREFIX = " [ RESPONSE ] "

	REPO_PREFIX        = "[ REPOSITORY ]"
	REPO_CREATE_FAILED = "FAILED TO CREATE REPOSITORY"

	RECONECT_DB            = "RECONNECTING TO DATABASE..."
	DISCONNECT_DB          = "DISCONNECTED FROM DATABASE"
	PG_PREFIX              = " [ POSTGRES ] "
	PG_CONNECTION_FAILED   = "FAILED TO CONNECT TO DATABASE"
	PG_CONNECT_SUCCESS     = "SUCCESSFULLY CONNECTED TO POSTGRES"
	PG_TRANSACTION_FAILED  = "FAILED TO FETCH TRANSACTION"
	PG_TRANSACTION_SUCCESS = "TRANSACTION SUCCESS"
	PG_COMMIT_FAILED       = "FAILED TO COMMIT TRANSACTION"

	TODO_PREFIX = " [ TODO ] "
)
