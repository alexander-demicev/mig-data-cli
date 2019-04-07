package ansible

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/alexander-demichev/mig-data-cli/pkg/shared"
)

// RunPlaybook run ansible playbook
func RunPlaybook(name string, withData, withBackup, withRestore bool) {
	migTestDataDir, err := shared.MigTestDataDir()
	if err != nil {
		log.Fatal(err)
	}

	nameArg := fmt.Sprintf("%s.yml", name)
	withDataArg := fmt.Sprintf("with_resources=%s", strconv.FormatBool(withData))
	withBackupArg := fmt.Sprintf("with_backup=%s", strconv.FormatBool(withBackup))
	withRestoreArg := fmt.Sprintf("with_restore=%s", strconv.FormatBool(withRestore))

	cmd := exec.Command("ansible-playbook", nameArg, "-e", withDataArg, "-e", withBackupArg, "-e", withRestoreArg)
	cmd.Dir = migTestDataDir
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(stdBuffer.String())
}
