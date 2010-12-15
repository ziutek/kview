package main

import (
    "syscall"
    "os"
    "log"
    "strconv"
)

func chrootuid() {
    if len(os.Args) != 3 {
        log.Exitf("Usage: %s [DIRECTORY UID]\n", os.Args[0])
    }

    uid, err := strconv.Atoi(os.Args[2])
    if err != nil {
        log.Exitln("Wrong UID:", err)
    }

    en := syscall.Chroot(os.Args[1])
    if en != 0 {
        log.Exitln("Chroot error:", os.Errno(en))
    }

    en = syscall.Setgid(uid)
    if en != 0 {
        log.Exitln("Setgid error:", os.Errno(en))
    }

    en = syscall.Setuid(uid)
    if en != 0 {
        log.Exitln("Setuid error:", os.Errno(en))
    }

    err = os.Chdir("/")
    if err != nil {
        log.Exitln(err)
    }
}
