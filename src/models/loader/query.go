package loader

type ModelCondition struct {
}

// {{ .Name }} returns a {{ .SnakeName }}.ConditionStart for {{ .Name }}Model
func {{ .Name }}() {{ .SnakeName }}.ConditionStart {
return {{ .SnakeName }}.ConditionStart{
ConditionStart: &models.ConditionStart{},
}
}
