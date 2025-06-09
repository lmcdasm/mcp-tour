package manager

import (
	"os"
	"path/filepath"
	"text/template"

	"mcp-go-server/models"
)

type ServerBuildContext struct {
	Definition models.McpServerDefinition
	Components models.McpServerComponents
	Status     models.McpServerStatus
}

var mainTemplate = `package main

import (
        "context"
        "mcp-go-server/mcplib"
)

func main() {
        server := mcplib.NewServer("{{.Definition.ID}}", "{{.Definition.Version}}", nil)

        {{range .Components.Tools}}
        server.AddTools(mcplib.NewTool("{{.Name}}", "{{.Description}}", {{.FuncName}}))
        {{end}}

        _ = server.Run(context.Background(), mcplib.New{{.Definition.Transport}}Transport())
}
`

type BuildTemplateContext struct {
	ID        string
	Version   string
	Transport string
	Tools     []models.ServerTool
}

func writeMainGo(ctx *ServerBuildContext, workspace string) error {
        dir := filepath.Join(workspace, ctx.Definition.ID)
        if err := os.MkdirAll(dir, 0755); err != nil {
                return err
        }

        f, err := os.Create(filepath.Join(dir, "main.go"))
        if err != nil {
                return err
        }
        defer f.Close()

        tmpl := template.Must(template.New("main").Parse(mainTemplate))
        return tmpl.Execute(f, ctx)
}

