package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

func main() {
	j, err := exec.Command("pulumi", "plugin", "ls", "--json").CombinedOutput()
	contract.AssertNoErrorf(err, "`pulumi plugin ls --json` failed")
	type info struct {
		Name    string `json:"name"`
		Kind    string `json:"kind"`
		Version string `json:"version"`
	}

	var infos []info
	fmt.Println(string(j))
	err = json.Unmarshal(j, &infos)
	contract.AssertNoErrorf(err, "`pulumi plugin ls --json` output parsing failed")
}
