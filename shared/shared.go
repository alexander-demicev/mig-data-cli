package shared

import (
	"fmt"
	"os/user"
)

// RepoURL https://github.com/fusor/ocp-mig-test-data.git
const RepoURL = "https://github.com/fusor/ocp-mig-test-data.git"

// MigTestDataDir return directrory for cloning mig-test-data repo
func MigTestDataDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/.ocp-mig-test-data/", usr.HomeDir), nil
}
