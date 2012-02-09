package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
)

var uid, gid int

func chrootuid(dir, user string) {
	pw_filename := "/etc/passwd"
	pwf, err := os.Open(pw_filename)
	if err != nil {
		log.Fatalf("%%Can't open %s: %s", pw_filename, err)
	}
	pwr := bufio.NewReader(pwf)
	for {
		line, err := pwr.ReadString('\n')
		if err != nil {
			log.Fatalf("%%Can't find UID for %s: %s", user, err)
		}
		pw_row := strings.SplitN(line, ":", 5)
		if len(pw_row) != 5 {
			continue
		}
		if pw_row[0] == user {
			uid, err = strconv.Atoi(pw_row[2])
			if err != nil {
				log.Fatalln("%Wrong UID:", err)
			}
			gid, err = strconv.Atoi(pw_row[3])
			if err != nil {
				log.Fatalln("%Wrong GID:", err)
			}
			break
		}
	}

	err = syscall.Chroot(dir)
	if err != nil {
		log.Fatalln("%Chroot error:", err)
	}

	err = syscall.Setgid(gid)
	if err != nil {
		log.Fatalln("%Setgid error:", err)
	}

	err = syscall.Setuid(uid)
	if err != nil {
		log.Fatalln("%Setuid error:", err)
	}

	err = os.Chdir("/")
	if err != nil {
		log.Fatalln("%Can't cd to '/':", err)
	}
}
