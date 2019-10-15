package swagger

import (
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
)

// ResourceType holds information about resources that can be displayed
type ResourceType struct {
	Display        string
	Endpoint       *endpoints.EndpointInfo
	Verb           string
	DeleteEndpoint *endpoints.EndpointInfo
	PatchEndpoint  *endpoints.EndpointInfo
	PutEndpoint    *endpoints.EndpointInfo
	Children       []ResourceType // Children are auto-loaded (must be able to build the URL => no additional template URL values)
	SubResources   []ResourceType // SubResources are not auto-loaded (these come from the request to the endpoint)
}

/////////////////////////////////////////////////////////////////////////////
// Path models

// Path represents a path that we want to consider emitting in code-gen. It is derived from
type Path struct {
	Name                  string
	CondensedEndpointPath string
	Endpoint              *endpoints.EndpointInfo // The logical endpoint. May be overridden for an operation
	Operations            PathOperations
	Children              []*Path
	SubPaths              []*Path
}

// PathOperations gives details on the operations for a resource
type PathOperations struct {
	Get    PathOperation
	Delete PathOperation
	Patch  PathOperation
	Post   PathOperation
	Put    PathOperation
}

// PathOperation represents an operation on the path (GET, PUT, ...)
type PathOperation struct {
	Permitted bool                    // true if the operation is permitted for the path
	Verb      string                  // Empty unless the Verb is overridden for the operation
	Endpoint  *endpoints.EndpointInfo // nil unless the endpoint is overridden for the operation
}

/*
	Path
		Endpoint (logical)
		Operations
			GET
				Accepted
				Verb (overridden)
				Endpoint (overridden)
			PUT
			POST
			DELETE
*/

/////////////////////////////////////////////////////////////////////////////
// Config

// Config handles configuration of url handling
type Config struct {
	// Overrides is keyed on url
	Overrides map[string]PathOverride
	// AdditionalGetPaths contains extra paths to include as GET
	AdditionalGetPaths []string
	// SuppressAPIVersion true to prevent the api version querystring
	SuppressAPIVersion bool
}

// PathOverride captures Path and/or Verb overrides
type PathOverride struct {
	Path    string // actual url to use
	GetVerb string // Verb to use for logical GET requests
}