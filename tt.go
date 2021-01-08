package tt

import (
	"bytes"
	"errors"
	"io"
	"net/url"

	"github.com/CloudyKit/jet"
)

type Vars = jet.VarMap

type TT struct {
	tt *jet.Set
}

var (
	global *TT = New()

	ErrNilTemplate error = errors.New("nil template object")
)

func New(dirs ...string) *TT {
	tt := &TT{tt: jet.NewHTMLSet(dirs...)}
	tt.regFilters()
	return tt
}

func (tt *TT) MakeVars() Vars {
	return make(Vars)
}

func MakeVars() Vars {
	return global.MakeVars()
}

func (tt *TT) regFilters() {
	tt.tt.AddGlobal("noescape", jet.SafeWriter(func(w io.Writer, b []byte) {
		w.Write(b)
	}))

	tt.tt.AddGlobal("pathescape", jet.SafeWriter(func(w io.Writer, b []byte) {
		w.Write([]byte(url.PathEscape(string(b))))
	}))

	tt.tt.AddGlobal("queryescape", jet.SafeWriter(func(w io.Writer, b []byte) {
		w.Write([]byte(url.QueryEscape(string(b))))
	}))
}

func Open(dir string) {
	global = New(dir)
}

func renderTemplate(t *jet.Template, vars Vars) ([]byte, error) {

	var w bytes.Buffer

	if err := t.Execute(&w, vars, nil); err != nil {
		return nil, err
	}

	return w.Bytes(), nil

}

func (tt *TT) Render(name string, vars Vars) ([]byte, error) {
	if tt == nil || tt.tt == nil {
		return nil, ErrNilTemplate
	}

	t, err := tt.tt.GetTemplate(name)
	if err != nil {
		return nil, err
	}

	return renderTemplate(t, vars)
}

func Render(name string, vars Vars) ([]byte, error) {
	return global.Render(name, vars)
}

func (tt *TT) RenderString(template string, vars Vars) ([]byte, error) {

	nt := New()

	t, err := nt.tt.LoadTemplate("template", template)
	if err != nil {
		return nil, err
	}

	return renderTemplate(t, vars)
}

func RenderString(template string, vars Vars) ([]byte, error) {
	return global.RenderString(template, vars)
}
