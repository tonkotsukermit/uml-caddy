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