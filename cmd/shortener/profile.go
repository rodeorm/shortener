package main

import (
	"os"
	"runtime"
	"runtime/pprof"
)

func profile(profileType int) error {
	if profileType != noneProfile {
		var (
			fmem *os.File
			err  error
		)

		if profileType == baseProfile {
			fmem, err = os.Create(`base.pprof`)
		} else {
			fmem, err = os.Create(`result.pprof`)
		}
		if err != nil {
			return err
		}
		defer fmem.Close()

		runtime.GC()
		if err := pprof.WriteHeapProfile(fmem); err != nil {
			return err
		}
	}
	return nil
}
