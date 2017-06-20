package nfsd

import "os/exec"

func command(command string) error {
	cmd := exec.Command("/sbin/nfsd", command)

	return cmd.Run()
}

func Enable() error {
	return command("enable")
}

func Disable() error {
	return command("disable")
}

func Start() error {
	return command("start")
}

func Stop() error {
	return command("stop")
}

func Restart() error {
	return command("restart")
}

func Update() error {
	return command("update")
}

func CheckExports() error {
	return command("checkexports")
}
