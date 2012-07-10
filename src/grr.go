package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

var verbose *bool = flag.Bool("v", false, "Be (very) verbose")
var jump *int64 = flag.Int64("j", 32, "Random jump maximum size")
var really *bool = flag.Bool("R", false, "Really do grr")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: grr [OPTS] PATH...\nPATH Is one or more filesystem paths\nOPTS Is one or more of the following:\n")
		flag.PrintDefaults()
	}
}

func main() {

	flag.Parse()

	targets := flag.Args()

	if len(targets) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	for i := range targets {
		target := targets[i]
		if *verbose {
			log.Printf("Processing target %v", target)
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
							for ; pos < info.Size(); pos = pos + rand.Int63n(*jump) + 1 {
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
