package main

import (
    "syscall"
    "os"
    "log"
    "strconv"
    "bufio"
    "strings"
)

var uid, gid int

func chrootuid(dir, user string) {
    pw_filename := "/etc/passwd"
    pwf, err := os.Open(pw_filename, os.O_RDONLY, 0)
    if err != nil {
        log.Fatalf("%%Can't open %s: %s", pw_filename, err)
    }
    pwr := bufio.NewReader(pwf)
    for {
        line, err := pwr.ReadString('\n')
        if err != nil {
            log.Fatalf("%%Can't find UID for %s: %s", user, err)
        }
        pw_row := strings.Split(line, ":", 5)
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

    en := syscall.Chroot(dir)
    if en != 0 {
        log.Fatalln("%Chroot error:", os.Errno(en))
    }

    en = syscall.Setgid(gid)
    if en != 0 {
        log.Fatalln("%Setgid error:", os.Errno(en))
    }

    en = syscall.Setuid(uid)
    if en != 0 {
        log.Fatalln("%Setuid error:", os.Errno(en))
    }

    err = os.Chdir("/")
    if err != nil {
        log.Fatalln("%Can't cd to '/':", err)
    }
}
