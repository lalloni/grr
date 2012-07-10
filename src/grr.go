package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

var verbose *bool = flag.Bool("v", false, "Be (very) verbose.")
var really *bool = flag.Bool("R", false, "Really grr.")

func main() {

	flag.Parse()

	targets := flag.Args()

	for i := range targets {
		target := targets[i]
		if *verbose {
			log.Printf("Processing target [%v]...", target)
		}
		filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
			if err == nil {
				if info.IsDir() {
					if *verbose {
						log.Printf("Ignoring dir %v", path)
					}
				} else {
					if *verbose {
						log.Printf("Trying to grr %v", path)
					}
					buf := make([]byte, 1)
					file, e := os.OpenFile(path, os.O_WRONLY, info.Mode())
					wrote := false
					if e != nil {
						log.Printf("Can't open %v: %v", path, e.Error())
					} else {
						pos, e := file.Seek(16, 0)
						if e != nil {
							log.Printf("Can't seek %v: %v", path, e.Error())
						} else {
							for ; pos+1 < info.Size(); pos = pos + rand.Int63n(16) {
								if *verbose {
									log.Printf("Seeking to %v", pos)
								}
								_, e := file.Seek(pos, 0)
								if e != nil {
									log.Printf("Can't seek %v: %v", path, e.Error())
								} else {
									buf[0] = byte(rand.Int())
									if *verbose {
										log.Printf("Writing %v at %v", buf[0], pos)
									}
									if *really {
										_, e := file.Write(buf)
										if e != nil {
											log.Printf("Can't write %v to %v at %v: %v", buf[0], path, pos, e.Error())
										} else {
											wrote = true
										}
									}
								}
							}
						}
						file.Close()
						if wrote {
							e := os.Chtimes(path, info.ModTime(), info.ModTime())
							if e != nil {
								log.Printf("Can't change %v times: %v", path, e.Error())
							}
						}
					}
				}
			}
			return nil
		})
	}

}
