package helpers

const (
	// FailedToCreateElement contains error message for failed to create element
	FailedToCreateElement = "FAILED_TO_CREATE_ELEMENT"

	// FailedToDeleteElement contains error message for failed to delete element
	FailedToDeleteElement = "FAILED_TO_DELETE_ELEMENT"

	// FailedToGetElements contains error message for failed to get elements
	FailedToGetElements = "FAILED_TO_GET_ELEMENTS"

	// FailedToRollback contains error message for failed to rollback
	FailedToRollback = "FAILED TO ROLLBACK"

	// FailedToClose contains error message for failed to close
	FailedToClose = "FAILED TO CLOSE"

	// RequestError contains error message for failed to error
	RequestError = "REQUEST_ERROR"

	// JSONParseError contains error message for failed to parse json
	JSONParseError = "JSON_PARSE_ERROR"

	// Success contains success message
	Success = "SUCCESS"

	// DefaultValueForFields contains default value for fields
	DefaultValueForFields = "NOT FOUND"

	// AppPrefix contains app prefix
	AppPrefix = " [ APP ] "

	// InfoPrefix contains info prefix
	InfoPrefix = " INFO "

	// StatusPrefix contains status prefix
	StatusPrefix = " STATUS "

	// ResponsePrefix contains response prefix
	ResponsePrefix = " [ RESPONSE ] "

	// UnmarshalError contains error message for unmarshal error
	UnmarshalError = "UNMARSHAL_ERROR"

	// ReadBodyError contains error message for read body error
	ReadBodyError = "READ_BODY_ERROR"

	// ReconnectDB contains reconnect db message
	ReconnectDB = "RECONNECTING TO DATABASE..."
	// DisconnectDB contains disconnect db message
	DisconnectDB = "DISCONNECTED FROM DATABASE"

	// PgPrefix contains database prefix for logs
	PgPrefix = " [ POSTGRES ] "

	// PgConnectFailed contains error message for failed to connect to database
	PgConnectFailed = "FAILED TO CONNECT TO DATABASE"

	// PgConnectSuccess contains success message for successfully connected to database
	PgConnectSuccess = "SUCCESSFULLY CONNECTED TO POSTGRES"

	// PgTransactionFailed contains error message for failed to fetch transaction
	PgTransactionFailed = "FAILED TO FETCH TRANSACTION"

	// PgMigrateFailed contains error message for failed to migrate database
	PgMigrateFailed = "FAILED TO MIGRATE DATABASE"

	// NoRowsAffected contains error message for no rows affected
	NoRowsAffected = "NO ROWS AFFECTED"
)
