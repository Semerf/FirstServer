package calculate

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/semerf/FirstServer/internal/database"
)

const CommonRes int = 5

type TimeStamp struct {
	task     database.Task
	end_time int
}

func Calculator(tasks []database.Task) {
	fmt.Print("\n Calculator \n ", tasks, "\n")
	c := make(chan int, 10)
	min := 1000000
	beginTime := time.Now()
	for i := 0; i < 1000000; i++ {
		go Generate(tasks, c)
		val := <-c
		if val < min {
			min = val
		}
	}
	fmt.Println(min)
	endTime := time.Now()
	fmt.Println(endTime.Sub(beginTime))

}

func Generate(tasks []database.Task, c chan int) {
	tasksCopy := make([]database.Task, len(tasks))
	n := copy(tasksCopy, tasks)
	if n != len(tasks) {
		log.Fatal()
	}
	rand.Seed(time.Now().UnixNano())
	time_stamps := make([]TimeStamp, 0) //создаем массив штампов времен
	currentRes := CommonRes             //текущее количество ресурсов равно начальному
	currentTime := 0                    //начальное время - 0
	complete := make([]int, 0)          //выполненных задач нет

	for len(tasksCopy) != 0 {
		prevTaskDone := CanRealizeByPrevWork(tasksCopy, complete)
		taskReady := CanRealizeByRes(prevTaskDone, currentRes)
		for {
			if len(taskReady) == 0 {
				break
			}
			//выбрать рандомную задачу из доступных
			ind := rand.Intn(len(taskReady))
			task := taskReady[ind]                       //выбрали задачу
			taskReady[ind] = taskReady[len(taskReady)-1] //убрали задачу из доступных
			taskReady = taskReady[:len(taskReady)-1]

			time_stamps = append(time_stamps, TimeStamp{task, currentTime + task.Duration}) //добавили в массив штампов
			currentRes -= task.Resource                                                     //вычли ресурсы
			taskReady = CanRealizeByRes(taskReady, currentRes)
		}
		//переходим к следующему временному штампу
		minInd := 0
		if len(time_stamps) == 0 {
			currentTime = -1
			break
		}
		min := time_stamps[minInd]
		for ind, v := range time_stamps {
			if v.end_time < min.end_time {
				min = v
				minInd = ind
			}
		}
		currentTime = min.end_time
		currentRes += min.task.Resource
		complete = append(complete, min.task.Task_id)
		for ind, task := range tasksCopy {
			if task.Task_id == min.task.Task_id {
				tasksCopy[ind] = tasksCopy[len(tasksCopy)-1]
				tasksCopy = tasksCopy[:len(tasksCopy)-1]
				break
			}
		}
		time_stamps[minInd] = time_stamps[len(time_stamps)-1]
		time_stamps = time_stamps[:len(time_stamps)-1]

	}
	c <- currentTime

}

func CanRealizeByRes(tasks []database.Task, res int) []database.Task {
	readyTasks := make([]database.Task, 0) // срез готовых к выполнению работ
	for _, task := range tasks {           //для всех задач
		if task.Resource <= res {
			readyTasks = append(readyTasks, task)
		}
	}
	return readyTasks
}

func CanRealizeByPrevWork(tasks []database.Task, complete []int) []database.Task {
	readyTasks := make([]database.Task, 0) // срез готовых к выполнению работ
	for _, task := range tasks {           //для всех задач
		if len(task.Prev_work) == 0 { //если предшествующих работ нет добавляем к возможным для выполнения и продолжаем

			readyTasks = append(readyTasks, task)
			continue
		}
		elsInt := make([]int, 0) //преобразуем строку в int срез
		elsStr := strings.Split(task.Prev_work, ";")

		for _, val := range elsStr {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			elsInt = append(elsInt, intVal)
		}

		if subslice(elsInt, complete) { //если необходимые работы являются подмножеством выполненнных
			readyTasks = append(readyTasks, task)
		}
	}
	return readyTasks
}

func contains(s []int, e int) bool { //содержит ли срез s элемент e
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func subslice(i1 []int, i2 []int) bool { //является ли срез i1 подмножеством среза i2
	if len(i1) > len(i2) {
		return false
	}
	for _, e := range i1 {
		if !contains(i2, e) {
			return false
		}
	}
	return true
}
