package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalmiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalmiddlewares: make([]Middleware, 0),
	}
}

func (m *Manager) Use(middlewares ...Middleware) {
	m.globalmiddlewares = append(m.globalmiddlewares, middlewares...)
}

func (m *Manager) With(next http.Handler, routeMiddlewares ...Middleware) http.Handler {
	handler := next

	// 1. Route-specific middlewares wraparound
	for i := len(routeMiddlewares) - 1; i >= 0; i-- {
		handler = routeMiddlewares[i](handler)
	}

	// 2. Global middlewares wraparound
	for i := len(m.globalmiddlewares) - 1; i >= 0; i-- {
		handler = m.globalmiddlewares[i](handler)
	}

	return handler
}