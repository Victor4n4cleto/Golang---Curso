// jogo copiado do site oficial da linguagem GO:
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// Field representa um campo bidimensional de células.
type Field struct {
	s    [][]bool
	w, h int
}

// NewField retorna um campo vazio de largura e altura especificadas.
func NewField(w, h int) *Field {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &Field{s: s, w: w, h: h}
}

// Set define o estado da célula especificada para o valor fornecido.
func (f *Field) Set(x, y int, b bool) {
	f.s[y][x] = b
}

// Alive informa se a célula especificada está ativa.
// If the x or y as coordenadas estão fora dos limites do campo elas são encapsuladas
// toroidally =  uma superfície de revolução com um orifício no meio. Por exemplo, um valor x de -1 é tratado como largura-1.
func (f *Field) Alive(x, y int) bool {
	x += f.w
	x %= f.w
	y += f.h
	y %= f.h
	return f.s[y][x]
}

// Next retorna o estado da célula especificada na próxima etapa de tempo.
func (f *Field) Next(x, y int) bool {
	// Conte as células adjacentes que estão alive.
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.Alive(x+i, y+j) {
				alive++
			}
		}
	}
	// Retorna o próximo estado de acordo com as regras do jogo:
	//   exatamente 3 vizinhos : on,
	//   exatamente 2 vizinhos: manter o estado atual,
	//   por outro lado: off.
	return alive == 3 || alive == 2 && f.Alive(x, y)
}

// A vida armazena o estado de uma rodada de Conway's Game of Life.
type Life struct {
	a, b *Field
	w, h int
}

// NewLife returns a novo estado de jogo de vida com um estado inicial aleatório.
func NewLife(w, h int) *Life {
	a := NewField(w, h)
	for i := 0; i < (w * h / 4); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &Life{
		a: a, b: NewField(w, h),
		w: w, h: h,
	}
}

// Step avança o jogo em um instante, recalculando e atualizando todas as células.
func (l *Life) Step() {
	// Atualizar o estado do next field (b) do campo atual (a).
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}
	// Trocar campos a and b.
	l.a, l.b = l.b, l.a
}

// String devolve o tabuleiro de jogo como string.
func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte(' ')
			if l.a.Alive(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func main() {
	l := NewLife(40, 15)
	for i := 0; i < 300; i++ {
		l.Step()
		fmt.Print("\x0c", l) // Limpar tela e campo de impressão.
		time.Sleep(time.Second / 30)
	}
}
