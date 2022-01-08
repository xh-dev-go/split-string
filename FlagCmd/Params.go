package FlagCmd

import "flag"

type StringParam struct {
	value *string
	name string
	defaultValue string
	usage string
}

func (param *StringParam) Value() string{
	return *param.value
}

func (param *StringParam) Bind(flag *flag.FlagSet) *StringParam{
	flag.StringVar(param.value, param.name, param.defaultValue, param.usage)
	return param
}

func NewStringParamDefault(name, usage string) *StringParam{
	defaultValue := ""
	v := StringParam{value: &defaultValue}
	v.name = name
	v.usage = usage
	return &v
}
func NewStringParam(name, defaultValue, usage string) *StringParam{
	v := NewStringParamDefault(name, usage)
	v.defaultValue = defaultValue
	return v
}
