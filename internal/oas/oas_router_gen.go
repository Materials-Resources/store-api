// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "account/"
				origElem := elem
				if l := len("account/"); len(elem) >= l && elem[0:l] == "account/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'b': // Prefix: "branch"
					origElem := elem
					if l := len("branch"); len(elem) >= l && elem[0:l] == "branch" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "PUT":
							s.handleSetActiveBranchRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "PUT")
						}

						return
					}
					switch elem[0] {
					case 'e': // Prefix: "es"
						origElem := elem
						if l := len("es"); len(elem) >= l && elem[0:l] == "es" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleListCustomerBranchesRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/active"
							origElem := elem
							if l := len("/active"); len(elem) >= l && elem[0:l] == "/active" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetActiveBranchesRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				case 'o': // Prefix: "orders"
					origElem := elem
					if l := len("orders"); len(elem) >= l && elem[0:l] == "orders" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleListOrdersRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						origElem := elem
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetOrderRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				case 'q': // Prefix: "quotes"
					origElem := elem
					if l := len("quotes"); len(elem) >= l && elem[0:l] == "quotes" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleListQuotesRequest([0]string{}, elemIsEscaped, w, r)
						case "POST":
							s.handleCreateQuoteRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET,POST")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						origElem := elem
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetQuoteRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'p': // Prefix: "products/"
				origElem := elem
				if l := len("products/"); len(elem) >= l && elem[0:l] == "products/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleGetProductRequest([1]string{
							args[0],
						}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

				elem = origElem
			case 's': // Prefix: "search/products"
				origElem := elem
				if l := len("search/products"); len(elem) >= l && elem[0:l] == "search/products" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleSearchProductsRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "account/"
				origElem := elem
				if l := len("account/"); len(elem) >= l && elem[0:l] == "account/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'b': // Prefix: "branch"
					origElem := elem
					if l := len("branch"); len(elem) >= l && elem[0:l] == "branch" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "PUT":
							r.name = SetActiveBranchOperation
							r.summary = "Set active branch for current user"
							r.operationID = "setActiveBranch"
							r.pathPattern = "/account/branch"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case 'e': // Prefix: "es"
						origElem := elem
						if l := len("es"); len(elem) >= l && elem[0:l] == "es" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = ListCustomerBranchesOperation
								r.summary = "Get available branches for customer"
								r.operationID = "listCustomerBranches"
								r.pathPattern = "/account/branches"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/active"
							origElem := elem
							if l := len("/active"); len(elem) >= l && elem[0:l] == "/active" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetActiveBranchesOperation
									r.summary = "Get active branch for user"
									r.operationID = "getActiveBranches"
									r.pathPattern = "/account/branches/active"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				case 'o': // Prefix: "orders"
					origElem := elem
					if l := len("orders"); len(elem) >= l && elem[0:l] == "orders" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = ListOrdersOperation
							r.summary = "Get a list of orders"
							r.operationID = "listOrders"
							r.pathPattern = "/account/orders"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						origElem := elem
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = GetOrderOperation
								r.summary = "Get an order by ID"
								r.operationID = "getOrder"
								r.pathPattern = "/account/orders/{id}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				case 'q': // Prefix: "quotes"
					origElem := elem
					if l := len("quotes"); len(elem) >= l && elem[0:l] == "quotes" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = ListQuotesOperation
							r.summary = "Get a list of quotes"
							r.operationID = "listQuotes"
							r.pathPattern = "/account/quotes"
							r.args = args
							r.count = 0
							return r, true
						case "POST":
							r.name = CreateQuoteOperation
							r.summary = "Create a new quote"
							r.operationID = "createQuote"
							r.pathPattern = "/account/quotes"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						origElem := elem
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = GetQuoteOperation
								r.summary = "Get quote by ID"
								r.operationID = "getQuote"
								r.pathPattern = "/account/quotes/{id}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'p': // Prefix: "products/"
				origElem := elem
				if l := len("products/"); len(elem) >= l && elem[0:l] == "products/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "GET":
						r.name = GetProductOperation
						r.summary = "Get a product by ID"
						r.operationID = "getProduct"
						r.pathPattern = "/products/{id}"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 's': // Prefix: "search/products"
				origElem := elem
				if l := len("search/products"); len(elem) >= l && elem[0:l] == "search/products" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "POST":
						r.name = SearchProductsOperation
						r.summary = "Search for products"
						r.operationID = "searchProducts"
						r.pathPattern = "/search/products"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}
