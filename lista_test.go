package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_VOLUMEN = 1000
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())

	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
}

func TestUnicoValorInsercionPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Prueba insertando un elemento al inicio.
	lista.InsertarPrimero(20)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 20, lista.VerPrimero())
	require.Equal(t, 20, lista.VerUltimo())

	doble := 0
	lista.Iterar(func(x int) bool {
		doble = 2 * x
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
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())

	iter = lista.Iterador()

	require.False(t, iter.HaySiguiente())
}

func TestUnicoValorInsercionUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Prueba insertando un elemento al final.
	lista.InsertarUltimo(33)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 33, lista.VerPrimero())
	require.Equal(t, 33, lista.VerUltimo())

	doble := 0
	lista.Iterar(func(x int) bool {
		doble = 2 * x
		return true
	})

	require.Equal(t, 66, doble)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 33, lista.VerPrimero())
	require.Equal(t, 33, lista.VerUltimo())
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 33, iter.VerActual())

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.Equal(t, 33, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())

	iter = lista.Iterador()

	require.False(t, iter.HaySiguiente())
}

func TestMultiplesInsercionesUnicoValor(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarPrimero(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarUltimo(2)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 2, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarUltimo(3)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.Equal(t, 3, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarPrimero(4)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())
	require.Equal(t, 4, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarUltimo(5)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 5, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 5, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

}

func TestInsercionPrimeroPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	lista.InsertarPrimero(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	lista.InsertarPrimero(2)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestInsercionPrimeroUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	lista.InsertarPrimero(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	lista.InsertarUltimo(2)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestInsercionUltimoPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	lista.InsertarUltimo(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	lista.InsertarPrimero(2)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestInsercionUltimoUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	lista.InsertarUltimo(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	lista.InsertarUltimo(2)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func VerPrimeroYUltimo(t *testing.T) {
	t.Log("Insertamos Primero y Despues a lo ultimo y vemos si cumple con el orden deseado")
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())

	lista.InsertarPrimero(2)
	lista.InsertarUltimo(3)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
}

func TestInsertarPrimeroYBorrar(t *testing.T) {
	t.Log("Insertamos el dato en la primera posicion muchas veces, borramos y vemos si cumple con el orden deseado")
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 3, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())

}
func TestInsertarUltimoYBorrar(t *testing.T) {
	t.Log("Insertamos el dato en la Ultima posicion muchas veces, borramos y vemos si cumple con el orden desado")
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	lista.BorrarPrimero()

	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())

}

func TestVolumen(t *testing.T) {
	t.Log("Hacemos pruebas de volumen ")
	lista := TDALista.CrearListaEnlazada[int]()

	//Insertar primero
	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, i, lista.Largo())
		lista.InsertarPrimero(i)
		require.False(t, lista.EstaVacia())

		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, 0, lista.VerUltimo())

	}
	for i := 999; i >= 0; i-- {
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, i, lista.BorrarPrimero())
		require.Equal(t, i, lista.Largo())
	}
	// Verificar si se vacio correctamente
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	//Insertar ultimo
	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, i, lista.Largo())
		lista.InsertarUltimo(i)
		require.EqualValues(t, 0, lista.VerPrimero())
		require.EqualValues(t, i, lista.VerUltimo())

	}
	for i := 0; i < _VOLUMEN; i++ {
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
	// Verificar si se vacio correctamente
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}
func TestListaFloats(t *testing.T) {
	t.Log("Insertamos a la lista floats ")
	lista := TDALista.CrearListaEnlazada[float32]()
	lista.InsertarPrimero(1.13)
	require.EqualValues(t, 1.13, lista.VerPrimero())
	lista.InsertarPrimero(1.15)
	require.EqualValues(t, 1.15, lista.VerPrimero())
	require.EqualValues(t, 1.15, lista.BorrarPrimero())
	require.EqualValues(t, 1.13, lista.VerPrimero())
	require.EqualValues(t, 1.13, lista.BorrarPrimero())

	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
}
func TestListaStrings(t *testing.T) {
	t.Log("Insertamos a la lista strings")
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("asd")
	require.EqualValues(t, "asd", lista.VerPrimero())
	lista.InsertarPrimero("asd1")
	require.EqualValues(t, "asd1", lista.VerPrimero())

	require.EqualValues(t, "asd1", lista.BorrarPrimero())
	require.EqualValues(t, "asd", lista.VerPrimero())

	require.EqualValues(t, "asd", lista.BorrarPrimero())

	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
}

// ----Tests iterador externo----
func TestIteradorInsertarAlInicio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verifico condiciones iniciales de lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Verifico estar al inicio de la lista:
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	// Inserto elemento al inicio de la lista:
	iter.Insertar(55)
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 55, iter.VerActual())

	// Itero el resto de la lista:
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()

	// Verifico cambios en lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 4, lista.Largo())
	require.Equal(t, 55, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 55, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestIteradorInsertarAlFinal(t *testing.T) {
	t.Log("Inserto al final un elemento cuando el iterador esta al final")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verifico condiciones iniciales de lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Itero:
	iter := lista.Iterador()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())

	//Inserto al final del iterador
	iter.Insertar(7)
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 7, iter.VerActual())

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())

	// Verifico cambios en lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 4, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 7, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 7, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestIteradorInsertarEnElMedio(t *testing.T) {
	t.Log("Inserto en el medio y comprobamos que haya sido entre el anterior y el actual")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verifico condiciones iniciales de lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Itero hasta la mitad de la lista:
	iter := lista.Iterador()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())

	//Inserto en el medio
	iter.Insertar(100)
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 100, iter.VerActual())

	// Itero hasta el final:
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())

	// Verifico cambios en lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 4, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 100, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}
func TestRemoverElementoInicioIterador(t *testing.T) {
	t.Log("Remover elemento cuando el iterador es creado")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verificar condiciones iniciales de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Crear iterador y verificar estar en el inicio de la pila.
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	// Borrar elemento.
	iter.Borrar()

	// Iterar hasta el final de la lista.
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()

	// Verificar condiciones finales.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}
func TestRemoverElementoFinalIterador(t *testing.T) {
	t.Log("Remover elemento cuando el iterador es llega a su fin")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verificar condiciones iniciales de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Crear iterador e iterar hasta estar en el final de la lista (sobre el nodo que contiene el número 3).
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())

	// Borrar elemento.
	iter.Borrar()
	// Verificar que se haya borrado.
	require.False(t, iter.HaySiguiente())

	// Verificar condiciones finales de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}
func TestRemoverElementoDelMedio(t *testing.T) {
	t.Log("Remover elemento cuando esta en el medio y verificar que no esta")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verificar condiciones iniciales de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Crear iterador e iterar hasta estar en el medio de la lista (en el nodo que contiene el número 2).
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())

	// Borrar elemento.
	iter.Borrar()

	// Recorrer el resto de la lista.
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())

	// Verificar condiciones finales de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}
