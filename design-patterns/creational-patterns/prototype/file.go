package main

import "fmt"

type File struct {
	name string
}

func (f *File) print(indetation string) {
	fmt.Println(indetation + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}
