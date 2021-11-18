package twui

import (
	"embed"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

//go:embed *.html
var templates embed.FS

// IncludeTemplates will add all *.html template files of twui.
// Make sure to call `tpl = tpl.Funcs(kyoto.TFuncMap())` first.
func IncludeTemplates(tpl *template.Template) (*template.Template, error) {
	return tpl.ParseFS(templates, "*.html")
}

func ExtractTemplates(targetPath string) error {
	files, err := templates.ReadDir("/")
	if err != nil {
		return err
	}
	for _, filex := range files {
		filex.Name()
		f, err := templates.Open(filex.Name())
		if err != nil {
			return err
		}
		_, fnm := filepath.Split(filex.Name())
		f2, err := os.Create(filepath.Join(targetPath, fnm))
		if err != nil {
			f.Close()
			return err
		}
		if _, err := io.Copy(f2, f); err != nil {
			f.Close()
			f2.Close()
			return err
		}
		f.Close()
		f2.Close()
	}
	return nil
}
