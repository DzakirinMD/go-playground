package assets

import _ "embed"

//go:embed gpg/somekey.gpg
var GPGKey string

//go:embed gpg/uuid.txt
var GPGUUID string
