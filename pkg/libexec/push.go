/*
  Copyright (c) 2018, Sylabs, Inc. All rights reserved.

  This software is licensed under a 3-clause BSD license.  Please
  consult LICENSE file distributed with the sources of this project regarding
  your rights to use or distribute this software.
*/
package libexec

import (
	"fmt"

	"github.com/singularityware/singularity/pkg/library/client"
	"log"
)

func PushImage(image string, library string, libraryURL string) {
	fmt.Printf("Pushing image: \"%s\" to library: \"%s\"\n", image, library)
	err := client.UploadImage(image, library, libraryURL)
	if err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	}
}