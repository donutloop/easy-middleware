package easy_middlware

import "net/http"

type Middleware func(http.Handler) http.Handler

// Chain acts as a list of middleware.
// Chain is effectively immutable:
// once created, it will always hold
// the same set of constructors in the same order.
type Chain struct {
	middleware []Middleware
}

// New creates a new chain,
// memorizing the given list of middleware.
// New serves no other function,
// constructors are only called upon a call to Then().
func New(middleware ...Middleware) Chain {
	return Chain{middleware:append([]Middleware{}, middleware...)}
}

// Create a chain, adding the specified middleware
// as the last ones in the request flow.
//
//     stdChain := easy_middleware.New(m1, m2)
//     extChain := Create(stdChain, m3, m4)
//     // requests in stdChain go m1 -> m2
//     // requests in extChain go m1 -> m2 -> m3 -> m4
func Create(chain Chain, middleware ...Middleware) Chain {
	newMiddleware := make([]Middleware, 0, len(chain.middleware) + len(middleware))
	newMiddleware = append(newMiddleware, chain.middleware...)
	newMiddleware = append(newMiddleware, middleware...)

	return Chain{middleware:newMiddleware}
}

// Append extends a chain, adding the specified constructors
// as the last ones in the request flow.
//
// Returns a new chain, leaving the originals one untouched.
//
//     stdChain := easy_middleware.New(m1, m2)
//     stdChain := easy_middleware.New(m3, m4)
//     mergedChain := Merge(m3, m4)
//     // requests in stdChain go m1 -> m2
//     // requests in extChain go m1 -> m2 -> m3 -> m4
func Merge(chainOne Chain, chainTwo Chain) Chain {
	newChain := make([]Middleware, 0, len(chainOne.middleware) + len(chainTwo.middleware))
	newChain = append(newChain, chainOne.middleware...)
	newChain = append(newChain, chainTwo.middleware...)

	return Chain{middleware:chainTwo.middleware}
}

// Then chains the middleware and returns the final http.Handler.
//     New(m1, m2, m3).Then(h)
// is equivalent to:
//     m1(m2(m3(h)))
// When the request comes in, it will be passed to m1, then m2, then m3
// and finally, the given handler
// (assuming every middleware calls the following one).
//
// A chain can be safely reused by calling Then() several times.
//     stdStack := easy_middleware.New(ratelimitHandler, csrfHandler)
//     indexPipe = stdStack.Then(indexHandler)
//     authPipe = stdStack.Then(authHandler)
// Note that constructors are called on every call to Then()
// and thus several instances of the same middleware will be created
// when a chain is reused in this way.
// For proper middleware, this should cause no problems.
//
// Then() treats nil as http.DefaultServeMux.
func (c Chain) Then(endpoint http.Handler) http.Handler {

	if endpoint == nil {
		endpoint = http.DefaultServeMux
	}

	for i := range c.middleware {
		endpoint = c.middleware[len(c.middleware)-1-i](endpoint)
	}

	return endpoint
}

// ThenFunc works identically to Then, but takes
// a HandlerFunc instead of a Handler.
//
// The following two statements are equivalent:
//     c.Then(http.HandlerFunc(fn))
//     c.ThenFunc(fn)
//
// ThenFunc provides all the guarantees of Then.
func (c Chain) ThenFunc(endpointFunc http.HandlerFunc) http.Handler {

	if endpointFunc == nil {
		return c.Then(nil)
	}

	return c.Then(endpointFunc)
}
