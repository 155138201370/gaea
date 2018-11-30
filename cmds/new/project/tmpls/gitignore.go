package tmpls

const gitignoreTmpl = `
*{{.projectName|lName}}
.idea
.vscode
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, build with go test -c
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out
`
