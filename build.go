package main

import (
	"errors"

	"github.com/lunny/log"
)

type BuildStatus byte

const (
	Downloading BuildStatus = iota + 1 // downloading the source code
	Pending                            // pending to build
	Building                           // building
	Finished                           // build successfully
	Failed                             // build failed
)

type Build struct {
	Pkg         string // pkg name
	Version     string // pkg version, could be master, tag, or git id
	BuildStatus        // build status
	BuildEror   string // build error if build failed
}

var (
	ErrNotImplemented = errors.New("not implemented")
)

// save pkg's info to nodb
func savePkg(build *Build) error {
	return ErrNotImplemented
}

func isPkgExist(pkg, version string) (bool, error) {
	return false, nil
}

func download(pkg, version string) error {
	// check if pkg is exist
	exist, err := isPkgExist(pkg, version)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	// save pkg info
	err = savePkg(&Build{
		Pkg:         pkg,
		Version:     version,
		BuildStatus: Downloading,
	})
	if err != nil {
		return err
	}

	// download

	return ErrNotImplemented
}

// build the pkg and special version
func build(pkg, version string) error {
	log.Debug("start building", pkg, version)
	defer func() {
		log.Debug("end building", pkg, version)
	}()

	err := download(pkg, version)
	if err != nil {
		return err
	}
	return ErrNotImplemented
}
