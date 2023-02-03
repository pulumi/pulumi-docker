package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"gopkg.in/yaml.v3"
)

//go:generate go run generate.go yaml .

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stdout, "Usage: %s <yaml source dir path> <markdown destination path>\n", os.Args[0])
		os.Exit(1)
	}
	yamlPath := os.Args[1]
	mdPath := os.Args[2]

	if !filepath.IsAbs(yamlPath) {
		cwd, err := os.Getwd()
		contract.AssertNoError(err)
		yamlPath = filepath.Join(cwd, yamlPath)
	}

	fileInfo, err := os.Lstat(mdPath)
	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(mdPath, 0600); err != nil {
			panic(err)
		}
	}

	if !fileInfo.IsDir() {
		fmt.Fprintf(os.Stderr, "Expect markdown destination %q to be a directory\n", mdPath)
		os.Exit(1)
	}

	yamlFiles, err := os.ReadDir(yamlPath)
	if err != nil {
		panic(err)
	}
	for _, yamlFile := range yamlFiles {
		if err := processYaml(filepath.Join(yamlPath, yamlFile.Name()), mdPath); err != nil {
			fmt.Fprintf(os.Stderr, "%+v", err)
			os.Exit(1)
		}
	}
}

func markdownExamples(examples []string) string {
	s := "{{% examples %}}\n## Example Usage\n"
	for _, example := range examples {
		s += example
	}
	s += "{{% /examples %}}\n"
	return s
}

func markdownExample(description string,
	typescript string,
	python string,
	csharp string,
	golang string,
	yaml string,
	java string) string {

	return fmt.Sprintf("{{%% example %%}}\n### %s\n\n"+
		"```typescript\n%s```\n"+
		"```python\n%s```\n"+
		"```csharp\n%s```\n"+
		"```go\n%s```\n"+
		"```yaml\n%s```\n"+
		"```java\n%s```\n"+
		"{{%% /example %%}}\n",
		description, typescript, python, csharp, golang, yaml, java)
}

func convert(language, tempDir, programFile string) (string, error) {
	exampleDir := filepath.Join(tempDir, "example"+language)
	cmd := exec.Command("pulumi", "convert", "--language", language, "--out", filepath.Join(tempDir, exampleDir), "--generate-only")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "convert %s failed, ignoring: %+v", language, err)
	}
	content, err := os.ReadFile(filepath.Join(tempDir, exampleDir, programFile))
	if err != nil {
		return "", err
	}
	return string(content), nil

}

func processYaml(path string, mdDir string) error {
	yamlFile, err := os.Open(path)
	if err != nil {
		return err
	}

	base := filepath.Base(path)
	md := strings.NewReplacer(".yaml", ".md", ".yml", ".md").Replace(base)

	defer contract.IgnoreClose(yamlFile)
	decoder := yaml.NewDecoder(yamlFile)
	exampleStrings := []string{}
	for {
		example := map[string]interface{}{}
		err := decoder.Decode(&example)
		if err == io.EOF {
			break
		}

		description := example["description"].(string)
		dir, err := os.MkdirTemp("", "")
		if err != nil {
			return err
		}

		defer func() {
			contract.IgnoreError(os.RemoveAll(dir))
		}()

		src, err := os.OpenFile(filepath.Join(dir, "Pulumi.yaml"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			return err
		}

		if err = yaml.NewEncoder(src).Encode(example); err != nil {
			return err
		}
		contract.AssertNoError(src.Close())

		typescript, err := convert("typescript", dir, "index.ts")
		python, err := convert("python", dir, "__main__.py")
		csharp, err := convert("csharp", dir, "Program.cs")
		golang, err := convert("go", dir, "main.go")
		java, err := convert("java", dir, "src/main/java/generated_program/App.java")

		yamlContent, err := os.ReadFile(filepath.Join(dir, "Pulumi.yaml"))
		if err != nil {
			return err
		}
		yaml := string(yamlContent)

		exampleStrings = append(exampleStrings, markdownExample(description, typescript, python, csharp, golang, yaml, java))
	}
	contract.AssertNoError(err)
	fmt.Fprintf(os.Stdout, "Writing %s\n", filepath.Join(mdDir, md))
	f, err := os.OpenFile(filepath.Join(mdDir, md), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer contract.IgnoreClose(f)
	_, err = f.Write([]byte(markdownExamples(exampleStrings)))
	contract.AssertNoError(err)
	return nil
}
