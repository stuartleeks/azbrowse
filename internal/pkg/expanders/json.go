package expanders

import (
	"context"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// Check interface
var _ Expander = &JSONExpander{}

// JSONExpander expands an item with "jsonItem" in its metadata
type JSONExpander struct {
	ExpanderBase
}

func (e *JSONExpander) setClient(c *armclient.Client) {
	// noop
}

// Name returns the name of the expander
func (e *JSONExpander) Name() string {
	return "JSONExpander"
}

// DoesExpand checks if this is an RG
func (e *JSONExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if _, ok := currentItem.Metadata["jsonItem"]; ok {
		return true, nil
	}
	return false, nil
}

// Expand returns Resources in the RG
func (e *JSONExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	return ExpanderResult{
		Err:               nil,
		Response:          ExpanderResponse{Response: currentItem.Metadata["jsonItem"], ResponseType: ResponseJSON},
		SourceDescription: "Deployments Subdeployment",
		IsPrimaryResponse: true,
	}
}

func (e *JSONExpander) testCases() (bool, *[]expanderTestCase) {
	return false, nil
}
