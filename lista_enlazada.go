package lista

// Nodo de la lista:

// Implementación del nodo.
type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

// crearNodo crea un nodo que almacena datos de tipo T y retorna un puntero al mismo.
func crearNodo[T any](dato T) *nodo[T] {
	nuevoNodo := new(nodo[T])
	nuevoNodo.dato = dato
	return nuevoNodo
}


// Lista:

// Implementación de la lista.
type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

// CrearListaEnlazada crea e inicializa una lista enlazada
func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

// EstaVacia informa si la lista está vacía (true) o no (false).
func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil && lista.largo == 0
}

// InsertarPrimero inserta al inicio de la lista el elemento pasado por argumento.
func (lista *listaEnlazada[T]) InsertarPrimero(T) {

}

// InsertarUltimo inserta al final de la lista el elemento pasado por argumento.
func (lista *listaEnlazada[T]) InsertarUltimo(T) {

}

// BorrarPrimero elimina el primer elemento de la lista y retorna su valor.
func (lista *listaEnlazada[T]) BorrarPrimero() T {

}

// VerPrimero muestra el valor del primer elemento de la lista.
func (lista *listaEnlazada[T]) VerPrimero() T {
	return lista.primero.dato
}

// VerUltimo muestra el valor del último elemento de la lista.
func (lista *listaEnlazada[T]) VerUltimo() T {
	return lista.ultimo.dato
}

// Largo muestra el largo de la lista.
func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

// Iterar recorre internamente la lista y aplica una función 'visitar' la cual
// debe cumplir que: devuelve 'true' si desea seguir iterando y 'false' en caso contrario.
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {

}

// Iterador devuelve un elemento IteradorLista el cual es un iterador externo de la lista.
// Las primitivas asociadas al mismo se encuentran más adelante.
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {

}


// Iterador externo de la lista:

// Implementación del iterador.
type iteradorLista[T any] struct {
	anterior *nodo[T]
	actual   *nodo[T]
}

// VerActual devuelve el dato contenido en el elemento la lista en el que se encuentra el iterador.
func (iter *iteradorLista[T]) VerActual() T {

}

// HaySiguiente indica si hay algún elemento para ver.
func (iter *iteradorLista[T]) HaySiguiente() bool {

}

// Siguiente hace que el iterador avance al siguiente elemento de la lista.
func (iter *iteradorLista[T]) Siguiente() {

}

// Insertar inserta un elemento nuevo en la lista entre el elemento en el que se encuentra el iterador y el anterior a ese.
// Luego de insertarlo el iterador se posiciona sobre el elemento insertado.
func (iter *iteradorLista[T]) Insertar(T) {

}

// Borrar elimina el elemento de la lista sobre el que se encuentra el iterador.
// Luego de borrarlo el iterador se posiciona sobre el elemento siguiente al elemento que se borra de la lista.
func (iter *iteradorLista[T]) Borrar() T {

}
