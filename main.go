// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers0 "github.com/revel/modules/testrunner/app/controllers"
	"github.com/revel/revel"
	"github.com/revel/revel/testing"
	"reflect"
	_ "webapp/app"
	controllers1 "webapp/app/controllers"
	"webapp/tests"
)

var (
	runMode    = flag.String("runMode", "dev", "Run mode.")
	port       = flag.Int("port", 0, "By default, read from app.conf")
	importPath = flag.String("importPath", "webapp", "Go Import Path for the app.")
	srcPath    = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.AppLog.Info("Running revel server")

	revel.RegisterController((*controllers.Static)(nil),
		[]*revel.MethodType{
			{
				Name: "Serve",
				Args: []*revel.MethodArg{
					{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			{
				Name: "ServeDir",
				Args: []*revel.MethodArg{
					{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			{
				Name: "ServeModule",
				Args: []*revel.MethodArg{
					{Name: "moduleName", Type: reflect.TypeOf((*string)(nil))},
					{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			{
				Name: "ServeModuleDir",
				Args: []*revel.MethodArg{
					{Name: "moduleName", Type: reflect.TypeOf((*string)(nil))},
					{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers0.TestRunner)(nil),
		[]*revel.MethodType{
			{
				Name: "Index",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					76: {
						"testSuites",
					},
				},
			},
			{
				Name: "Suite",
				Args: []*revel.MethodArg{
					{Name: "suite", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			{
				Name: "Run",
				Args: []*revel.MethodArg{
					{Name: "suite", Type: reflect.TypeOf((*string)(nil))},
					{Name: "test", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
					125: {},
				},
			},
			{
				Name:           "List",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers1.App)(nil),
		[]*revel.MethodType{
			{
				Name: "Index",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					12: {},
				},
			},
			{
				Name: "Users",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					17: {},
				},
			},
		})

	revel.RegisterController((*controllers1.BaseController)(nil),
		[]*revel.MethodType{})

	revel.RegisterController((*controllers1.CommentController)(nil),
		[]*revel.MethodType{
			{
				Name:           "Index",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			{
				Name: "Show",
				Args: []*revel.MethodArg{
					{Name: "id", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			{
				Name:           "Create",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			{
				Name:           "Update",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			{
				Name: "Delete",
				Args: []*revel.MethodArg{
					{Name: "id", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers1.User)(nil),
		[]*revel.MethodType{
			{
				Name:           "Index",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.DefaultValidationKeys = map[string]map[int]string{}
	testing.TestSuites = []interface{}{
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
