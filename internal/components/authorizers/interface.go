package authorizers

type Authorizer interface {
	// Authorize - authorizes the user and returns a token or an error
	Authorize() (string, error)
}
