package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rhumb-lang/rhi/internal/compiler"
	"github.com/rhumb-lang/rhi/internal/config"
	"github.com/rhumb-lang/rhi/internal/loader"
	mapval "github.com/rhumb-lang/rhi/internal/map"
	_ "github.com/rhumb-lang/rhi/internal/natlib" // Register nat_math
	// Import the new package
	"github.com/rhumb-lang/rhi/internal/tooling"
	"github.com/rhumb-lang/rhi/internal/vm"
)

var (
	traceParser   = flag.Bool("trace-parser", false, "Enable parser tracing")
	traceBytecode = flag.Bool("trace-bytecode", false, "Enable bytecode tracing")
	traceStack    = flag.Bool("trace-stack", false, "Enable stack tracing")
	traceFrame    = flag.Bool("trace-frame", false, "Enable frame tracing")
	traceSpace    = flag.Bool("trace-space", false, "Enable space/concurrency tracing")
	traceLoader   = flag.Bool("trace-loader", false, "Enable loader tracing")
	lastValue     = flag.Bool("last-value", false, "Print the last value of the execution")
	testMode      = flag.Bool("test", false, "Run in test mode (check assertions)")
	sha256Flag    = flag.Bool("sha256", false, "Auto-fill empty anchors")
)

type Session struct {
	Compiler *compiler.Compiler
	VM       *vm.VM
	IsRepl   bool
}

func findProjectRoot(startPath string) string {
	dir := startPath
	if stat, err := os.Stat(dir); err == nil && !stat.IsDir() {
		dir = filepath.Dir(dir)
	}

	for {
		entries, err := os.ReadDir(dir)
		if err == nil {
			for _, e := range entries {
				if !e.IsDir() && strings.HasSuffix(e.Name(), ".rhy") {
					return dir
				}
			}
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	// Fallback to CWD
	cwd, _ := os.Getwd()
	return cwd
}

func NewSession(cfg *config.Config, scriptPath string) *Session {
	v := vm.NewVM()
	v.Config = cfg

	// Initialize Loader
	projectRoot := findProjectRoot(scriptPath)

	// Load Catalog to check for SourceRoot
	var rootCatalog *loader.Catalog
	var sourceRoot string
	entries, _ := os.ReadDir(projectRoot)
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".rhy") {
			catalog, err := loader.LoadCatalog(filepath.Join(projectRoot, e.Name()))
			if err == nil {
				rootCatalog = catalog
				if catalog.SourceRoot != "" {
					sourceRoot = catalog.SourceRoot
				}
			}
			break
		}
	}

	// Adjust ProjectRoot to include SourceRoot
	finalRoot := projectRoot
	if sourceRoot != "" {
		finalRoot = filepath.Join(projectRoot, sourceRoot)
	}

	v.Loader = &loader.LibraryLoader{
		Registry:       make(map[string]mapval.Value),
		Sitemap:        make(map[string]string),
		ProjectRoot:    finalRoot,
		Config:         cfg,
		VM:             v,
		RootCatalog:    rootCatalog, // Assign the loaded catalog
		CurrentRuntime: config.RuntimeVersion,
	}

	return &Session{
		Compiler: compiler.NewCompiler(),
		VM:       v,
	}
}

func main() {
	flag.Parse()

	// Environment variable overrides
	if !*traceParser && os.Getenv("RHI_TRACE_PARSER") == "1" {
		*traceParser = true
	}
	if !*traceBytecode && os.Getenv("RHI_TRACE_BYTECODE") == "1" {
		*traceBytecode = true
	}
	if !*traceStack && os.Getenv("RHI_TRACE_STACK") == "1" {
		*traceStack = true
	}
	if !*traceFrame && os.Getenv("RHI_TRACE_FRAME") == "1" {
		*traceFrame = true
	}
	if !*traceSpace && os.Getenv("RHI_TRACE_SPACE") == "1" {
		*traceSpace = true
	}
	if !*traceLoader && os.Getenv("RHI_TRACE_LOADER") == "1" {
		*traceLoader = true
	}

	cfg := &config.Config{
		TraceParser:   *traceParser,
		TraceBytecode: *traceBytecode,
		TraceStack:    *traceStack,
		TraceFrame:    *traceFrame,
		TraceSpace:    *traceSpace,
		TraceLoader:   *traceLoader,
	}

	args := flag.Args()

	if len(args) == 0 {
		scriptPath, _ := os.Getwd()
		// Pre-Flight Integrity Check for REPL (using CWD)
		projectRoot := findProjectRoot(scriptPath)
		entries, _ := os.ReadDir(projectRoot)
		var catalogPath string
		for _, e := range entries {
			if strings.HasSuffix(e.Name(), ".rhy") {
				catalogPath = filepath.Join(projectRoot, e.Name())
				break
			}
		}

		if catalogPath != "" {
			if err := tooling.EnsureIntegrity(catalogPath, projectRoot, *sha256Flag); err != nil {
				fmt.Fprintf(os.Stderr, "Panic: %v\n", err)
				os.Exit(1)
			}
		}

		session := NewSession(cfg, scriptPath)
		session.IsRepl = true
		session.runREPL()
	} else {
		failed := false
		for _, arg := range args {
			scriptPath, _ := filepath.Abs(arg)

			// Pre-Flight Integrity Check
			projectRoot := findProjectRoot(scriptPath)
			entries, _ := os.ReadDir(projectRoot)
			var catalogPath string
			for _, e := range entries {
				if strings.HasSuffix(e.Name(), ".rhy") {
					catalogPath = filepath.Join(projectRoot, e.Name())
					break
				}
			}

			if catalogPath != "" {
				if err := tooling.EnsureIntegrity(catalogPath, projectRoot, *sha256Flag); err != nil {
					fmt.Fprintf(os.Stderr, "Panic: %v\n", err)
					os.Exit(1)
				}
			}

			session := NewSession(cfg, scriptPath)
			if *testMode {
				fmt.Printf("Checking assertions in file '%s'...\n", filepath.Base(scriptPath))
			}
			if !session.runFile(arg) {
				failed = true
			}
		}

		if failed {
			os.Exit(1)
		}
	}
}
