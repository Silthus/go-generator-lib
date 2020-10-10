package api

// Specifies what templates belong to a generator and what variables it needs to run.
//
// Will be read from a generator-*.yaml file in the root directory of the generator.
//
// The values of the variables as well as what generator to use come from a RenderSpec instead.
type GeneratorSpec struct {
	// The list of templates to render (if their condition evaluates to true)
	Templates []TemplateSpec `yaml:"templates"`

	// The list of available variables
	Variables []VariableSpec `yaml:"variables"`
}

// Specifies a template to process, or a list to iterate over, if WithItems is nonempty (setting {{ item }} each run)
//
// Every field is evaluated as a template itself, so you can use variables in all fields.
type TemplateSpec struct {
	RelativeSourcePath string        `yaml:"source"`
	RelativeTargetPath string        `yaml:"target"`
	Condition          string        `yaml:"condition"`
	WithItems          []interface{} `yaml:"with_items"`
}

// Specifies a variable that this generator uses, so it is made available in the templates.
//
// Actual values for an invocation of the generator are set in a RenderSpec, not the GeneratorSpec.
type VariableSpec struct {
	// Human readable description for the variable.
	Description string `yaml:"description"`

	// Regex validation pattern that the value must match.
	ValidationPattern string `yaml:"pattern"`

	// Default value. If left empty, the variable is considered required.
	DefaultValue string `yaml:"default"`
}
