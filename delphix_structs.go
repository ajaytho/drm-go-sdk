package delphix

// APIErrorStruct - Description of an error encountered during an API call.
// extends TypedObject
type APIErrorStruct struct {
	// Action to be taken by the user, if any, to fix the underlying
	// problem.
	Action string `json:"action,omitempty"`
	// Extra output, often from a script or other external process, that
	// may give more insight into the cause of this error.
	CommandOutput string `json:"commandOutput,omitempty"`
	// For validation errors, a map of fields to APIError objects. For
	// all other errors, a string with further details of the error.
	Details string `json:"details,omitempty"`
	// Results of diagnostic checks run, if any, if the job failed.
	Diagnoses []*DiagnosisResultStruct `json:"diagnoses,omitempty"`
	// A stable identifier for the class of error encountered.
	Id string `json:"id,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// APISessionStruct - Describes a Delphix web service session and is the result of an
// initial handshake.
// extends TypedObject
type APISessionStruct struct {
	// Client software identification token.
	// required = false
	// maxLength = 64
	Client string `json:"client,omitempty"`
	// Locale as an IETF BCP 47 language tag, defaults to 'en-US'.
	// format = locale
	// required = false
	Locale string `json:"locale,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of the API to use.
	// required = true
	Version *APIVersionStruct `json:"version,omitempty"`
}

// APIVersionStruct - Describes an API version.
// extends TypedObject
type APIVersionStruct struct {
	// Major API version number.
	// minimum = 0
	// required = true
	Major *int `json:"major,omitempty"`
	// Micro API version number.
	// minimum = 0
	// required = true
	Micro *int `json:"micro,omitempty"`
	// Minor API version number.
	// minimum = 0
	// required = true
	Minor *int `json:"minor,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}


// LoginRequestStruct - Represents a Delphix user authentication request.
// extends TypedObject
type LoginRequestStruct struct {
	// Whether to keep session alive for all requests or only via
	// 'KeepSessionAlive' request headers. Defaults to ALL_REQUESTS if
	// omitted.
	// default = ALL_REQUESTS
	// enum = [ALL_REQUESTS KEEP_ALIVE_HEADER_ONLY]
	KeepAliveMode string `json:"keepAliveMode,omitempty"`
	// The password of the user to authenticate.
	// format = password
	// required = true
	Password string `json:"password,omitempty"`
	// The authentication domain.
	// enum = [DOMAIN SYSTEM]
	Target string `json:"target,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the user to authenticate.
	// required = true
	Username string `json:"username,omitempty"`
}

// GroupStruct - Database group.
// extends NamedUserObject
type GroupStruct struct {
	// Optional description for the group.
	// create = optional
	// update = optional
	// maxLength = 1024
	Description string `json:"description,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}


// DiagnosisResultStruct - Details from a diagnosis check that was run due to a failed operation.
// extends TypedObject
type DiagnosisResultStruct struct {
	// True if this was a check that did not pass.
	Failure *bool `json:"failure,omitempty"`
	// Localized message.
	Message string `json:"message,omitempty"`
	// Message code associated with the event.
	MessageCode string `json:"messageCode,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}