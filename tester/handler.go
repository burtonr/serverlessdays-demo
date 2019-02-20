package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	input := ""

	if len(req) > 0 {
		input = fmt.Sprintf("You said: %s", string(req))
	}
	return fmt.Sprintf("Hello ServerlessDays ATX! %s", input)
}
