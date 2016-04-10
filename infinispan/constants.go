package infinispan

const (

	//DefaultCache is an empty string
	DefaultCache = ""

	// Magic numbers for requests and responses
	RequestMagic  = 0XA0
	ResponseMagic = 0XA1

	// Different Protocol Versions
	Protocol25 = 25

	//Requests
	PutRequest                  = 0X01
	GetRequest                  = 0X03
	PutIfAbsentRequest          = 0X05
	ReplaceRequest              = 0X07
	ReplaceIfUnmodifiedRequest  = 0X09
	RemoveRequest               = 0X0B
	RemoveIfUnmodifiedRequest   = 0X0D
	ContainsKeyRequest          = 0X0F
	GetWithVersionRequest       = 0X11
	ClearRequest                = 0X13
	StatsRequest                = 0X15
	PingRequest                 = 0X17
	BulkGetRequest              = 0X19
	GetWithMetadataRequest      = 0X1B
	BulkGetKeysRequest          = 0X1D
	QueryRequest                = 0X1F
	AuthMechListRequest         = 0X21
	AuthRequest                 = 0X23
	AddClientListenerRequest    = 0X25
	RemoveClientListenerRequest = 0X27
	SizeRequest                 = 0X29
	ExecRequest                 = 0X2B
	PutAllRequest               = 0X2D
	GetAllRequest               = 0X2F
	IterationStartRequest       = 0X31
	IterationNextRequest        = 0X33
	IterationEndRequest         = 0X35

	//Responses
	PutResponse                     = 0X02
	GetResponse                     = 0X04
	PutIfAbsentResponse             = 0X06
	ReplaceResponse                 = 0X08
	ReplaceIfUnmodifiedResponse     = 0X0A
	RemoveResponse                  = 0X0C
	RemoveIfUnmodifiedResponse      = 0X0E
	ContainsKeyResponse             = 0X10
	GetWithVersionResponse          = 0X12
	ClearResponse                   = 0X14
	StatsResponse                   = 0X16
	PingResponse                    = 0X18
	BulkGetResponse                 = 0X1A
	GetWithMetadataResponse         = 0X1C
	BulkGetKeysResponse             = 0X1E
	QueryResponse                   = 0X20
	AuthMechListResponse            = 0X22
	AuthResponse                    = 0X24
	AddClientListenerResponse       = 0X26
	RemoveClientListenerResponse    = 0X28
	SizeResponse                    = 0X2A
	ExecResponse                    = 0X2C
	PutAllResponse                  = 0X2E
	GetAllResponse                  = 0X30
	IterationStartResponse          = 0X32
	IterationNextResponse           = 0X34
	IterationEndResponse            = 0X36
	ErrorResponse                   = 0X50
	CacheEntryCreatedEventResponse  = 0X60
	CacheEntryModifiedEventResponse = 0X61
	CacheEntryRemovedEventResponse  = 0X62
	CacheEntryExpiredEventResponse  = 0X63

	//Response Status
	NoErrorStatus                 = 0X00
	NotPutRemovedReplacedStatus   = 0X01
	KeyDoesNotExistStatus         = 0X02
	SuccessWithPrevious           = 0X03
	NotExecutedWithPrevious       = 0X04
	InvalidIteration              = 0X05
	NoErrorStatusCompat           = 0X06
	SuccessWithPreviousCompat     = 0X07
	NotExecutedWithPreviousCompat = 0X08

	InvalidMagicOrMessageIdStatus = 0X81
	RequestParsingErrorStatus     = 0X84
	UnknownCommandStatus          = 0X82
	ServerErrorStatus             = 0X85
	UnknownVersionStatus          = 0X83
	CommandTimeoutStatus          = 0X86
	NodeSuspected                 = 0X87
	IllegalLifecycleState         = 0X88

	ClientIntelligenceBasic                 = 0X01
	ClientIntelligenceTopologyAware         = 0X02
	ClientIntelligenceHashDistributionAware = 0X03

	DefaultDuration  = 0x7
	InfiniteDuration = 0x8

	NoTopology            = 0
	DefaultCacheTopology  = 1
	SwitchClusterTopology = 2
)

var responses = map[int]string{
	PutResponse:                     "PutResponse",
	GetResponse:                     "GetResponse",
	PutIfAbsentResponse:             "PutIfAbsentResponse",
	ReplaceResponse:                 "ReplaceResponse",
	ReplaceIfUnmodifiedResponse:     "ReplaceIfUnmodifiedResponse",
	RemoveResponse:                  "RemoveResponse",
	RemoveIfUnmodifiedResponse:      "RemoveIfUnmodifiedResponse",
	ContainsKeyResponse:             "ContainsKeyResponse",
	GetWithVersionResponse:          "GetWithVersionResponse",
	ClearResponse:                   "ClearResponse",
	StatsResponse:                   "StatsResponse",
	PingResponse:                    "PingResponse",
	BulkGetResponse:                 "BulkGetResponse",
	GetWithMetadataResponse:         "GetWithMetadataResponse",
	BulkGetKeysResponse:             "BulkGetKeysResponse",
	QueryResponse:                   "QueryResponse",
	AuthMechListResponse:            "AuthMechListResponse",
	AuthResponse:                    "AuthResponse",
	AddClientListenerResponse:       "AddClientListenerResponse",
	RemoveClientListenerResponse:    "RemoveClientListenerResponse",
	SizeResponse:                    "SizeResponse",
	ExecResponse:                    "ExecResponse",
	PutAllResponse:                  "PutAllResponse",
	GetAllResponse:                  "GetAllResponse",
	IterationStartResponse:          "IterationStartResponse",
	IterationNextResponse:           "IterationNextResponse",
	IterationEndResponse:            "IterationEndResponse",
	ErrorResponse:                   "ErrorResponse",
	CacheEntryCreatedEventResponse:  "CacheEntryCreatedEventResponse",
	CacheEntryModifiedEventResponse: "CacheEntryModifiedEventResponse",
	CacheEntryRemovedEventResponse:  "CacheEntryRemovedEventResponse",
	CacheEntryExpiredEventResponse:  "CacheEntryExpiredEventResponse",
}

var status = map[int]string{
	NoErrorStatus:                 "NoErrorStatus",
	NotPutRemovedReplacedStatus:   "NotPutRemovedReplacedStatus",
	KeyDoesNotExistStatus:         "KeyDoesNotExistStatus",
	SuccessWithPrevious:           "SuccessWithPrevious",
	NotExecutedWithPrevious:       "NotExecutedWithPrevious",
	InvalidIteration:              "InvalidIteration",
	NoErrorStatusCompat:           "NoErrorStatusCompat",
	SuccessWithPreviousCompat:     "SuccessWithPreviousCompat",
	NotExecutedWithPreviousCompat: "NotExecutedWithPreviousCompat",
	InvalidMagicOrMessageIdStatus: "InvalidMagicOrMessageIdStatus",
	RequestParsingErrorStatus:     "RequestParsingErrorStatus",
	UnknownCommandStatus:          "UnknownCommandStatus",
	ServerErrorStatus:             "ServerErrorStatus",
	UnknownVersionStatus:          "UnknownVersionStatus",
	CommandTimeoutStatus:          "CommandTimeoutStatus",
	NodeSuspected:                 "NodeSuspected",
	IllegalLifecycleState:         "IllegalLifecycleState",
}
