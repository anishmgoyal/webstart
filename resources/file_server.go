package resources

import (
	"net/http"
	"os"
)

// TODO: Implement caching for these static files, possibly some
//       "compilation" to reduce file size?

// NoListingFileSystem wraps a file system in which directory
// listings are disabled
type NoListingFileSystem struct {
	fs http.FileSystem
}

// Open opens a file
func (wrapper NoListingFileSystem) Open(name string) (http.File, error) {
	f, err := wrapper.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return NoListingReadDir{f}, nil
}

// NoListingReadDir wraps a file preventing access to listings of a directory
type NoListingReadDir struct {
	http.File
}

// Readdir is a generic wrapper function that just returns nil
func (f NoListingReadDir) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

// MapCSSHandler creates the /css route for getting CSS files
func MapCSSHandler() {
	fs := NoListingFileSystem{http.Dir("webapp/css")}
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(fs)))
}

// MapJSHandler creates the /js route for getting JS files
func MapJSHandler() {
	fs := NoListingFileSystem{http.Dir("webapp/js")}
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(fs)))
}

// MapImageHandler creates the /img route for getting image files
func MapImageHandler() {
	fs := NoListingFileSystem{http.Dir("webapp/img")}
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(fs)))
}
