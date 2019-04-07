package list

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/alexander-demichev/ocp-mig-test-data-cli/log"
	"github.com/alexander-demichev/ocp-mig-test-data-cli/shared"
)

// List list available samples
func List() []string {
	migTestDataDir, err := shared.MigTestDataDir()
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("ls")
	cmd.Dir = migTestDataDir
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out[:])

	splittedOutput := strings.Split(output, "\n")

	filtered := make([]string, 0)
	for _, file := range splittedOutput {
		if strings.HasSuffix(file, ".yml") && strings.Compare(file, "all.yml") != 0 {
			formatedFile := strings.TrimSuffix(file, ".yml")
			filtered = append(filtered, formatedFile)
		}
	}
	return filtered
}

// PrintList print playbook list
func PrintList(list []string) {
	fmt.Println("Available samples:")
	for _, file := range list {
		fmt.Println(file)
	}
}

// IsValidPlaybook true, if playbook is valid
func IsValidPlaybook(playbook string) bool {
	list := List()

	for _, file := range list {
		if file == playbook {
			return true
		}
	}
	return false
}
