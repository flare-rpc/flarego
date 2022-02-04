package main

import (
	"flag"
	"fmt"
	"go/build"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/flare-rpc/flare-go/tool/xgen/parser"
)

var (
	processPkg    = flag.Bool("pkg", false, "process the whole package instead of just the given file")
	specifiedName = flag.String("o", "", "specify the filename of the output")
	buildTags     = flag.String("tags", "", "build tags to add to generated file")
	registry      = flag.String("r", "", "registry type. support etcd, consul, zookeeper, mdns")
)

func main() {
	flag.Parse()

	files := flag.Args()

	if *processPkg {
		if len(flag.Args()) == 0 {
			log.Fatalf("not set packages")
		}
		var pkgFiles []string
		for _, f := range files {
			pkgFiles = append(pkgFiles, filepath.Join(filepath.Join(build.Default.GOPATH, "src", f)))
		}
		files = pkgFiles
	}

	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var parsers []*parser.Parser

	for _, fname := range files {
		fInfo, err := os.Stat(fname)
		if err != nil {
			panic(err)
		}

		if !filepath.IsAbs(fname) {
			fname, err = filepath.Abs(fname)
			if err != nil {
				panic(err)
			}
		}

		relDir, err := filepath.Rel(filepath.Join(build.Default.GOPATH, "src"), fname)
		if err != nil {
			fmt.Printf("provided directory not under GOPATH (%s): %v",
				build.Default.GOPATH, err)
			return
		}

		p := &parser.Parser{PkgFullName: relDir, StructNames: make(map[string]bool)}
		if err := p.Parse(fname, fInfo.IsDir()); err != nil {
			fmt.Printf("Error parsing %v: %v", fname, err)
			return
		}

		parsers = append(parsers, p)
		_ = generate(parsers)
	}

}

func generate(parsers []*parser.Parser) error {
	var w io.Writer
	if *specifiedName == "" {
		w = os.Stdout
	} else {
		f, err := os.Create(*specifiedName)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f
	}

	if *buildTags != "" {
		fmt.Fprintln(w, "// +build ", *buildTags)
		fmt.Fprintln(w)
	}
	fmt.Fprintln(w, "// AUTOGENERATED FILE: rpcx server stub code")
	fmt.Fprintln(w, "// compilable during generation.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "package main")

	fmt.Fprintln(w)
	fmt.Fprintln(w, "import (")
	fmt.Fprintln(w, `  "flag"`)
	fmt.Fprintln(w, `  "time"`)
	fmt.Fprintln(w)
	fmt.Fprintln(w, `  metrics "github.com/rcrowley/go-metrics"`)
	fmt.Fprintln(w, `  "github.com/flare-rpc/flare-go/server"`)
	fmt.Fprintln(w, `  "github.com/flare-rpc/flare-go/serverplugin"`)

	var importedPackages = make(map[string]bool)
	for _, p := range parsers {
		if !importedPackages[p.PkgFullName] {
			fmt.Fprintln(w, `  "`+p.PkgFullName+`"`)
			importedPackages[p.PkgFullName] = true
		}
	}
	fmt.Fprintln(w, ")")
	fmt.Fprintln(w)

	var mainFn = `
var (
	addr     = flag.String("addr", "localhost:8972", "server address")`
	if *registry == "etcd" || *registry == "consul" || *registry == "zookeeper" {
		mainFn = mainFn + `
	rAddr = flag.String("rAddr", "localhost:2379", "register address")`
	}

	mainFn = mainFn + `
	basePath = flag.String("base", "/rpcx", "prefix path")
)
	
func main() {
	flag.Parse()

	_ = time.Second
	_ = metrics.UseNilMetrics
	_ = serverplugin.GetFunctionName

	s := server.NewServer()
	addRegistryPlugin(s)

	registerServices(s)

	s.Serve("tcp", *addr)
}`
	fmt.Fprintln(w, mainFn)

	fmt.Fprintln(w, `func registerServices(s *server.Server) {`)
	for _, p := range parsers {
		for n := range p.StructNames {
			fmt.Fprintln(w, `	s.Register(new(`+p.PkgName+"."+n+`), "")`)
		}
	}

	fmt.Fprintln(w, `}`)

	useRegistry := false

	fmt.Fprintln(w, `func addRegistryPlugin(s *server.Server) {`)
	switch *registry {
	case "etcd":
		fmt.Fprintln(w, `	// add registery
	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*rAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}`)
		useRegistry = true
	case "consul":
		fmt.Fprintln(w, `	// add registery
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		ConsulServers:  []string{*rAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}`)
		useRegistry = true
	case "zookeeper":
		fmt.Fprintln(w, `	// add registery
	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *addr,
		ZooKeeperServers: []string{*rAddr},
		BasePath:         *basePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}`)
		useRegistry = true
	case "mdns":
		fmt.Fprintln(w, `
			r := serverplugin.NewMDNSRegisterPlugin("tcp@"+*addr, 8972, metrics.NewRegistry(), time.Minute, "")`)
		useRegistry = true
	default:
		fmt.Fprintln(w, `
		`)
	}

	if useRegistry {
		fmt.Fprintln(w, `
	err := r.Start()
	if err != nil {
		//log.Fatal(err)
	}
	s.Plugins.Add(r)`)
	}

	fmt.Fprintln(w, `}`)
	return nil
}
