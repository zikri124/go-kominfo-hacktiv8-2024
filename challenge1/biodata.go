package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/zikri124/go-kominfo-2024/challenge1/db"
)

func main() {
	studentDb := db.StudentDb{}
	initDbData(&studentDb)

	id := os.Args

	if len(id) < 2 {
		err := errors.New("error: user Id not provided, please insert user Id")
		fmt.Println(err.Error())
	} else {
		userId, _ := strconv.Atoi(id[1])
		printStudentById(&studentDb, userId)
	}
}

func initDbData(studentDb *db.StudentDb) {
	rawDatas := [][]string{
		{"MULIA HARTAWAN NEGARA", "Jakarta", "Batch 1 Golang Participant", "Ingin mendalami back end enginering"},
		{"DZULFIQAR AKMALUDDIN", "Jakarta", "Batch 1 Golang Participant", "Ingin mendalami back end enginering"},
		{"RIEFKY ARIF IBRAHIM", "Jakarta", "Batch 1 Golang Participant", "Ingin mendalami back end enginering"},
		{"Permadi Hidayat", "Jakarta", "Batch 1 Golang Participant", "Ingin mendalami back end enginering"},
		{"Talitha Alda Zafirah Dewi", "Jakarta", "Batch 1 Golang Participant", "Ingin mendalami back end enginering"},
		{"ZIKRI KURNIA AIZET", "Padang", "Batch 1 Golang Participant", "Ingin mendalami back end enginering"},
		{"Fungki Prasetya", "Jakarta", "Batch 1 Golang Participant", "Ingin mendalami back end enginering"},
		{"SYAMSUL TRI ANDIKA", "Jakarta", "Batch 1 Golang Participant", "Ingin mendalami back end enginering"},
		{"ALFUL LAILA S", "Jakarta", "Batch 1 Golang Participant", "Ingin mendalami back end enginering"},
		{"SALMIA RAHMAWATI", "Jakarta", "Batch 1 Golang Participant", "Ingin mendalami back end enginering"},
	}

	for _, data := range rawDatas {
		student := db.Student{Name: data[0], Address: data[1], Work: data[2], JoinReason: data[3]}
		studentDb.AddStudent(student)
	}
}

func printStudentById(studentDb *db.StudentDb, id int) {
	student, err := studentDb.GetStudentById(id)

	if err == nil {
		student.PrintStudent()
	} else {
		fmt.Println(err.Error())
	}
}
