package sanitization

import "github.com/microcosm-cc/bluemonday"

// Do this once for each unique policy, and use the policy for the life of the program
// Policy creation/editing is not safe to use in multiple goroutines
var StrictPolicy = bluemonday.StrictPolicy()
var HtmlPolicy = bluemonday.UGCPolicy()
