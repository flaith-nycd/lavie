/**
 * https://github.com/riobard/go-virtualbox/blob/master/vbm.go
 */
package vmbox

import (
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
)

var VBoxManagePath string // Path to VBoxManage utility.

/**
 * The init function
 * (https://golang.org/doc/effective_go.html#init)
 *
 * Finally, each source file can define its own niladic init function to set up whatever state is required.
 * (Actually each file can have multiple init functions.)
 * And finally means finally: init is called after all the variable declarations in the package have evaluated their initializers,
 * and those are evaluated only after all the imported packages have been initialized.
 *
 * Besides initializations that cannot be expressed as declarations, a common use of init functions is to verify
 * or repair correctness of the program state before real execution begins.
 */
func init() {
    VBoxManagePath = ""
    //if p := os.Getenv("VBOX_INSTALL_PATH"); p != "" && runtime.GOOS == "windows" {
    if p := os.Getenv("VBOX_MSI_INSTALL_PATH"); p != "" && runtime.GOOS == "windows" {
        VBoxManagePath = filepath.Join(p, "VBoxManage.exe")
    }
}

func vbm(args ...string) error {
    cmd := exec.Command(VBoxManagePath, args...)

    if err := cmd.Run(); err != nil {
        return err
    }
    return nil
}

func vbmOut(args ...string) (string, error) {
    cmd := exec.Command(VBoxManagePath, args...)

    b, err := cmd.Output()
    return string(b), err
}
