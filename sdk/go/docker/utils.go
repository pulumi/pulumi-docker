// Copyright 2016-2020, Pulumi Corporation.

package docker

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

func checkRepositoryURL(repositoryURL string) error {
	_, tag := getImageNameAndTag(repositoryURL)

	// We want to report an advisory error to users so that they don't accidentally include a 'tag'
	// in the repo url they supply.  i.e. their repo url can be:
	//
	//      docker.mycompany.com/namespace/myimage
	//
	// but should not be:
	//
	//      docker.mycompany.com/namespace/myimage:latest
	//
	// We could consider removing this check entirely.  However, it is likely valuable to catch
	// clear mistakes where a tag was included in a repo url inappropriately.
	//
	// However, since we do have the check, we need to ensure that we do allow the user to specify
	// a *port* on their repository that the are communicating with.  i.e. it's fine to have:
	//
	//      docker.mycompany.com:5000 or
	//      docker.mycompany.com:5000/namespace/myimage
	//
	// So check if this actually does look like a port, and don't report an error in that case.
	//
	// From: https://www.w3.org/Addressing/URL/url-spec.txt
	//
	//      port        digits
	//
	// Regex = any number of digits, optionally followed by / and any remainder.
	validTag := regexp.MustCompile(`^\d+(\/.*)?`)
	if len(tag) > 0 && !validTag.MatchString(tag) {
		return errors.Errorf("%s should not contain a tag: %s", repositoryURL, tag)
	}

	return nil
}

func getImageNameAndTag(baseImageName string) (string, string) {
	// From https://docs.docker.com/engine/reference/commandline/tag
	//
	// "A tag name must be valid ASCII and may contain lowercase and uppercase letters, digits,
	// underscores, periods and dashes. A tag name may not start with a period or a dash and may
	// contain a maximum of 128 characters."
	//
	// So it is safe for us to just look for the colon, and consume whatever follows as the tag
	// for the image.
	lastColon := strings.LastIndex(baseImageName, ":")
	if lastColon < 0 {
		return baseImageName, ""
	}
	return baseImageName[:lastColon], baseImageName[lastColon+1:]
}

func localStageImageName(imageName string, stage string) string {
	return fmt.Sprintf("%s-%s", imageName, stage)
}
