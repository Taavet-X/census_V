package loaders

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"censusV/database"
	"censusV/models"
)

// Loads the data from a file and returns a string
func ReadData(filePath string) list.List {
    // Open the file
    rows := list.New()
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening the file: ", err)
        return *rows
    }
    defer file.Close()

    // Read the first line of the file
    scanner := bufio.NewScanner(file)
    if scanner.Scan() {
        firstLine := scanner.Text()
        fmt.Println("First line of the file: ", firstLine)
    }

    // Read the rest of the file    
    
    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, ",")
        rows.PushBack(split)
    }

    return *rows
}

func SaveDataIntoDB (data list.List) error {

    var wg sync.WaitGroup
    db := database.DB
    
    for e := data.Front(); e != nil; e = e.Next() {
        fields := e.Value.([]string)
        fmt.Println(fields)
        // Convertir los valores necesarios

        age, err := strconv.Atoi(strings.TrimSpace(fields[0]))
        if err != nil {
            fmt.Printf("Error al convertir la edad: %v para la línea: %s\n", err)
            continue
        }

        hours, err := strconv.Atoi(strings.TrimSpace(fields[12]))
        if err != nil {
            fmt.Printf("Error al convertir las horas: %v para la línea: %s\n", err)
            continue
        }
        data := models.CensusData{
            Age:           age,
            Workclass:     strings.TrimSpace(fields[1]),
            Education:     strings.TrimSpace(fields[3]),
            MaritalStatus: strings.TrimSpace(fields[5]),
            Occupation:    strings.TrimSpace(fields[6]),
            Relationship:  strings.TrimSpace(fields[7]),
            Race:          strings.TrimSpace(fields[8]),
            Sex:           strings.TrimSpace(fields[9]),
            HoursPerWeek:  hours,
            Income:        strings.TrimSpace(fields[14]),
        }

        wg.Add(1)
        go func(data models.CensusData) {
            defer wg.Done()
            log.Println(data)
            if err := db.Create(&data).Error; err != nil {
                fmt.Printf("Error al guardar los datos en la base: %v\n", err)
            }
        }(data)
    }
    wg.Wait()
    return nil
}