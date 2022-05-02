package etrisfpocctnmgmt

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var description = `
ctnmgmt module is container management module for etri-poc
It is not yet ready for deployment
`

func PrintModuleDescription() {
	log.Println(description)
}

func IsExist(name string) bool {
	cmd := strings.Split("container\\ls\\--format\\'{{.Image}} {{.Names}}'\\-a", "\\")
	bout, err := exec.Command("docker", cmd...).Output()
	if err != nil {
		log.Fatalln(err)
	}

	sout := strings.Split(string(bout), "\n")

	for _, e := range sout {
		l := strings.Split(e, " ")

		if len(l) < 2 {
			continue
		}

		if name == l[0] {
			return true
		}
	}

	return false
}

func CreateContainer(name string) error {

	if IsExist(name) {
		return nil
	}
	args := strings.Split(fmt.Sprintf("container\\run\\-d\\%s", name), "\\")
	fmt.Println(args)
	_, err := exec.Command("docker", args...).Output()
	if err != nil {
		return err
	}

	return nil
}
