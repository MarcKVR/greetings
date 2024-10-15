# Saludos en GO

Este paquete proporciona una forma simple de obtener saludos personales en GO

## Installation

Ejecuta el siguiente comando para instalar

```bash
go get -u github.com/MarcKVR/greetings
```

## Uso

aquí tienes un ejemplo de cómo utilizar el paquete en tu código

```bash

package main

import (
	"fmt"
	"github.com/MarcKVR/greetings"
)

func main() {
    message, err := greetings.Hello("Esteban")
}
```
