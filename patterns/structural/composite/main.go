// Паттерн Composite (Компоновщик)
//
// Это структурный паттерн, который позволяет работать с одиночными объектами и их группами через единый интерфейс.
// Идея:
// - "Часть и целое имеют общий интерфейс"
// - Composite = рекурсивная структура + единый интерфейс.
// Принцип действия:
// - Строится дерево объектов
// - Клиент работает через интерфейс Component
// - Composite рекурсивно делегирует вызовы детям
// - Leaf выполняет реальную работу

package main

import "fmt"

// Component — общий контракт
// Любой объект в дереве:
// - может вернуть размер
// - может быть напечатан
// Клиенту становится не важно — файл это или папка.
type Component interface {
	GetName() string
	GetSize() int
	Print(ident string)
}

// File — Leaf, конечный элемент дерева.
// Не содержит детей
// Возвращает свой size
// Просто печатает себя
type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) Print(ident string) {
	fmt.Println(ident + f.name)
}

// Directory — Composite, просто вызывает метод интерфейса.
// - Хранит []Component
// - Делегирует вызовы детям
// - Работает рекурсивно
// - не знает: File это или ещё одна Directory
type Directory struct {
	name     string
	children []Component
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
	}
}

func (d *Directory) Add(component Component) {
	d.children = append(d.children, component)
}

func (d *Directory) Remove(component Component) {
	for i, child := range d.children {
		if child == component {
			d.children = append(d.children[:i], d.children[i+1:]...)
			break
		}
	}
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	totalSize := 0
	for _, child := range d.children {
		totalSize += child.GetSize()
	}
	return totalSize
}

func (d *Directory) Print(ident string) {
	fmt.Println(ident + d.name)
	for _, child := range d.children {
		child.Print(ident + "   ")
	}
}

func main() {
	file1 := NewFile("file1.txt", 100)
	file2 := NewFile("file2.txt", 200)

	subDir := NewDirectory("subdir")
	subDir.Add(NewFile("file3.txt", 300))

	subDir1 := NewDirectory("subdir_1")
	subDir1.Add(NewFile("file4.txt", 400))
	subDir1.Add(NewFile("file5.txt", 500))
	subDir.Add(subDir1)

	root := NewDirectory("root")
	root.Add(file1)
	root.Add(file2)
	root.Add(subDir)

	root.Print("")
	fmt.Println("Total size:", root.GetSize())
}

// Вывод:
/*
root
   file1.txt
   file2.txt
   subdir
      file3.txt
      subdir_1
         file4.txt
         file5.txt
Total size:
*/
