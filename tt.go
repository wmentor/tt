package tt

import (
	"bytes"
	"errors"
	"io"
	"net/url"

	"github.com/CloudyKit/jet"
)

type Vars = jet.VarMap

var (
	tmpl *jet.Set = newTT()
)

func newTT(dirs ...string) *jet.Set {
	nt := jet.NewHTMLSet(dirs...)
	regFilters(nt)
	return nt
}

func MakeVars() Vars {
	return make(Vars)
}

func regFilters(nt *jet.Set) {
	nt.AddGlobal("noescape", jet.SafeWriter(func(w io.Writer, b []byte) {
		w.Write(b)
	}))

	nt.AddGlobal("pathescape", jet.SafeWriter(func(w io.Writer, b []byte) {
		w.Write([]byte(url.PathEscape(string(b))))
	}))

	nt.AddGlobal("queryescape", jet.SafeWriter(func(w io.Writer, b []byte) {
		w.Write([]byte(url.QueryEscape(string(b))))
	}))
}

// open template directory
func Open(dir string) {
	tmpl = newTT(dir)
}

func renderTemplate(t *jet.Template, vars Vars) ([]byte, error) {

	var w bytes.Buffer

	if err := t.Execute(&w, vars, nil); err != nil {
		return nil, err
	}

	return w.Bytes(), nil

}

func Render(name string, vars Vars) ([]byte, error) {
	if tmpl == nil {
		err := errors.New("nit template object")
		return nil, err
	}

	t, err := tmpl.GetTemplate(name)
	if err != nil {
		return nil, err
	}

	return renderTemplate(t, vars)
}

func RenderString(template string, vars Vars) ([]byte, error) {

	nt := newTT()

	t, err := nt.LoadTemplate("template", template)
	if err != nil {
		return nil, err
	}

	return renderTemplate(t, vars)
}
