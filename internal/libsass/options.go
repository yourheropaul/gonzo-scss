package libsass

/*
#include "sass_context.h"
*/
import "C"

// Links:
// https://github.com/sass/libsass/wiki/API-Sass-Context-Internal
// https://github.com/sass/libsass/wiki/API-Sass-Context

type OutputStyle int

const (
	StyleNested OutputStyle = iota
	StyleExpanded
	StyleCompact
	StyleCompressed
)

type CompilerState int

const (
	CompilerCreated CompilerState = iota
	CompilerParsed
	CompilerExecuted
)

type Options interface {
	Precision() int
	GetOutputStyle() OutputStyle

	SourceComments() bool
	SourceMapEmbed() bool
	SourceMapContent() bool
	OmitSourceMapURL() bool
	IsIndentedSyntaxSrc() bool
	Indent() string
	Linefeed() string
	InputPath() string
	OutputPath() string
	PluginPath() string
	IncludePath() string
	SourceMapFile() string
	SourceMapRoot() string
	//Sass_Importer_List sass_option_get_c_headers (struct Sass_Options* options);
	//Sass_Importer_List sass_option_get_c_importers (struct Sass_Options* options);
	//Sass_Function_List sass_option_get_c_functions (struct Sass_Options* options);

	SetPrecision(precision int)
	SetOutputStyle(style OutputStyle)

	SetSourceComments(comment bool)
	SetSourceMapEmbed(embed bool)
	SetSourceMapContent(content bool)
	SetOmitSourceMapURL(omit bool)
	SetIsIndentedSyntaxSrc(is bool)
	SetIndent(indent string)
	SetLinefeed(linefeed string)
	SetInputPath(input string)
	SetOutputPath(output string)
	SetPluginPath(plugin string)
	SetIncludePath(include string)
	SetSourceMapFile(mapfile string)
	SetSourceMapRoot(maproot string)
	PushPluginPath(path string)
	PushIncludePath(path string)

	optc() *C.struct_Sass_Options
	//TODO:
	// Setters for son values
	// void sass_option_set_c_headers (struct Sass_Options* options, Sass_Importer_List c_headers);
	// void sass_option_set_c_importers (struct Sass_Options* options, Sass_Importer_List c_importers);
	// void sass_option_set_c_functions (struct Sass_Options* options, Sass_Function_List c_functions);
	// Push function for paths (no manipulation support for now)
}

func NewOptions() Options {
	return (*options)(C.sass_make_options())
}

type options C.struct_Sass_Options

func (opt *options) optc() *C.struct_Sass_Options {
	return (*C.struct_Sass_Options)(opt)
}

func (opt *options) Precision() int {
	return int(C.sass_option_get_precision(opt.optc()))
}

func (opt *options) SetPrecision(precision int) {
	C.sass_option_set_precision(opt.optc(), C.int(precision))
}

func (opt *options) GetOutputStyle() OutputStyle {
	return OutputStyle(C.sass_option_get_output_style(opt.optc()))
}

func (opt *options) SetOutputStyle(style OutputStyle) {
	C.sass_option_set_output_style(opt.optc(), uint32(style))
}

func (opt *options) SourceComments() bool {
	return bool(C.sass_option_get_source_comments(opt.optc()))
}

func (opt *options) SetSourceComments(comment bool) {
	C.sass_option_set_source_comments(opt.optc(), C._Bool(comment))
}

func (opt *options) SourceMapEmbed() bool {
	return bool(C.sass_option_get_source_map_embed(opt.optc()))
}

func (opt *options) SetSourceMapEmbed(embed bool) {
	C.sass_option_set_source_map_embed(opt.optc(), C._Bool(embed))
}

func (opt *options) SourceMapContent() bool {
	return bool(C.sass_option_get_source_map_contents(opt.optc()))
}

func (opt *options) SetSourceMapContent(content bool) {
	C.sass_option_set_source_map_contents(opt.optc(), C._Bool(content))
}

func (opt *options) OmitSourceMapURL() bool {
	return bool(C.sass_option_get_omit_source_map_url(opt.optc()))
}

func (opt *options) SetOmitSourceMapURL(omit bool) {
	C.sass_option_set_omit_source_map_url(opt.optc(), C._Bool(omit))
}

func (opt *options) IsIndentedSyntaxSrc() bool {
	return bool(C.sass_option_get_is_indented_syntax_src(opt.optc()))
}

func (opt *options) SetIsIndentedSyntaxSrc(is bool) {
	C.sass_option_set_is_indented_syntax_src(opt.optc(), C._Bool(is))
}

func (opt *options) Indent() string {
	return C.GoString(C.sass_option_get_indent(opt.optc()))
}

func (opt *options) SetIndent(indent string) {
	C.sass_option_set_indent(opt.optc(), C.CString(indent))
}

func (opt *options) Linefeed() string {
	return C.GoString(C.sass_option_get_linefeed(opt.optc()))
}

func (opt *options) SetLinefeed(linefeed string) {
	C.sass_option_set_linefeed(opt.optc(), C.CString(linefeed))
}

func (opt *options) InputPath() string {
	return C.GoString(C.sass_option_get_input_path(opt.optc()))
}

func (opt *options) SetInputPath(input string) {
	C.sass_option_set_input_path(opt.optc(), C.CString(input))
}

func (opt *options) OutputPath() string {
	return C.GoString(C.sass_option_get_output_path(opt.optc()))
}

func (opt *options) SetOutputPath(output string) {
	C.sass_option_set_output_path(opt.optc(), C.CString(output))
}

func (opt *options) PluginPath() string {
	return C.GoString(C.sass_option_get_plugin_path(opt.optc()))
}

func (opt *options) SetPluginPath(plugin string) {
	C.sass_option_set_plugin_path(opt.optc(), C.CString(plugin))
}

func (opt *options) IncludePath() string {
	return C.GoString(C.sass_option_get_include_path(opt.optc()))
}

func (opt *options) SetIncludePath(include string) {
	C.sass_option_set_include_path(opt.optc(), C.CString(include))
}

func (opt *options) SourceMapFile() string {
	return C.GoString(C.sass_option_get_source_map_file(opt.optc()))
}

func (opt *options) SetSourceMapFile(mapfile string) {
	C.sass_option_set_source_map_file(opt.optc(), C.CString(mapfile))
}

func (opt *options) SourceMapRoot() string {
	return C.GoString(C.sass_option_get_source_map_root(opt.optc()))
}

func (opt *options) SetSourceMapRoot(maproot string) {
	C.sass_option_set_source_map_root(opt.optc(), C.CString(maproot))
}

//Sass_Importer_List sass_option_get_c_headers (struct Sass_options* options);
//Sass_Importer_List sass_option_get_c_importers (struct Sass_options* options);
//Sass_Function_List sass_option_get_c_functions (struct Sass_options* options);

//func (opt *options) SetOutputStyle(style OutputStyle) {
//}

// Setters for son values
// void sass_option_set_c_headers (struct Sass_options* options, Sass_Importer_List c_headers);
// void sass_option_set_c_importers (struct Sass_options* options, Sass_Importer_List c_importers);
// void sass_option_set_c_functions (struct Sass_options* options, Sass_Function_List c_functions);
// Push function for paths (no manipulation support for now)
func (opt *options) PushPluginPath(path string) {
	C.sass_option_push_plugin_path(opt.optc(), C.CString(path))
}
func (opt *options) PushIncludePath(path string) {
	C.sass_option_push_include_path(opt.optc(), C.CString(path))
}
