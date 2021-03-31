package torrentfile

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	bencode "github.com/IncSW/go-bencode"
)