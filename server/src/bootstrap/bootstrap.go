package bootstrap

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/util"

type Bootstrap interface {
	Initialice(config util.Config) error
}
