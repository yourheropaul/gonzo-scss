package libsass

import "testing"

func TestDataContext(t *testing.T) {

	scss_src := "div { a { color: blue; font-size: 5px} }"

	dc := NewDataContext(scss_src)
	opt := dc.Context().Options()
	opt.SetPrecision(10)
	opt.SetSourceComments(true)
	//??
	//dc.SetOptions(opt)

	compiler := dc.Compiler()
	err := compiler.Parse()
	if err != nil {
		t.Fatal(err)
	}

	err = compiler.Execute()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(dc.Context().OutputString())
}
