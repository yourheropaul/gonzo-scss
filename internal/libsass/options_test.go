package libsass

import "testing"

func TestOptionsPrecision(t *testing.T) {
	opt := NewOptions()

	expect := 2
	opt.SetPrecision(expect)
	got := opt.Precision()
	if got != expect {
		t.Fatalf("Options.Precison: Expectd %d Got %d", expect, got)
	}
}

func TestOutputStyle(t *testing.T) {
	opt := NewOptions()

	expect := StyleNested
	opt.SetOutputStyle(expect)
	got := opt.GetOutputStyle()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %v, Got %v", expect, got)
	}
}

func TestSourceComments(t *testing.T) {
	opt := NewOptions()

	expect := true
	opt.SetSourceComments(expect)
	got := opt.SourceComments()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %t, Got %t", expect, got)
	}
}

func TestSourceMapEmbed(t *testing.T) {
	opt := NewOptions()

	expect := true
	opt.SetSourceMapEmbed(expect)
	got := opt.SourceMapEmbed()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %t, Got %t", expect, got)
	}
}

func TestSourceMapContent(t *testing.T) {
	opt := NewOptions()

	expect := true
	opt.SetSourceMapContent(expect)
	got := opt.SourceMapContent()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %t,Got %t", expect, got)
	}
}

func TestOmitSourceMapURL(t *testing.T) {
	opt := NewOptions()

	expect := true
	opt.SetOmitSourceMapURL(expect)
	got := opt.OmitSourceMapURL()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %t, Got %t", expect, got)
	}
}

func TestIsIndentedSyntaxSrc(t *testing.T) {
	opt := NewOptions()

	expect := true
	opt.SetIsIndentedSyntaxSrc(expect)
	got := opt.IsIndentedSyntaxSrc()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %t, Got %t", expect, got)
	}
}

func TestIndent(t *testing.T) {
	opt := NewOptions()

	expect := "string"
	opt.SetIndent(expect)
	got := opt.Indent()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %s, Got %s", expect, got)
	}
}

func TestLinefeed(t *testing.T) {
	opt := NewOptions()

	expect := "string"
	opt.SetLinefeed(expect)
	got := opt.Linefeed()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %s, Got %s", expect, got)
	}
}

func TestInputPath(t *testing.T) {
	opt := NewOptions()

	expect := "string"
	opt.SetInputPath(expect)
	got := opt.InputPath()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %s, Got %s", expect, got)
	}
}

func TestOutputPath(t *testing.T) {
	opt := NewOptions()

	expect := "string"
	opt.SetOutputPath(expect)
	got := opt.OutputPath()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %s, Got %s", expect, got)
	}
}

func TestPluginPath(t *testing.T) {
	opt := NewOptions()

	expect := "string"
	opt.SetPluginPath(expect)
	got := opt.PluginPath()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %s, Got %s", expect, got)
	}
}

func TestIncludePath(t *testing.T) {
	opt := NewOptions()

	expect := "string"
	opt.SetIncludePath(expect)
	got := opt.IncludePath()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %s, Got %s", expect, got)
	}
}

func TestSourceMapFile(t *testing.T) {
	opt := NewOptions()

	expect := "string"
	opt.SetSourceMapFile(expect)
	got := opt.SourceMapFile()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %s, Got %s", expect, got)
	}
}

func TestSourceMapRoot(t *testing.T) {
	opt := NewOptions()

	expect := "string"
	opt.SetSourceMapRoot(expect)
	got := opt.SourceMapRoot()
	if got != expect {
		t.Fatalf("Options.OutputStyle: Expected %s, Got %s", expect, got)
	}
}

//TODO(rudes):Figure out how to test this
func TestPushPluginPath(t *testing.T) {
	t.Skip("TODO(rudes):Figure out how to test this")
	opt := NewOptions()

	s := "string"
	opt.PushPluginPath(s)
}

//TODO(rudes):Figure out how to test this
func TestPushIncludePath(t *testing.T) {
	t.Skip("TODO(rudes):Figure out how to test this")
	opt := NewOptions()

	s := "string"
	opt.PushIncludePath(s)
}
