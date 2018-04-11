package main

// there is a lot of better library out there. specially I hate the standard flag library.
// but I want this one, clean and tidy. so hell with it :)
import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	listFile *string
	foxy     *string
	pac      *string
	ppx      *string
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func openFile(f string) (io.WriteCloser, error) {
	fl, err := os.Create(f)
	if err != nil {
		return nil, err
	}
	return fl, nil
}

func load(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var domains []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		domains = append(domains, sc.Text())
	}
	return domains, sc.Err()
}

func execute(p string, fn func(io.Writer, ...string) error, domains ...string) error {
	if p != "" {
		fl, err := openFile(p)
		check(err)
		defer fl.Close()

		return fn(fl, domains...)
	}

	return nil
}

func main() {
	flag.Parse()
	domains, err := load(*listFile)
	check(err)

	check(execute(*foxy, foxyProxy, domains...))
	check(execute(*ppx, proxifier, domains...))
	check(execute(*pac, pacFile, domains...))
}

func init() {
	p, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	p = filepath.Join(p, "domains")
	listFile = flag.String("domains", p, "domains file list")
	foxy = flag.String("foxyproxy", "", "path to build foxyproxy pattern, empty means ignore")
	pac = flag.String("pac", "", "path to build pac file, empty means ignore")
	ppx = flag.String("proxifier", "", "path to build proxifier config, empty means ignore")
}
