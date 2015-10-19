package stm

import (
	"net/url"
	"os"
	"path/filepath"
)

func NewLocation() *Location {
	loc := &Location{
		adapter:    NewFileAdapter(),
		publicPath: "public/",
	}
	return loc
}

type Location struct {
	adapter Adapter

	verbose      bool
	host         string
	publicPath   string
	sitemapsPath string
}

func (loc *Location) SetPublicPath(path string) {
	loc.publicPath = path
}

func (loc *Location) SetSitemapsPath(path string) {
	loc.sitemapsPath = path
}

// func (loc *Location) with(opts={})
// self.merge(opts)
// }

func (loc *Location) Directory() string {
	return filepath.Join(loc.publicPath, loc.sitemapsPath)
}

func (loc *Location) Path() string {
	return filepath.Join(loc.publicPath, loc.sitemapsPath, loc.Filename())
}

func (loc *Location) PathInPublic() string {
	return filepath.Join(loc.sitemapsPath, loc.Filename())
}

func (loc *Location) URL() string {
	base, _ := url.Parse(loc.host)

	var u *url.URL
	for _, ref := range []string{loc.sitemapsPath, loc.Filename()} {
		u, _ = url.Parse(ref)
		base.ResolveReference(u)
	}

	return base.String()
}

// Return the size of the file at
func (loc *Location) Filesize() int64 {
	f, _ := os.Open(loc.Path())
	defer f.Close()
	fi, _ := f.Stat()
	return fi.Size()
}

// Return the filename.  Raises an exception if no filename or namer is set.
// If using a namer once the filename has been retrieved from the namer its
// value is locked so that it is unaffected by further changes to the namer.
func (loc *Location) Filename() string {
	return ""

	// raise SitemapGenerator::SitemapError, "No filename or namer set" unless self[:filename] || self[:namer]
	// unless self[:filename]
	// self.send(:[]=, :filename, self[:namer].to_s, :super => true)

	// // Post-process the filename for our compression settings.
	// // Strip the `.gz` from the extension if we aren't compressing this file.
	// // If you're setting the filename manually, :all_but_first won't work as
	// // expected.  Ultimately I should force using a namer in all circumstances.
	// // Changing the filename here will affect how the FileAdapter writes out the file.
	// if self[:compress] == false || (self[:namer] && self[:namer].start? && self[:compress] == :all_but_first) {
	// self[:filename].gsub!(/\.gz$/, '')
	// }
	// self[:filename]
}

// If a namer is set, reserve the filename and increment the namer.
// Returns the reserved name.
// func (loc *Location) ReserveName() {
// if self[:namer]
// filename
// self[:namer].next
// end
// self[:filename]
// }

// Return true if this location has a fixed filename.  If no name has been
// reserved from the namer, for instance, returns false.
// func (loc *Location) IsReservedName() bool {
// !!self[:filename]
// }

// func (loc *Location) namer() {
// self[:namer]
// }

func (loc *Location) IsVerbose() bool {
	return loc.verbose
}

// If you set the filename, clear the namer and vice versa.
// func (loc *Location) []=(key, value, opts={})
// if !opts[:super]
// case key
// when :namer
// super(:filename, nil)
// when :filename
// super(:namer, nil)
// end
// end
// super(key, value)
// }

// Write `data` out to a file.
// Output a summary line if verbose is true.
func (loc *Location) Write(data []byte, linkCount int) {
	loc.adapter.Write(loc, data)
	if loc.IsVerbose() {
		pp.Println(loc.Summary(linkCount))
	}
}

// Return a summary string
func (loc *Location) summary(linkCount int) string {
	// filesize = number_to_human_size(loc.Filesize())
	// width = self.class::PATH_OUTPUT_WIDTH
	// path = SitemapGenerator::Utilities.ellipsis(self.path_in_public, width)
	// fmt.Sprintf("+ #{('%-'+width.to_s+'s') % path} #{'%10s' % link_count} links / #{'%10s' % filesize}")
	return ""
}
