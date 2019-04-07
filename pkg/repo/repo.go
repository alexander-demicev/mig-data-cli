package repo

import (
	"fmt"
	"os/exec"

	"github.com/alexander-demichev/mig-data-cli/pkg/log"
	"github.com/alexander-demichev/mig-data-cli/pkg/shared"
)

// CloneRepo clone mig test data dir
func CloneRepo() {
	migTestDataDir, err := shared.MigTestDataDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Create directory for ocp-mig-test-data playbooks")
	out, err := exec.Command("mkdir", "-p", migTestDataDir).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Debugln(out)

	fmt.Println("Started cloning ocp-mig-test-data playbooks")
	cmd := exec.Command("git", "clone", shared.RepoURL, ".")
	cmd.Dir = migTestDataDir
	out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Debugln(out)
	fmt.Println("Finished cloning ocp-mig-test-data playbooks")
}

// UpdateRepo Update ocp-mig-test-data repo
func UpdateRepo() {
	migTestDataDir, err := shared.MigTestDataDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updating ocp-mig-test-data repo")
	cmd := exec.Command("git", "pull", "origin")
	cmd.Dir = migTestDataDir
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Debugln(out)
}
