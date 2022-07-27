package boilerplate

import (
	"fmt"
	"os"
	"os/exec"
)

func Brewer(pkgname string, appname string) (err error) {
	fmt.Println("Welcome Cuk !", pkgname)

	if err := createDirectories(appname); err != nil {
		return err
	}

	goExe, err := exec.LookPath("go")
	if err != nil {
		return err
	}

	if err := createFiles(pkgname, appname); err != nil {
		return err
	}

	cmdGo := &exec.Cmd{
		Path:   goExe,
		Args:   []string{goExe, "mod", "init"},
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

	if err := cmdGo.Run(); err != nil {
		return err
	}

	cmdGo = &exec.Cmd{
		Path:   goExe,
		Args:   []string{goExe, "mod", "tidy"},
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

	if err := cmdGo.Run(); err != nil {
		return err
	}

	return nil
}

func createDirectories(appname string) (err error) {
	if err := os.MkdirAll("cmd/"+appname, os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir("internal", os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir("pkg", os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll("configs/env", os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll("configs/db", os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll("configs/log", os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir("local", os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir("mocks", os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir("test", os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir("router", os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir("api", os.ModePerm); err != nil {
		return err
	}
	return nil
}
