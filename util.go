package gopaypal

import "github.com/dchest/uniuri"

// CreateNonce creates and returns an arbitrary ID that may only be used once
func CreateNonce() string {
	return uniuri.NewLen(nonceLength)
}
