package parse

import (
	"bufio"
	"challengeFinal/models"
	"os"
	"strings"
	"fmt"
)

func LoadData(filePath string) ([]models.Record, error) {
    var records []models.Record 

    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Split(line, ", ")

        if len(fields) < 14 { 
            continue
        }

        
        record := models.Record{
            Age:           parseInt(fields[0]),              
            Workclass:     models.ParseWorkclass(fields[1]), 
            Fnlwgt:        parseInt(fields[2]),              
            Education:     fields[3],                        
            EducationNum:  parseInt(fields[4]),              
            MaritalStatus: fields[5],                       
            Occupation:    fields[6],                       
            Relationship:  fields[7],                   
            Race:          fields[8],                       
            Sex:           fields[9], 
            CapitalGain:   parseInt(fields[10]), 
	        CapitalLoss:   parseInt(fields[11]),    
	        HoursPerWeek:  parseInt(fields[12]), 
	        NativeCountry: fields[13], 
	        Income:        fields[14],                      
        }

        records = append(records, record)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return records, nil
}


func parseInt(value string) int {
    var result int
    fmt.Sscanf(value, "%d", &result)
    return result
}
