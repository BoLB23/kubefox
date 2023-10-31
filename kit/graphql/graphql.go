package graphql

import (
	"context"
	"net/http"

	"github.com/hasura/go-graphql-client"
	"github.com/xigxog/kubefox/kit"
)

type Client struct {
	ktx     kit.Kontext
	wrapped *graphql.Client
}

func New(ktx kit.Kontext, dependency kit.Dependency) *Client {
	c := graphql.NewClient("http://hasura-graphql-engine.default:8080/v1/graphql", &http.Client{
		Transport: ktx.Transport(dependency),
	}).WithRequestModifier(func(r *http.Request) {
		r.Header.Set("x-hasura-admin-secret", "hasura")
	})

	return &Client{
		ktx: ktx,
		// TODO url should be set in broker by config
		wrapped: c,
	}
}

// Query executes a single GraphQL query request, with a query derived from q,
// populating the response into it. q should be a pointer to struct that
// corresponds to the GraphQL schema.
func (c *Client) Query(q interface{}, variables map[string]interface{}, options ...graphql.Option) error {
	return c.QueryCtx(c.ktx.Context(), q, variables, options...)
}

// See Query for details.
func (c *Client) QueryCtx(ctx context.Context, q interface{}, variables map[string]interface{}, options ...graphql.Option) error {
	return c.wrapped.Query(ctx, q, variables, options...)
}

// Mutate executes a single GraphQL mutation request, with a mutation derived
// from m, populating the response into it. m should be a pointer to struct that
// corresponds to the GraphQL schema.
func (c *Client) Mutate(m interface{}, variables map[string]interface{}, options ...graphql.Option) error {
	return c.MutateCtx(c.ktx.Context(), m, variables, options...)
}

// See Mutate for details.
func (c *Client) MutateCtx(ctx context.Context, m interface{}, variables map[string]interface{}, options ...graphql.Option) error {
	return c.wrapped.Mutate(ctx, m, variables, options...)
}

// Query executes a single GraphQL query request, with a query derived from q,
// populating the response into it. q should be a pointer to struct that
// corresponds to the GraphQL schema. return raw bytes message.
func (c *Client) QueryRaw(q interface{}, variables map[string]interface{}, options ...graphql.Option) ([]byte, error) {
	return c.QueryRawCtx(c.ktx.Context(), q, variables, options...)
}

// See QueryRaw for details.
func (c *Client) QueryRawCtx(ctx context.Context, q interface{}, variables map[string]interface{}, options ...graphql.Option) ([]byte, error) {
	return c.wrapped.QueryRaw(ctx, q, variables, options...)
}

// MutateRaw executes a single GraphQL mutation request, with a mutation derived
// from m, populating the response into it. m should be a pointer to struct that
// corresponds to the GraphQL schema. return raw bytes message.
func (c *Client) MutateRaw(m interface{}, variables map[string]interface{}, options ...graphql.Option) ([]byte, error) {
	return c.MutateRawCtx(c.ktx.Context(), m, variables, options...)
}

// See MutateRaw for details.
func (c *Client) MutateRawCtx(ctx context.Context, m interface{}, variables map[string]interface{}, options ...graphql.Option) ([]byte, error) {
	return c.wrapped.MutateRaw(ctx, m, variables, options...)
}
