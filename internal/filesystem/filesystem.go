package filesystem

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
)

// FileSystem ...
type FileSystem struct {
}

// GetPartial ...
func (f *FileSystem) GetPartial(partial string) string {
	return ""
}

// InitPublicFolder ...
func (f *FileSystem) InitPublicFolder() {
	f.removeContents(f.GetPublicPath())
	f.copyAssets()
}

func (f *FileSystem) copyAssets() {
	copy.Copy(f.GetAssetsPath(), f.GetPublicAssetsPath())
}

// GetPublicPath ...
func (f *FileSystem) GetPublicPath() string {
	return f.getPath("public")
}

// GetPublicAssetsPath ...
func (f *FileSystem) GetPublicAssetsPath() string {
	return f.GetPublicPath() + "/assets"
}

// GetAssetsPath ...
func (f *FileSystem) GetAssetsPath() string {
	return f.GetCurrentThemePath() + "/assets"
}

// GetDataPath ...
func (f *FileSystem) GetDataPath() string {
	return f.getPath("data")
}

// GatherAllContent scans and returns all Markdown source files, in form of
// array of filenames (with path)
func (f *FileSystem) GatherAllContent() ([]string, error) {
	return f.glob(f.GetDataPath(), ".md")
}

// GetThemesPath returns the path for the main themes folder
func (f *FileSystem) GetThemesPath() string {
	return f.getPath("themes")
}

// GetCurrentThemePath returns the path to the current theme folder
func (f *FileSystem) GetCurrentThemePath() string {
	return f.GetThemesPath() + "/" + f.GetCurrentTheme()
}

// GetCurrentTheme returns the name of the current theme (and theme folder)
func (f *FileSystem) GetCurrentTheme() string {
	return "superstrat"
}

// GatherAllLayouts returns a list of HTML templates in the theme folder
func (f *FileSystem) GatherAllLayouts() ([]string, error) {
	return f.glob(f.GetCurrentThemePath(), ".html")
}

// Create ...
func (f *FileSystem) Create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}

	return os.Create(p)
}

func (f *FileSystem) glob(dir string, ext string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func (f *FileSystem) getPath(relative string) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir + "/" + strings.Trim(relative, "/")
}

func (f *FileSystem) removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
