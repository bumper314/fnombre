package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/spf13/pflag"
	"github.com/bumper314/fnombre"
)

const (
	usage = `fnombre genera nombres aleatorios y fáciles de recordar, como
"awful-fossil" o "constant-process".

Uso: 
  fnombre [opciones]

Ejemplos:
  # genera una sola frase de nombre usando las opciones predeterminadas
  fnombre

  # genera 10 nombres usando un delimitador personalizado
  fnombre --delimiter "." --quantity 10

Opciones:`

	contact = `
Author: Christopher Murphy <flyweight@pm.me>
Source: https://github.com/splode/fnombre`
)

var (
	version = ""
)

//go:embed banner
var banner []byte

func main() {
	pflag.Usage = generateUsage

	var (
		casing    string = "lower"
		delimiter string = "-"
		help      bool
		ver       bool
		quantity  int   = 1
		size      uint  = 2
		seed      int64 = -1
		// TODO: add option to use custom dictionary
	)

	pflag.StringVarP(&casing, "casing", "c", casing, "establece el formato de mayúsculas y minúsculas del nombre generado <title|upper|lower>")
	pflag.StringVarP(&delimiter, "delimiter", "d", delimiter, "establece el delimitador utilizado para unir palabras")
	pflag.IntVarP(&quantity, "quantity", "q", quantity, "establece la cantidad de nombres a generar")
	pflag.UintVarP(&size, "size", "z", size, "establece el número de palabras en el nombre generado (mínimo 2, máximo 4)")
	pflag.Int64VarP(&seed, "seed", "s", seed, "semilla para el generador aleatorio")
	pflag.BoolVarP(&help, "help", "h", help, "muestra el uso de fnombre")
	pflag.BoolVarP(&ver, "version", "v", ver, "muestra la versión de fnombre")
	pflag.Parse()

	if help {
		pflag.Usage()
		os.Exit(0)
	}

	if ver {
		fmt.Println(getVersion())
		os.Exit(0)
	}

	c, err := fnombre.CasingFromString(casing)
	handleError(err)

	opts := []fnombre.GeneratorOption{
		fnombre.WithCasing(c),
		fnombre.WithDelimiter(delimiter),
	}

	if seed != -1 {
		opts = append(opts, fnombre.WithSeed(seed))
	}
	if size != 2 {
		opts = append(opts, fnombre.WithSize(size))
	}

	rng := fnombre.NewGenerator(opts...)

	for i := 0; i < quantity; i++ {
		name, err := rng.Generate()
		handleError(err)
		fmt.Println(name)
	}
}

func generateUsage() {
	fmt.Println(string(banner))
	fmt.Println(usage)
	pflag.PrintDefaults()
	fmt.Println(contact)
}

func getVersion() string {
	if version != "" {
		return version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" {
		return "unknown"
	}

	version = info.Main.Version
	if info.Main.Sum != "" {
		version += fmt.Sprintf(" (%s)", info.Main.Sum)
	}

	return version
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
