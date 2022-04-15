package parser

import "net/http"

type Translator[TPayload any, TInput any] func(p TPayload) (*TInput, error)

func ParseAndTranslate[TPayload any, TInput any](

	r *http.Request,
	f Translator[TPayload, TInput],

) (*TInput, error) {

	p, err := ParseRequestAs[TPayload](r)
	if err != nil {
		return nil, err
	}
	return f(p)
}
