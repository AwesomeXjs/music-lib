package helpers

const (
	FailedToCreateElement = "FAILED_TO_CREATE_ELEMENT" // FailedToCreateElement contains error message for failed to create element
	FailedToDeleteElement = "FAILED_TO_DELETE_ELEMENT" // FailedToDeleteElement contains error message for failed to delete element
	FailedToGetElements   = "FAILED_TO_GET_ELEMENTS"   // FailedToGetElements contains error message for failed to get elements
	JSONParseError        = "JSON_PARSE_ERROR"         // JSONParseError contains error message for failed to parse json
	DefaultValueForFields = "NOT FOUND"                // DefaultValueForFields contains default value for fields

	AppPrefix      = " [ APP ] "      // AppPrefix contains app prefix
	ResponsePrefix = " [ RESPONSE ] " // ResponsePrefix contains response prefix
	InfoPrefix     = " INFO "         // InfoPrefix contains info prefix
	Success        = "SUCCESS"        // Success contains success message
	RequestError   = "REQUEST_ERROR"  // RequestError contains error message for failed to error
	StatusPrefix   = " STATUS "       // StatusPrefix contains status prefix

	UnmarshalError = "UNMARSHAL_ERROR" // UnmarshalError contains error message for unmarshal error
	ReadBodyError  = "READ_BODY_ERROR" // ReadBodyError contains error message for read body error

	PgPrefix            = " [ POSTGRES ] "                     // PgPrefix contains database prefix for logs
	ReconnectDB         = "RECONNECTING TO DATABASE..."        // ReconnectDB contains reconnect db message
	DisconnectDB        = "DISCONNECTED FROM DATABASE"         // DisconnectDB contains disconnect db message
	PgConnectFailed     = "FAILED TO CONNECT TO DATABASE"      // PgConnectFailed contains error message for failed to connect to database
	PgConnectSuccess    = "SUCCESSFULLY CONNECTED TO POSTGRES" // PgConnectSuccess contains success message for successfully connected to database
	PgTransactionFailed = "FAILED TO FETCH TRANSACTION"        // PgTransactionFailed contains error message for failed to fetch transaction
	PgMigrateFailed     = "FAILED TO MIGRATE DATABASE"         // PgMigrateFailed contains error message for failed to migrate database
	NoRowsAffected      = "NO ROWS AFFECTED"                   // NoRowsAffected contains error message for no rows affected
	FailedToRollback    = "FAILED TO ROLLBACK"                 // FailedToRollback contains error message for failed to rollback
	FailedToClose       = "FAILED TO CLOSE"                    // FailedToClose contains error message for failed to close
)
