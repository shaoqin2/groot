/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package security

import (
	//"html"

	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kennygrant/sanitize"
)

func SanitizeRequest(r *http.Request) {

	if !enabled {
		return
	}

	content, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	sanitizedHTML := sanitize.HTML(string(content))
	r.Body = ioutil.NopCloser(strings.NewReader(sanitizedHTML))
	r.ContentLength = int64(len(sanitizedHTML))

}
