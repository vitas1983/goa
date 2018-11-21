// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// chatter HTTP client types
//
// Command:
// $ goa gen goa.design/goa/examples/chatter/design -o
// $(GOPATH)/src/goa.design/goa/examples/chatter

package client

import (
	goa "goa.design/goa"
	chattersvc "goa.design/goa/examples/chatter/gen/chatter"
	chattersvcviews "goa.design/goa/examples/chatter/gen/chatter/views"
)

// SummaryResponseBody is the type of the "chatter" service "summary" endpoint
// HTTP response body.
type SummaryResponseBody []*ChatSummaryResponseBody

// HistoryResponseBody is the type of the "chatter" service "history" endpoint
// HTTP response body.
type HistoryResponseBody struct {
	// Message sent to the server
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Length of the message sent
	Length *int `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Time at which the message was sent
	SentAt *string `form:"sent_at,omitempty" json:"sent_at,omitempty" xml:"sent_at,omitempty"`
}

// LoginUnauthorizedResponseBody is the type of the "chatter" service "login"
// endpoint HTTP response body for the "unauthorized" error.
type LoginUnauthorizedResponseBody string

// EchoerInvalidScopesResponseBody is the type of the "chatter" service
// "echoer" endpoint HTTP response body for the "invalid-scopes" error.
type EchoerInvalidScopesResponseBody string

// EchoerUnauthorizedResponseBody is the type of the "chatter" service "echoer"
// endpoint HTTP response body for the "unauthorized" error.
type EchoerUnauthorizedResponseBody string

// ListenerInvalidScopesResponseBody is the type of the "chatter" service
// "listener" endpoint HTTP response body for the "invalid-scopes" error.
type ListenerInvalidScopesResponseBody string

// ListenerUnauthorizedResponseBody is the type of the "chatter" service
// "listener" endpoint HTTP response body for the "unauthorized" error.
type ListenerUnauthorizedResponseBody string

// SummaryInvalidScopesResponseBody is the type of the "chatter" service
// "summary" endpoint HTTP response body for the "invalid-scopes" error.
type SummaryInvalidScopesResponseBody string

// SummaryUnauthorizedResponseBody is the type of the "chatter" service
// "summary" endpoint HTTP response body for the "unauthorized" error.
type SummaryUnauthorizedResponseBody string

// HistoryInvalidScopesResponseBody is the type of the "chatter" service
// "history" endpoint HTTP response body for the "invalid-scopes" error.
type HistoryInvalidScopesResponseBody string

// HistoryUnauthorizedResponseBody is the type of the "chatter" service
// "history" endpoint HTTP response body for the "unauthorized" error.
type HistoryUnauthorizedResponseBody string

// ChatSummaryResponseBody is used to define fields on response body types.
type ChatSummaryResponseBody struct {
	// Message sent to the server
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Length of the message sent
	Length *int `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Time at which the message was sent
	SentAt *string `form:"sent_at,omitempty" json:"sent_at,omitempty" xml:"sent_at,omitempty"`
}

// NewLoginUnauthorized builds a chatter service login endpoint unauthorized
// error.
func NewLoginUnauthorized(body LoginUnauthorizedResponseBody) chattersvc.Unauthorized {
	v := chattersvc.Unauthorized(body)
	return v
}

// NewEchoerInvalidScopes builds a chatter service echoer endpoint
// invalid-scopes error.
func NewEchoerInvalidScopes(body EchoerInvalidScopesResponseBody) chattersvc.InvalidScopes {
	v := chattersvc.InvalidScopes(body)
	return v
}

// NewEchoerUnauthorized builds a chatter service echoer endpoint unauthorized
// error.
func NewEchoerUnauthorized(body EchoerUnauthorizedResponseBody) chattersvc.Unauthorized {
	v := chattersvc.Unauthorized(body)
	return v
}

// NewListenerInvalidScopes builds a chatter service listener endpoint
// invalid-scopes error.
func NewListenerInvalidScopes(body ListenerInvalidScopesResponseBody) chattersvc.InvalidScopes {
	v := chattersvc.InvalidScopes(body)
	return v
}

// NewListenerUnauthorized builds a chatter service listener endpoint
// unauthorized error.
func NewListenerUnauthorized(body ListenerUnauthorizedResponseBody) chattersvc.Unauthorized {
	v := chattersvc.Unauthorized(body)
	return v
}

// NewSummaryChatSummaryCollectionOK builds a "chatter" service "summary"
// endpoint result from a HTTP "OK" response.
func NewSummaryChatSummaryCollectionOK(body SummaryResponseBody) chattersvcviews.ChatSummaryCollectionView {
	v := make([]*chattersvcviews.ChatSummaryView, len(body))
	for i, val := range body {
		v[i] = &chattersvcviews.ChatSummaryView{
			Message: val.Message,
			Length:  val.Length,
			SentAt:  val.SentAt,
		}
	}
	return v
}

// NewSummaryInvalidScopes builds a chatter service summary endpoint
// invalid-scopes error.
func NewSummaryInvalidScopes(body SummaryInvalidScopesResponseBody) chattersvc.InvalidScopes {
	v := chattersvc.InvalidScopes(body)
	return v
}

// NewSummaryUnauthorized builds a chatter service summary endpoint
// unauthorized error.
func NewSummaryUnauthorized(body SummaryUnauthorizedResponseBody) chattersvc.Unauthorized {
	v := chattersvc.Unauthorized(body)
	return v
}

// NewHistoryChatSummaryOK builds a "chatter" service "history" endpoint result
// from a HTTP "OK" response.
func NewHistoryChatSummaryOK(body *HistoryResponseBody) *chattersvcviews.ChatSummaryView {
	v := &chattersvcviews.ChatSummaryView{
		Message: body.Message,
		Length:  body.Length,
		SentAt:  body.SentAt,
	}
	return v
}

// NewHistoryInvalidScopes builds a chatter service history endpoint
// invalid-scopes error.
func NewHistoryInvalidScopes(body HistoryInvalidScopesResponseBody) chattersvc.InvalidScopes {
	v := chattersvc.InvalidScopes(body)
	return v
}

// NewHistoryUnauthorized builds a chatter service history endpoint
// unauthorized error.
func NewHistoryUnauthorized(body HistoryUnauthorizedResponseBody) chattersvc.Unauthorized {
	v := chattersvc.Unauthorized(body)
	return v
}

// Validate runs the validations defined on ChatSummaryResponseBody.
func (body *ChatSummaryResponseBody) Validate() (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.SentAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.sent_at", *body.SentAt, goa.FormatDateTime))
	}
	return
}
