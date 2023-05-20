package security

type TokenGenerator interface {
	Generate(principal *Principal) (*string, error)
}
