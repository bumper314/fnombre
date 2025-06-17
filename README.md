# fnombre

Este es un fork rápido y sucio en español de https://github.com/Splode/fname.

Genera nombres aleatorios y fáciles de recordar, como `determinado-panqueque` o `descubrimiento siniestro`. fnombre es como un *diceware* gramaticalmente consciente para nombres o identificadores únicos.

fnombre no pretende proporcionar un identificador seguro y único a nivel global, pero con más de 500 mil millones de combinaciones posibles, es suficientemente bueno para la mayoría de los casos de uso no críticos.

## Tabla de Contenidos

- [fnombre](#fnombre)
  - [Tabla de Contenidos](#tabla-de-contenidos)
  - [Instalación](#instalación)
    - [Go](#go)
    - [Fuente](#fuente)
  - [Uso](#uso)
    - [CLI](#cli)
    - [Librería](#librería)
      - [Instalación](#instalación-1)
      - [Uso Básico](#uso-básico)
      - [Personalización](#personalización)
  - [Descargos de Responsabilidad](#descargos-de-responsabilidad)
  - [Contribuciones](#contribuciones)
    - [Reportar Problemas](#reportar-problemas)
    - [Sugerir Mejoras](#sugerir-mejoras)
  - [Licencia](#licencia)
  - [Proyectos Relacionados](#proyectos-relacionados)

## Instalación

### Go

```sh
go install github.com/bumper314/fnombre/cmd/fnombre@latest
```

### Fuente

```sh
git clone https://github.com/bumper314/fnombre.git
cd fnombre
go install ./cmd/fnombre
```

## Uso

### CLI
  
Genera una sola frase de nombre aleatoria:

```sh
$ fnombre
extinto-verde
```

Genera múltiples frases de nombre, pasando el número de nombres como argumento:

```sh
$ fnombre --quantity 3
influente-longitud
directo-oído
almacenamiento-cultural
```

Genera una frase de nombre con un delimitador personalizado:

```sh
$ fnombre --delimiter "__"
llamativa__percepción
```

Genera una frase de nombre con más palabras:

```sh
$ fnombre --size 3
vengativo-juguete-identificado

$ fnombre --size 4
hechizante-proyecto-presentado-completamente
```

Nota: el tamaño mínimo de la frase es 2 (predeterminado), y el máximo es 4.

Genera una frase de nombre con un formato específico:

```sh
$ fnombre --casing upper
TRÁGICA-MONTAÑA

$ fnombre --casing title
Fiesta-Caprichosa
```

Especifica la semilla para generar nombres:

```sh
$ fnombre --seed 123 --quantity 2
alegre-gozo
inquilino-elegible

$ fnombre --seed 123 --quantity 2
alegre-gozo
inquilino-elegible
```

### Librería

#### Instalación

```sh
go get github.com/bumper314/fnombre
```

#### Uso Básico

```go
package main

import (
  "fmt"

  "github.com/bumper314/fnombre"
)

func main() {
  rng := fnombre.NewGenerator()
  phrase, err := rng.Generate()
  fmt.Println(phrase)
  // => "influente-longitud"
}
```

#### Personalización

```go
package main

import (
  "fmt"

  "github.com/bumper314/fnombre"
)

func main() {
  rng := fnombre.NewGenerator(fnombre.WithDelimiter("__"), fnombre.WithSize(3))
  phrase, err := rng.Generate()
  fmt.Println(phrase)
  // => "establecido__tiburón__destruido"
}
```

## Descargos de Responsabilidad

fnombre no es seguro criptográficamente, y no debe ser utilizado para nada que requiera un identificador verdaderamente único. Está destinado a ser una alternativa divertida y fácil de recordar a los UUIDs.

El diccionario de fnombre está curado para excluir palabras que sean ofensivas, o que puedan considerarse ofensivas, ya sea solas o cuando se generen en una frase. Sin embargo, no todos los casos están ni pueden estar cubiertos. Si encuentras una palabra que piensas que debería ser eliminada, por favor [abre un problema](https://github.com/bumper314/fnombre/issues).

## Contribuciones

¡Damos la bienvenida a las contribuciones al proyecto fnombre! Ya sea reportando errores, sugiriendo mejoras o enviando nuevas características, tu aporte es valioso para nosotros. Aquí te mostramos cómo puedes comenzar:

1. Haz un fork del repositorio en GitHub.
2. Clona tu fork y crea una nueva rama para tus cambios.
3. Realiza tus cambios y haz commit en tu rama.
4. Crea una solicitud de pull, y proporciona una descripción clara de tus cambios.

Antes de enviar una solicitud de pull, por favor asegúrate de que tus cambios estén bien probados y se adhieran al estilo de código utilizado en todo el proyecto. Si no estás seguro de cómo proceder o necesitas ayuda, no dudes en abrir un problema o hacer una pregunta en la sección de [discusiones](https://github.com/bumper314/fnombre/discussions).

### Reportar Problemas

Si encuentras un error o cualquier problema, por favor [abre un problema](https://github.com/bumper314/fnombre/issues) en GitHub. Al reportar un error, trata de incluir tanta información como sea posible, como los pasos para reproducir el problema, el comportamiento esperado y el comportamiento actual. Esto nos ayudará a diagnosticar y solucionar el problema de manera más eficiente.

### Sugerir Mejoras

Siempre estamos buscando formas de mejorar fnombre. Si tienes una sugerencia para una nueva característica o una mejora a una característica existente, por favor [abre un problema](https://github.com/bumper314/fnombre/issues) o inicia una discusión en la sección de [discusiones](https://github.com/bumper314/fnombre/discussions). Asegúrate de explicar tu idea en detalle, y si es posible, proporciona ejemplos o casos de uso.

¡Gracias por tu interés en contribuir a fnombre!

## Licencia

[Licencia MIT](./LICENSE)

## Proyectos Relacionados

- [go-diceware](https://github.com/sethvargo/go-diceware)
- [wordnet-random-name](https://github.com/kohsuke/wordnet-random-name)