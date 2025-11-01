package parsing

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"go.yaml.in/yaml/v4"
)

func ParseYml(filePath string) (map[string]any, error) {
	yamlData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var data map[string]any
	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type GhActionInput struct {
	Default     any
	Description any
	Required    any
}

type GhActionOutput struct {
	Description any
	Value       any
}

type GhAction struct {
	Name        string
	Description *string
	Inputs      map[string]GhActionInput
	Steps       []string
	Outputs     map[string]GhActionOutput
}

func (gha GhAction) StringifyNils() GhAction {
	g := gha
	if g.Description == nil {
		s := ""
		g.Description = &s
	}
	if len(g.Inputs) > 0 {
		for k, val := range g.Inputs {
			v := val
			if v.Description == nil {
				v.Description = "no description"
			}
			if v.Default == nil {
				v.Default = "no default"
			}
			if v.Required == nil {
				v.Required = false
			}
			g.Inputs[k] = v
		}
	}
	if len(g.Outputs) > 0 {
		for k, val := range g.Outputs {
			v := val
			if v.Description == nil {
				v.Description = "no description"
			}
			if v.Value == nil {
				v.Value = "no value"
			}
			g.Outputs[k] = v
		}
	}
	return g
}

func (g GhAction) ToMarkdownString() string {
	header := "# " + g.Name
	if g.Description != nil && *g.Description != "" {
		header = fmt.Sprintf("# %s\n\n## %s\n\n", g.Name, *g.Description)
	}
	if len(g.Inputs) > 0 {
		header += "### Inputs\n\n"
		header += "| Name | Description | Default | Required |\n"
		header += "|------|-------------|---------|----------|\n"
		sortedKeys := slices.Sorted(maps.Keys(g.Inputs))
		for _, k := range sortedKeys {
			v := g.Inputs[k]
			header += fmt.Sprintf("| `%s` | %v | `%v` | %v |\n", k, v.Description, v.Default, v.Required)
		}
	} else {
		header += "### Inputs\n\nNo inputs\n\n"
	}
	if len(g.Outputs) > 0 {
		header += "\n\n### Outputs\n\n"
		header += "| Name | Description | Value |\n"
		header += "|------|-------------|-------|\n"
		sortedKeys := slices.Sorted(maps.Keys(g.Outputs))
		for _, k := range sortedKeys {
			v := g.Outputs[k]
			header += fmt.Sprintf("| `%s` | %v | `%v` |\n", k, v.Description, v.Value)
		}
	} else {
		header += "\n\n### Outputs\n\nNo outputs"
	}

	if len(g.Steps) > 0 {
		header += "\n\n## Steps\n\n- "
		header += strings.Join(g.Steps, "\n- ")
	} else {
		header += "\n\n### Steps\n\nThis action has no reported steps <!-- presumably a JavaScript action -->"
	}
	return header
}

func ParseActionData(data map[string]any) (string, error) {
	action := GhAction{Inputs: map[string]GhActionInput{}, Outputs: map[string]GhActionOutput{}, Steps: []string{}}
	name, ok := data["name"]
	if !ok {
		return "", errors.New("the action must have a name")
	}
	typedName, okType := name.(string)
	if !okType {
		return "", errors.New("action name is not a string")
	}
	action.Name = typedName
	description, ok := data["description"]
	if ok {
		typedDescription, okType := description.(string)
		if !okType {
			return "", errors.New("action description is not a string")
		}
		action.Description = &typedDescription
	}
	inputs, ok := data["inputs"]
	if ok {
		typedInputs, okType := inputs.(map[string]any)
		if !okType {
			return "", errors.New("action inputs are not correctly structured")
		}
		for k := range typedInputs {
			typedMap, okType := typedInputs[k].(map[string]any)
			if !okType {
				return "", errors.New("action inputs are not correctly structured (malformed YAML)")
			}
			inpt := GhActionInput{}
			defaultV, ok := typedMap["default"]
			if ok {
				inpt.Default = defaultV
			}
			desc, ok := typedMap["description"]
			if ok {
				inpt.Description = desc
			}
			required, ok := typedMap["required"]
			if ok {
				inpt.Required = required
			}
			action.Inputs[k] = inpt
		}
	}
	outputs, ok := data["outputs"]
	if ok {
		typedOutputs, okType := outputs.(map[string]any)
		if !okType {
			return "", errors.New("action outputs are not correctly structured")
		}
		for k := range typedOutputs {
			typedMap, okType := typedOutputs[k].(map[string]any)
			if !okType {
				return "", errors.New("action inputs are not correctly structured (malformed YAML)")
			}
			outpt := GhActionOutput{}
			val, ok := typedMap["value"]
			if ok {
				outpt.Value = val
			}
			desc, ok := typedMap["description"]
			if ok {
				outpt.Description = desc
			}
			action.Outputs[k] = outpt
		}
	}
	run, ok := data["runs"]
	if !ok {
		return "", errors.New("action does not contain a run specification")
	}
	typedRun, okType := run.(map[string]any)
	if !okType {
		return "", errors.New("action run workflow is not correctly structured")
	}
	steps, ok := typedRun["steps"]
	if ok {
		stepsTyped, ok := steps.([]any)
		if !ok {
			return "", errors.New("action run workflow is not correctly structured (should have a list of steps)")
		}
		stepNames := []string{}
		for _, el := range stepsTyped {
			typedEl, okType := el.(map[string]any)
			if !okType {
				return "", errors.New("run step incorrectly structured (malformed YAML)")
			}
			stepName, ok := typedEl["name"]
			if !ok {
				return "", errors.New("each step should have a name")
			}
			stepNameTyped, okType := stepName.(string)
			if !okType {
				return "", errors.New("step name should be a string")
			}
			stepNames = append(stepNames, stepNameTyped)
		}
		action.Steps = stepNames
	}
	correctedAction := action.StringifyNils()
	return correctedAction.ToMarkdownString(), nil
}
