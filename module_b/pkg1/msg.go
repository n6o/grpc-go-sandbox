package pkg1

import apkg "github.com/n6o/grpc-go-sandbox/module_a/pkg1/subpkg"

func HeHello() string {
	return "He" + apkg.Hello()
}
