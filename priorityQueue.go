package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	nombre    string
	prioridad int
	indice    int
}
type PriorityQueue []*Item

func (cola PriorityQueue) Len() int { return len(cola) }

func (cola PriorityQueue) Less(i, j int) bool {

	return cola[i].prioridad > cola[j].prioridad
}

func (cola PriorityQueue) Swap(i, j int) {
	cola[i], cola[j] = cola[j], cola[i]
	cola[i].indice = i
	cola[j].indice = j
}

func (cola *PriorityQueue) Push(x any) {
	n := len(*cola)
	item := x.(*Item)
	item.indice = n
	*cola = append(*cola, item)
}

func (cola *PriorityQueue) Pop() any {
	old := *cola
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.indice = -1
	*cola = old[0 : n-1]
	return item
}

func (cola *PriorityQueue) update(item *Item, nombre string, prioridad int) {
	item.nombre = nombre
	item.prioridad = prioridad
	heap.Fix(cola, item.indice)
}

func main() {

	items := map[string]int{
		"Perro": 3, "Gato": 2, "Oso": 4,
	}

	cola := make(PriorityQueue, len(items))
	i := 0
	for nombre, prioridad := range items {
		cola[i] = &Item{
			nombre:    nombre,
			prioridad: prioridad,
			indice:    i,
		}
		i++
	}
	heap.Init(&cola)

	item := &Item{
		nombre:    "Zorro",
		prioridad: 1,
	}
	heap.Push(&cola, item)
	cola.update(item, item.nombre, 5)

	for cola.Len() > 0 {
		item := heap.Pop(&cola).(*Item)
		fmt.Printf("%.2d:%s ", item.prioridad, item.nombre)
	}
}
