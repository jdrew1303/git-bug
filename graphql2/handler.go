//go:generate gorunpkg github.com/vektah/gqlgen

package graphql2

import (
	"github.com/MichaelMure/git-bug/graphql2/resolvers"
	"github.com/MichaelMure/git-bug/repository"
	"github.com/vektah/gqlgen/handler"
	"net/http"
)

func NewHandler(repo repository.Repo) http.Handler {
	backend := resolvers.NewRootResolver()

	backend.RegisterDefaultRepository(repo)

	return handler.GraphQL(resolvers.NewExecutableSchema(backend))
}