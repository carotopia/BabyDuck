package main

import (
	"BabyDuckCompiler/builder"
	"BabyDuckCompiler/quads"
	"BabyDuckCompiler/vm"
	"fmt"
	"fyne.io/fyne/v2"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Crear aplicación
	myApp := app.New()
	myApp.SetIcon(theme.DocumentIcon())
	myWindow := myApp.NewWindow("🐣 BabyDuck Compiler 🐣")
	myWindow.Resize(fyne.NewSize(1200, 800))

	// Editor de código
	codeEditor := widget.NewMultiLineEntry()
	codeEditor.SetPlaceHolder("Escribe tu código BabyDuck aquí...")
	codeEditor.Wrapping = fyne.TextWrapWord

	// Código de ejemplo por defecto
	defaultCode := `

program factorial;

var n, factorial, i: int;

void printNumber(n : int) [
{
    print("El numero es", n + n * 10);
}
];

void a(n: int) [
   {
    printNumber(n);
   }
];

main
{
    n = 5;
    factorial = 1;

    i = n;
    while (i > 1) do {
        factorial = factorial * i;
        i = i - 1;
    };

    print("factorial", factorial);
    a(5);
}

end`
	codeEditor.SetText(defaultCode)

	// Área de salida
	outputArea := widget.NewMultiLineEntry()
	outputArea.SetPlaceHolder("La salida aparecerá aquí...")
	outputArea.MultiLine = true
	outputArea.Wrapping = fyne.TextWrapWord

	// Checkbox para debug
	debugCheck := widget.NewCheck("Modo Debug", nil)

	// Botón compilar
	compileBtn := widget.NewButton("🌟 Compilar y Ejecutar 🌟", func() {
		sourceCode := codeEditor.Text
		debug := debugCheck.Checked

		if strings.TrimSpace(sourceCode) == "" {
			outputArea.SetText("🍓 Error: No hay código para  🍓")
			return
		}

		// Compilar y ejecutar
		output := compileAndRun(sourceCode, debug)
		outputArea.SetText(output)
	})
	compileBtn.Importance = widget.HighImportance

	// Botón limpiar
	clearBtn := widget.NewButton("🥨️ Limpiar Editor", func() {
		codeEditor.SetText("")
		outputArea.SetText("")
	})

	// Botón ejemplo
	exampleBtn := widget.NewButton("🐱 Cargar Ejemplo", func() {
		exampleCode := `program Fibonacci;

program Fibonacci;

var a, b, temp: int;

main 
{ 
    a = 0;
    b = 1;

    while (a < 100)  { 
        print(a);
        temp = a + b; 
        a = b;
        b = temp;
    } do;
print("Hola", 5*5 /5  + (3 -1));

}

end

`
		codeEditor.SetText(exampleCode)
		outputArea.SetText("Ejemplo cargado. Presiona 'Compilar y Ejecutar'")
	})

	// Botón acerca de
	aboutBtn := widget.NewButton("🌠 Acerca de", func() {
		dialog.ShowInformation("Acerca de BabyDuck Compiler",
			"🐥 BabyDuck Compiler 🐥\n\n"+
				"Compilador con máquina virtual integrada\n"+
				"Características:\n"+
				"• Funciones con parámetros\n"+
				"• Ciclos while\n"+
				"• Operaciones aritméticas\n"+
				"• Variables locales y globales\n"+
				"• Debug paso a paso\n\n"+
				"¡Desarrollado en Go!", myWindow)
	})

	// Layout de botones
	buttonContainer := container.NewHBox(
		compileBtn,
		clearBtn,
		exampleBtn,
		debugCheck,
		aboutBtn,
	)

	// Layout principal
	leftPanel := container.NewBorder(
		widget.NewLabel("🌼 Editor de Código 🌼"),
		buttonContainer,
		nil, nil,
		codeEditor,
	)

	rightPanel := container.NewBorder(
		widget.NewLabel("⭐ Salida del Compilador ⭐"),
		nil, nil, nil,
		outputArea,
	)

	// Dividir pantalla
	content := container.NewHSplit(leftPanel, rightPanel)
	content.SetOffset(0.5) // 50-50 split

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

// Función para compilar y ejecutar código
func compileAndRun(sourceCode string, debug bool) string {
	var output strings.Builder

	output.WriteString("🌟 COMPILANDO 🌟\n")
	output.WriteString(strings.Repeat("=", 50) + "\n")

	// Crear parser
	parser := builder.NewPureVisitorParser(sourceCode, debug)

	// Compilar
	symbolTable, errors := parser.Parse()

	if len(errors) > 0 {
		output.WriteString("🍓 ERRORES DE COMPILACIÓN 🍓\n")
		for i, err := range errors {
			output.WriteString(fmt.Sprintf("%d. %s\n", i+1, err))
		}
		return output.String()
	}

	output.WriteString("✨ Compilación exitosa! ✨\n\n")

	// Mostrar tabla de símbolos si está en debug
	if debug && symbolTable != nil {
		output.WriteString("🍰 TABLA DE SÍMBOLOS 🍰\n")
		for scope, funcInfo := range symbolTable.Directory {
			output.WriteString(fmt.Sprintf("  Scope: %s\n", scope))
			if len(funcInfo.Variables) > 0 {
				for varName, varDetails := range funcInfo.Variables {
					output.WriteString(fmt.Sprintf("    %s (%s) -> %d\n",
						varName, varDetails.Type, varDetails.MemoryAddress))
				}
			}
		}
		output.WriteString("\n")
	}

	// Obtener cuádruplos
	quadruples := parser.GetQuadruples()
	var typedQuadruples []quads.Quadruple
	for _, q := range quadruples {
		if quad, ok := q.(quads.Quadruple); ok {
			typedQuadruples = append(typedQuadruples, quad)
		}
	}

	if debug && len(typedQuadruples) > 0 {
		output.WriteString("💫 CUÁDRUPLOS GENERADOS 💫 :\n")
		for i, quad := range typedQuadruples {
			output.WriteString(fmt.Sprintf("%3d: %-10s %-10v %-10v %-10v\n",
				i, quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result))
		}
		output.WriteString("\n")
	}

	// Ejecutar con máquina virtual
	if len(typedQuadruples) > 0 {
		output.WriteString("⚡ EJECUTANDO ⚡ \n")
		output.WriteString(strings.Repeat("=", 50) + "\n")

		// Crear VM
		virtualMachine := vm.NewVirtualMachine(debug)

		// NUEVO: Configurar writers para capturar output
		var programOutput strings.Builder
		var debugOutput strings.Builder

		virtualMachine.SetOutputWriter(&programOutput)
		if debug {
			virtualMachine.SetDebugWriter(&debugOutput)
		}

		// Convertir cuádruplos
		vmQuadruples := make([]vm.Quadruple, len(typedQuadruples))
		for i, quad := range typedQuadruples {
			vmQuadruples[i] = vm.Quadruple{
				Operator:     quad.Operator,
				LeftOperand:  quad.LeftOperand,
				RightOperand: quad.RightOperand,
				Result:       quad.Result,
			}
		}

		virtualMachine.LoadQuadruples(vmQuadruples)

		// Cargar constantes
		constantTable := parser.GetConstantTable()
		if constantTable != nil {
			constants := constantTable.GetConstants()
			virtualMachine.LoadConstants(constants)
		} else {
			virtualMachine.LoadConstants(make(map[int]interface{}))
		}

		// Ejecutar
		err := virtualMachine.Execute()

		// Mostrar output del programa
		programOutputStr := programOutput.String()
		if programOutputStr != "" {
			output.WriteString("SALIDA DEL PROGRAMA:\n")
			output.WriteString(programOutputStr)
			output.WriteString("\n")
		}

		// Mostrar debug si está habilitado
		if debug {
			debugOutputStr := debugOutput.String()
			if debugOutputStr != "" {
				output.WriteString("DEBUG INFO:\n")
				output.WriteString(debugOutputStr)
				output.WriteString("\n")
			}
		}

		if err != nil {
			output.WriteString(fmt.Sprintf("🌧 Error de ejecución: %v\n 🌧", err))
		} else {
			output.WriteString("⚡ Ejecución completada exitosamente!\n")
		}
	} else {
		output.WriteString("🏵️️ No hay cuádruplos para ejecutar 🏵️\n")
	}

	return output.String()
}
