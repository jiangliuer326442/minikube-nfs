package nfsd

import "os/exec"

func command(command string) error {
	cmd := exec.Command("/sbin/nfsd", command)

	return cmd.Run()
}

// Enable enables the nfsd service.
func Enable() error {
	return command("enable")
}

// Disable disables the nfsd service.
func Disable() error {
	return command("disable")
}

// Start starts the nfsd service.
func Start() error {
	return command("start")
}

// Stop stops the nfsd service.
func Stop() error {
	return command("stop")
}

// Restart restarts the nfsd service.
func Restart() error {
	return command("restart")
}

// Update sends a SIGHUP to the running nfsd daemon to cause it to update its configuration.
func Update() error {
	return command("update")
}

// CheckExports checks the exports file and reports any errors.
func CheckExports() error {
	return command("checkexports")
}
