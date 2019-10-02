package handler

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

type middlewareContext struct {
	InvokedNext   bool
	ResultContext graphql.RequestContext
	Response      graphql.Response
}

func testMiddleware(m Middleware, initialContexts ...graphql.RequestContext) middlewareContext {
	var c middlewareContext
	initial := &graphql.RequestContext{}
	if len(initialContexts) > 0 {
		initial = &initialContexts[0]
	}

	m(func(ctx context.Context, writer Writer) {
		c.ResultContext = *graphql.GetRequestContext(ctx)
		c.InvokedNext = true
	})(graphql.WithRequestContext(context.Background(), initial), func(response *graphql.Response) {
		c.Response = *response
	})

	return c
}