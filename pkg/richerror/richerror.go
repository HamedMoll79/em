package richerror

import "log"

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
	KindRateLimit
	KindBadRequest
	KindUnauthorized
)

type Op string

type RichError struct {
	operation    Op // a unique string describing a method or a function.
	wrappedError error
	message      string
	kind         Kind // category of errors.
	meta         map[string]interface{}
}

func New(op Op) RichError {
	return RichError{operation: op}
}

func (r RichError) WithOp(op Op) RichError {
	r.operation = op
	return r
}

func (r RichError) WithErr(err error) RichError {
	r.wrappedError = err
	return r
}

func (r RichError) WithMessage(message string) RichError {
	r.message = message
	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind
	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta
	return r
}

func (r RichError) Meta() map[string]interface{} {
	if len(r.meta) != 0 {
		return r.meta
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {
		return map[string]interface{}{}
	}

	return re.Meta()
}

func (r RichError) Error() string {
	return r.message
}

func (r RichError) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {
		return 0
	}

	return re.Kind()
}

func (r RichError) Message() string {
	log.Println(r.wrappedError)
	if r.message != "" {
		return r.message
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {
		return r.wrappedError.Error()
	}

	return re.Message()
}
func (r RichError) WrappedError() string {
	log.Println(r.wrappedError)
	return r.wrappedError.Error()
}
