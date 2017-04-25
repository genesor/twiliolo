package twiliolo

import "strconv"

type RequestOption interface {
	GetValue() (string, string)
}

type Page int

func (o Page) GetValue() (string, string) {
	return "Page", strconv.Itoa(int(o))
}

type PageSize int

func (o PageSize) GetValue() (string, string) {
	return "PageSize", strconv.Itoa(int(o))
}

type Beta bool

func (o Beta) GetValue() (string, string) {
	return "Beta", strconv.FormatBool(bool(o))
}

type AreaCode string

func (o AreaCode) GetValue() (string, string) {
	return "AreaCode", string(o)
}

type Contains int

func (o Contains) GetValue() (string, string) {
	return "Contains", string(o)
}

type SMSEnabled bool

func (o SMSEnabled) GetValue() (string, string) {
	return "SmsEnabled", strconv.FormatBool(bool(o))
}

type MMSEnabled bool

func (o MMSEnabled) GetValue() (string, string) {
	return "MmsEnabled", strconv.FormatBool(bool(o))
}

type VoiceEnabled bool

func (o VoiceEnabled) GetValue() (string, string) {
	return "VoiceEnabled", strconv.FormatBool(bool(o))
}

type FaxEnabled bool

func (o FaxEnabled) GetValue() (string, string) {
	return "FaxEnabled", strconv.FormatBool(bool(o))
}

type ExcludeAllAddressRequired bool

func (o ExcludeAllAddressRequired) GetValue() (string, string) {
	return "ExcludeAllAddressRequired", strconv.FormatBool(bool(o))
}

type ExcludeForeignAddressRequired bool

func (o ExcludeForeignAddressRequired) GetValue() (string, string) {
	return "ExcludeForeignAddressRequired	", strconv.FormatBool(bool(o))
}

type ExcludeLocalAddressRequired bool

func (o ExcludeLocalAddressRequired) GetValue() (string, string) {
	return "ExcludeLocalAddressRequired", strconv.FormatBool(bool(o))
}
