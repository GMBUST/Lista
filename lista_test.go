package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.VerPrimero()})
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.VerUltimo()})
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.BorrarPrimero()})
	require.Equal(t, 0, lista.Largo())
	
	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.VerPrimero()})
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.VerUltimo()})
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.BorrarPrimero()})
	require.Equal(t, 0, lista.Largo())
}

func TestUnicoValor(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Prueba insertando un elemento al inicio.
	lista.InsertarPrimero(20)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 20, lista.VerPrimero())
	require.Equal(t, 20, lista.VerUltimo())

	doble := 0
	lista.Iterar(func(x int) bool {
		doble = 2*x
		return true
	})

	require.Equal(t, 40, doble)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 20, lista.VerPrimero())
	require.Equal(t, 20, lista.VerUltimo())
	
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 20, iter.VerActual())

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.Equal(t, 20, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.VerPrimero()})
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.VerUltimo()})
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.BorrarPrimero()})
	require.Equal(t, 0, lista.Largo())
	
	iter = lista.Iterador()

	require.False(t, iter.HaySiguiente())

	// Prueba insertando un elemento al final.
	lista.InsertarUltimo(33)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 33, lista.VerPrimero())
	require.Equal(t, 33, lista.VerUltimo())

	doble = 0
	lista.Iterar(func(x int) bool {
		doble = 2*x
		return true
	})

	require.Equal(t, 66, doble)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 33, lista.VerPrimero())
	require.Equal(t, 33, lista.VerUltimo())
	
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 33, iter.VerActual())

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.Equal(t, 33, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.VerPrimero()})
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.VerUltimo()})
	require.PanicsWithValue(t, "La lista esta vacia", func() {lista.BorrarPrimero()})
	require.Equal(t, 0, lista.Largo())
	
	iter = lista.Iterador()

	require.False(t, iter.HaySiguiente())
}