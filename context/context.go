// Package context provides a layer of abstraction above go's native context, providing access to
// metadata which can be set by passing headers to the M3O API.
package context

import "context"

// Context provided to a service handler
type Context struct {
	context.Context
}

// GetRequestHeader from the request
func (c *Context) GetRequestHeader(key string) string {
	return ""
}

// SetRequestHeader on the context
func (c *Context) SetRequestHeader(key, value string) {
}

// SetResponseHeader on the context
func (c *Context) SetResponseHeader(key, value string) {
}

// SetResponseCode on the context
func (c *Context) SetResponseCode(code int) {
}
