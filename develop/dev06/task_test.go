package main

import (
	"os"
	"testing"
)

func TestCut(t *testing.T) {
	type Args struct {
		separatedPtr bool
		delimiterPtr string
		fieldsPtr    string
	}

	tests := []struct {
		name             string
		args             Args
		input            string
		expectedOutput   string
		expectedError    bool
		expectedPanic    bool
		expectedPanicMsg string
	}{
		{
			name: "No flags, single line",
			args: Args{
				false, "\t", "",
			},
			input:          "Hello\tWorld\n",
			expectedOutput: "Hello\tWorld\n",
			expectedError:  false,
			expectedPanic:  false,
		},
		{
			name: "No flags, multiple lines",
			args: Args{
				false, "\t", "",
			},
			input:          "Hello\tWorld\nFoo\tBar\n",
			expectedOutput: "Hello\tWorld\nFoo\tBar\n",
			expectedError:  false,
			expectedPanic:  false,
		},
		{
			name: "Select fields, single line",
			args: Args{
				false, "\t", "1",
			},
			input:          "Hello\tWorld\tFoo",
			expectedOutput: "Hello\n",
			expectedError:  false,
			expectedPanic:  false,
		},
		{
			name: "Select fields, multiple lines",
			args: Args{
				false, "\t", "2,3",
			},
			input:          "Hello\tWorld\tFoo\n1\t2\t3",
			expectedOutput: "World\tFoo\n2\t3\n",
			expectedError:  false,
			expectedPanic:  false,
		},
		{
			name: "Invalid field index",
			args: Args{
				false, "\t", "4",
			},
			input:            "Hello\tWorld\tFoo",
			expectedOutput:   "",
			expectedError:    false,
			expectedPanic:    false,
			expectedPanicMsg: "runtime error: index out of range",
		},
		{
			name: "Separated flag, single line without delimiter",
			args: Args{
				true, "\t", "",
			},
			input:          "Hello World",
			expectedOutput: "",
			expectedError:  false,
			expectedPanic:  false,
		},
		{
			name: "Separated flag, single line with delimiter",
			args: Args{
				true, ",", "",
			},
			input:          "Hello,World",
			expectedOutput: "Hello,World\n",
			expectedError:  false,
			expectedPanic:  false,
		},
		{
			name: "Separated flag, multiple lines with/without delimiter",
			args: Args{
				true, ",", "",
			},
			input:          "Hello,World\nFoo,Bar\n1,2,3\nTest",
			expectedOutput: "Hello,World\nFoo,Bar\n1,2,3\n",
			expectedError:  false,
			expectedPanic:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Пайп для эмуляции os.Stdin
			stdinReader, stdinWriter, err := os.Pipe()
			if err != nil {
				t.Error("failed to create pipe: ", err)
			}

			// Пайп для эмуляции os.Stdout
			stdoutReader, stdoutWriter, err := os.Pipe()
			if err != nil {
				t.Error("failed to create pipe: ", err)
			}

			// Запись строки теста в фукнцию
			stdinWriter.Write([]byte(test.input))
			stdinWriter.Close()

			// Переопределение Stdin и Stdout
			oldStdin := os.Stdin
			os.Stdin = stdinReader

			oldStdout := os.Stdout
			os.Stdout = stdoutWriter

			if test.expectedPanic {
				defer func() {
					if r := recover(); r != nil {
						if r != test.expectedPanicMsg {
							t.Errorf("unexpected panic message: got %v, want %v", r, test.expectedPanicMsg)
						}
					} else {
						t.Error("expected panic, but no panic occurred")
					}
				}()
			}

			Cut(test.args.separatedPtr, test.args.delimiterPtr, test.args.fieldsPtr)

			stdinReader.Close()
			stdoutWriter.Close()

			// Чтение результата функции
			buffer := make([]byte, len(test.expectedOutput))
			stdoutReader.Read(buffer)
			stdoutReader.Close()
			output := string(buffer)

			// Восстановление Stdin и Stdout
			os.Stdin = oldStdin
			os.Stdout = oldStdout

			if output != test.expectedOutput {
				t.Errorf("unexpected output: got %v, want %v", output, test.expectedOutput)
			}
		})
	}
}
