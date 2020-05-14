package constants

const (
	TestResourcesTag = "Test Resources"

	ERR_MSG_REQID_NOT_FOUND                = "Request ID mismatch."
	ERR_MSG_INTERNAL_ERR                   = "Internal Error"
	ERR_MSG_UUID_EXIST                     = "UUID already exist"
	ERR__SESSION_FILTER_EMPTY              = "Session filter cannot be empty."
	ERR_MSG_NOT_IN_INACTIVE                = "The rule state is not inactive."
	ERR_MSG_PENDING_IN_PROGRESS            = "The rule is still in progress"
	ERR_MSG_PARAMETERS_CONFLICT            = "Request Parameters conflicts with the backend."
	ERR_MSG_NOT_IN_PROGRESS                = "The request status is not inprogress, maybe duplicate postbacks."
	ERR_REQUIRED_ITEMS_INVALID             = "Some required items are invalid."
	ERR_QUERY_PARAMETER_INVALID            = "Some query parameters are invalid."
	ERR_DOMAIN_NAME_DUPLICATED             = "Domain name already been taken"
	ERR_DOMAIN_NAME_CANNOT_CHANGE          = "Domain name cannot be change or delete due to one or more dns records refs to this domain."
	ERR_HIJACK_TOP_LEVEL_ID_NOT_FOUND      = "Cannot find top level domain id"
	ERR_DOMAIN_RECORD_ALREADY_EXISTS       = "Domain records already exists. Use update instead of create"
	ERR_ANSWER_ALREADY_EXISTS              = "Answer with same ip already exist. No need create"
	ERR_DOMAIN_RECORD_ID_NOT_FOUND         = "Domain recoreds not found with given id"
	ERR_DOMAIN_RECORD_NOT_MATCH_TLD_DOMAIN = "Domain name not end with tld`s domain name"
	ERR_TOP_LEVEL_DOMAIN_NOT_FOUND         = "Top level domain id or name not found"

	HTTP_200 = "It is used to indicate nonspecific success. The response body contains a representation of the resource."
	HTTP_204 = "It is used to indicate nonspecific success. The response body contains nothing."
	HTTP_201 = "Upon success, the HTTP response shall include a Location HTTP header that contains the resource URI of the created resource."
	HTTP_400 = "Bad Request. It is used to indicate that incorrect parameters were passed to the request."
	HTTP_403 = "Forbidden. The operation is not allowed given the current status of the resource."
	HTTP_404 = "Not Found. It is used when a client provided a URI that cannot be mapped to a valid resource URI."
	HTTP_412 = "Precondition Failed. It is used when a condition has failed during conditional requests, e.g. when using ETags to avoid write conflicts."
	HTTP_414 = "It is used to indicate that the server is refusing to process the request because the request URI is longer than the server is willing or able to process."
	HTTP_500 = ERR_MSG_INTERNAL_ERR

	ROUTE_DESC_userId = "Represents a user instance."

	TAG_COMMON_echo = "liveness: echo"
	TAG_COMMON_mock = "liveness: mock"

	STATUS_POSTBACK_INPROGRESS = "inprogress"
	STATUS_POSTBACK_SUCCESS    = "successfully"
	STATUS_POSTBACK_ERROR      = "error"

	PAGE_SIZE   = "The size of per page."
	PAGE_NUMBER = "The number of page."
)
