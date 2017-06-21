package nfsd

import "os/exec"

func command(command string) error {
	cmd := exec.Command("/sbin/nfsd", command)

	return cmd.Run()
}

// Enables the nfsd service.
func Enable() error {
	return command("enable")
}

// Disables the nfsd service.
func Disable() error {
	return command("disable")
}

// Starts the nfsd service.
func Start() error {
	return command("start")
}

// Stops the nfsd service.
func Stop() error {
	return command("stop")
}

// Restarts the nfsd service.
func Restart() error {
	return command("restart")
}

// Sends a SIGHUP to the running nfsd daemon to cause it to update its configuration.
func Update() error {
	return command("update")
}

// Checks the exports file and reports any errors.
func CheckExports() error {
	return command("checkexports")
}
