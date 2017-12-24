package api

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

//Version according Semantic Versioning 2.0.0
type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"release"`
}

//NewVersion constructor from string
func NewVersion(version string) (*Version, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return nil, errors.New("Version: bad string version format")
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, errors.New("Version: bad version format ")
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errors.New("Version: bad version format")
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, errors.New("Version: bad version format")
	}

	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}

//GetVersion return version
func (v *Version) GetVersion() *Version {
	return v
}

//ToString return version
func (v *Version) ToString() string {
	version := ""

	if v != nil {
		major := strconv.Itoa(v.Major)
		minor := strconv.Itoa(v.Minor)
		patch := strconv.Itoa(v.Patch)

		version = major + "." + minor + "." + patch
	}

	return version
}

//ToJSON return version
func (v *Version) ToJSON() string {
	json, _ := json.Marshal(v)
	return string(json)
}
