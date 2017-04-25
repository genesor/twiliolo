package twiliolo

import "strconv"

type RequestOption interface {
	GetValue() (string, string)
}

type OptionPage int

func (o OptionPage) GetValue() (string, string) {
	return "Page", strconv.Itoa(int(o))
}

type OptionPageSize int

func (o OptionPageSize) GetValue() (string, string) {
	return "PageSize", strconv.Itoa(int(o))
}

type OptionBeta bool

func (o OptionBeta) GetValue() (string, string) {
	return "Beta", strconv.FormatBool(bool(o))
}

type OptionAreaCode string

func (o OptionAreaCode) GetValue() (string, string) {
	return "AreaCode", string(o)
}

type OptionContains int

func (o OptionContains) GetValue() (string, string) {
	return "Contains", string(o)
}

type OptionSMSEnabled bool

func (o OptionSMSEnabled) GetValue() (string, string) {
	return "SmsEnabled", strconv.FormatBool(bool(o))
}

type OptionMMSEnabled bool

func (o OptionMMSEnabled) GetValue() (string, string) {
	return "MmsEnabled", strconv.FormatBool(bool(o))
}

type OptionVoiceEnabled bool

func (o OptionVoiceEnabled) GetValue() (string, string) {
	return "VoiceEnabled", strconv.FormatBool(bool(o))
}

type OptionFaxEnabled bool

func (o OptionFaxEnabled) GetValue() (string, string) {
	return "FaxEnabled", strconv.FormatBool(bool(o))
}

type OptionExcludeAllAddressRequired bool

func (o OptionExcludeAllAddressRequired) GetValue() (string, string) {
	return "ExcludeAllAddressRequired", strconv.FormatBool(bool(o))
}

type OptionExcludeForeignAddressRequired bool

func (o OptionExcludeForeignAddressRequired) GetValue() (string, string) {
	return "ExcludeForeignAddressRequired	", strconv.FormatBool(bool(o))
}

type OptionExcludeLocalAddressRequired bool

func (o OptionExcludeLocalAddressRequired) GetValue() (string, string) {
	return "ExcludeLocalAddressRequired", strconv.FormatBool(bool(o))
}
