package option

import "strconv"

// RequestOption is the interface implemented by each querystring parameter used by Twilio API
type RequestOption interface {
	GetValue() (string, string)
}

// Page type for querystring parameter
type Page int

// GetValue returns the query string compliant name and value
func (o Page) GetValue() (string, string) {
	return "Page", strconv.Itoa(int(o))
}

// PageSize type for querystring parameter
type PageSize int

// GetValue returns the query string compliant name and value
func (o PageSize) GetValue() (string, string) {
	return "PageSize", strconv.Itoa(int(o))
}

// Beta type for querystring parameter
type Beta bool

// GetValue returns the query string compliant name and value
func (o Beta) GetValue() (string, string) {
	return "Beta", strconv.FormatBool(bool(o))
}

// AreaCode type for querystring parameter
type AreaCode string

// GetValue returns the query string compliant name and value
func (o AreaCode) GetValue() (string, string) {
	return "AreaCode", string(o)
}

// Contains type for querystring parameter
type Contains string

// GetValue returns the query string compliant name and value
func (o Contains) GetValue() (string, string) {
	return "Contains", string(o)
}

// SMSEnabled type for querystring parameter
type SMSEnabled bool

// GetValue returns the query string compliant name and value
func (o SMSEnabled) GetValue() (string, string) {
	return "SmsEnabled", strconv.FormatBool(bool(o))
}

// MMSEnabled type for querystring parameter
type MMSEnabled bool

// GetValue returns the query string compliant name and value
func (o MMSEnabled) GetValue() (string, string) {
	return "MmsEnabled", strconv.FormatBool(bool(o))
}

// VoiceEnabled type for querystring parameter
type VoiceEnabled bool

// GetValue returns the query string compliant name and value
func (o VoiceEnabled) GetValue() (string, string) {
	return "VoiceEnabled", strconv.FormatBool(bool(o))
}

// FaxEnabled type for querystring parameter
type FaxEnabled bool

// GetValue returns the query string compliant name and value
func (o FaxEnabled) GetValue() (string, string) {
	return "FaxEnabled", strconv.FormatBool(bool(o))
}

// ExcludeAllAddressRequired type for querystring parameter
type ExcludeAllAddressRequired bool

// GetValue returns the query string compliant name and value
func (o ExcludeAllAddressRequired) GetValue() (string, string) {
	return "ExcludeAllAddressRequired", strconv.FormatBool(bool(o))
}

// ExcludeForeignAddressRequired type for querystring parameter
type ExcludeForeignAddressRequired bool

// GetValue returns the query string compliant name and value
func (o ExcludeForeignAddressRequired) GetValue() (string, string) {
	return "ExcludeForeignAddressRequired	", strconv.FormatBool(bool(o))
}

// ExcludeLocalAddressRequired type for querystring parameter
type ExcludeLocalAddressRequired bool

// GetValue returns the query string compliant name and value
func (o ExcludeLocalAddressRequired) GetValue() (string, string) {
	return "ExcludeLocalAddressRequired", strconv.FormatBool(bool(o))
}
