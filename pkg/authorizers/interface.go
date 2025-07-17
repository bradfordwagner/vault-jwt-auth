package authorizers

import "context"

type Authorizer interface {
	// Authorize - authorizes the user and returns a token or an error
	Authorize(ctx context.Context) (string, error)
}
