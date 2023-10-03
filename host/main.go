package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

var (
	r      wazero.Runtime
	code   wazero.CompiledModule
	binary = flag.String("wasm-binary", "../module/echo.wasm", "binary to run against every line of input from connections")
)

func main() {
	ctx := context.Background()

	data, err := os.ReadFile(*binary)
	if err != nil {
		log.Fatal(err)
	}

	r = wazero.NewRuntime(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	code, err = r.CompileModule(ctx, data)
	if err != nil {
		log.Fatal(err)
	}

	config := wazero.NewModuleConfig().
		WithStdout(os.Stdout).
		WithStdin(os.Stdin).
		WithStderr(os.Stderr).
		WithArgs("excellent").
		WithName("excellent")

	fmt.Println("Instantiating module")
	mod, err := r.InstantiateModule(ctx, code, config)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return
	}
	fmt.Println("Finished instantiating")
	defer mod.Close(ctx)
}
