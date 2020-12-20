package boilingcore

// Custom defines custom parameters for the generation run
type Custom struct {
	Tables map[string]CustomGroup `toml:"tables,omitempty" json:"tables,omitempty"`
}

type CustomGroup struct {
	Groups map[string]CustomParameters `toml:"groups,omitempty" json:"groups,omitempty"`
}

type CustomParameters struct {
	Params map[string]string `toml:"params,omitempty" json:"params,omitempty"`
}

func (a Custom) Table(table string) CustomGroup {
	t, ok := a.Tables[table]
	if !ok {
		return CustomGroup{}
	}

	return t
}

func (t CustomGroup) Group(group string) CustomParameters {
	c, ok := t.Groups[group]
	if !ok {
		return CustomParameters{}
	}

	return c
}

func (t CustomParameters) Parameter(param string) string {
	c, ok := t.Params[param]
	if !ok {
		return ""
	}

	return c
}