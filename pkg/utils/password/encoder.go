package password

type EncodedPassword string

type Encoder interface {
	EncodePassword(raw string) (EncodedPassword, error)
	IsMatchPassword(raw string, encoded EncodedPassword) bool
}
