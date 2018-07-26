// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	"github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers0 "github.com/revel/modules/testrunner/app/controllers"
	_ "webapp/app"
	controllers1 "webapp/app/controllers"
	"webapp/tests"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.AppLog.Info("Running revel server")

	revel.RegisterController((*controllers.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
				},
			},
			&revel.MethodType{
				Name: "ServeDir",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
				},
			},
			&revel.MethodType{
				Name: "ServeModuleDir",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
				},
			},
		})

	revel.RegisterController((*controllers0.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{
				},
				RenderArgNames: map[int][]string{
					76: []string{
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
					125: []string{
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{
				},
				RenderArgNames: map[int][]string{
				},
			},
		})

	revel.RegisterController((*controllers1.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{
				},
				RenderArgNames: map[int][]string{
					12: []string{
					},
				},
			},
			&revel.MethodType{
				Name: "Users",
				Args: []*revel.MethodArg{
				},
				RenderArgNames: map[int][]string{
					17: []string{
					},
				},
			},
		})

	revel.RegisterController((*controllers1.BaseController)(nil),
		[]*revel.MethodType{

		})

	revel.RegisterController((*controllers1.CommentController)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{
				},
				RenderArgNames: map[int][]string{
				},
			},
			&revel.MethodType{
				Name: "Show",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
				},
			},
			&revel.MethodType{
				Name: "Create",
				Args: []*revel.MethodArg{
				},
				RenderArgNames: map[int][]string{
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{
				},
				RenderArgNames: map[int][]string{
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
				},
			},
		})

	revel.RegisterController((*controllers1.User)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{
				},
				RenderArgNames: map[int][]string{
				},
			},
		})

	revel.DefaultValidationKeys = map[string]map[int]string{
	}
	testing.TestSuites = []interface{}{
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
