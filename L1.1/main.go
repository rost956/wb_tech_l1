// Package main демонстрирует встраивание методов через композицию в Go.
//
// Пакет показывает как структура Action может наследовать методы
// от родительской структуры Human с помощью embedded структур.
package main

import (
	"fmt"
	"strings"
)

// Human представляет человека с некоторыми данными - именем, возрастом, полом, весом
type Human struct {
	Name   string
	Age    int
	Gender string
	Weight int
}

// В Action встраиваются поля и становятся доступны методы структуры Human
type Action struct {
	Human
}

// Go_walk выводит информацию, когда человек вышел гулять
func (a Action) Go_walk() {
	if strings.Compare(a.Gender, "мальчик") == 0 {
		fmt.Printf("%v вышел погулять.\n", a.Name)
	} else {
		fmt.Printf("%v вышла погулять.\n", a.Name)
	}

}

// Was_born позволяет задать параметры для только что родившегося человека
func (h *Human) Was_born(name string, gender string, weight int) {
	fmt.Printf("На свет только что появился %v с весом %v грамм! Ура, это %v. \n", name, weight, gender)
	h.Gender = gender
	h.Name = name
	h.Age = 0
	h.Weight = weight
}

// main демонстирует встраивание полей и методов
func main() {
	human_petr := Action{}
	human_petr.Was_born("Петр", "мальчик", 3124)
	human_petr.Go_walk()
}
