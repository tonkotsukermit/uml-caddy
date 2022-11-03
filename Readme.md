# UML-Caddy

Uml-caddy is a tool for templating plantuml from known metadata

## Running

The docker-compose file allows you to run standard docker commands to run the container locally for debugging

`docker-compose up --build`

## Development

### Recommended Installed Tooling

- Golang v1.19+
- Docker
- Docker-compose
- PlantUML

### vscode

Example `launch.json`
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/uml-caddy",
            "cwd": "${workspaceFolder}"
        }
    ]
}
```
----
### Adding new models (Pre-DB implementation)

- (Optional) add a new `./pkg/models.models_<system>.go` file in the `models` package
- Add a UML template utilizing the golang templating methodology. [Sprig functions](http://masterminds.github.io/sprig/) are also supported. 
    - Template example: 
``` golang
const CustomUMLBase = `
@startuml {{ .Name | default "UML Output" }}

{{ if .Header}}
  header {{ .Header }}
{{ end }}

{{ if .Title }}
  title {{ .Title }}
{{ end }}

{{ partial "Custom" .K8s }}

@enduml
`
```

- Add all necessary nested partials as constants, and reference them with the `{{ partial "PartialModelName" .Object }}` syntax
- Add all constants to the model index at `./pkg/models/models_index.go` in the `models` package
    - Models index example:
```golang
models := map[string]string{
    "K8sUMLVirtualBase":  K8sUMLVirtualBase,
    "K8sVirtual":         K8sVirtual,
    "NamespaceModel":     NamespaceModel,
    "DeploymentModel":    DeploymentModel,
    "PodModel":           PodModel,
    "CustomUMLBase":      CustomUMLBase,
}
```
----
----
### Using models 

- (Optional) Extend the `UML` struct for a custom data object
    - Example struct extension:
```golang
type CustomUML struct {
	UML
	Data	importer.CustomResources //this assumes a new importer package associated with the custom resource
}

```

- Add a new generator function in the `gen` package
    - Example Generator function:
```golang
func (c *CustomUML) GenerateCustomUML() error{

data := c.Data.Get()

t := models.Template{
    Model:  models.CustomUMLBase,
    Output: c.Output,
}

return t.Execute(data)
```
- Import/Utilize the package in main 
    - Example main package function
    ``` golang
    func generateCustomPUML(w io.Writer){

        u := gen.CustomUML{
            UML: gen.UML{
                Name:         "CustomName",
                Header:       "CustomHeader",
                Title:        "CustomTitle",
                Output:       w,
            },
        }

        err := u.GenerateCustomUML()
        if err != nil {
            w.Write([]byte(err.Error()))
            return
        }

    }
    ```