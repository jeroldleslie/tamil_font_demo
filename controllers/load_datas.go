package controllers

import (
	"log"
	//"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"regexp"
	"encoding/csv"
	"tamil_font_demo/models"
	"sync"
)

var schools = make([]models.School, 0)
var homeDir = ""

type datas struct {
	Schools []models.School
}

var instance *datas
var once sync.Once

func GetDatas() *datas {
	log.Printf("inside singleton")
	once.Do(func() {
		LoadDefaultDatas()
		instance = &datas{
			Schools: schools,
		}
	})
	return instance
}

func LoadDefaultDatas() {
	log.Printf("Loading default datas")
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeDir = usr.HomeDir
	parseCsv()
	picParse()
	for _, school := range schools {
		print(school.Images)
	}
}

func isDataLine(line []string) bool {
	re := regexp.MustCompile("^[A-Z]+ [A-Z]+/[A-Z]*$")
	return (!re.Match([]byte(line[1])))
}

func isGpaLine(line []string) bool {
	re := regexp.MustCompile("^[0-9][.][0-9]$")
	return re.Match([]byte(line[1]))
}

func parseGpaLine(line []string) models.GpaReq {
	gpa, _ := strconv.ParseFloat(line[1], 64)
	act, _ := strconv.Atoi(line[2])
	sat, _ := strconv.Atoi(line[3])
	return models.GpaReq{Gpa: gpa, Act: act, Sat: sat}
}

func parsePercLine(line []string) models.PercReq {
	perc, _ := strconv.Atoi(line[1])
	act, _ := strconv.Atoi(line[2])
	sat, _ := strconv.Atoi(line[3])
	return models.PercReq{Perc: perc, Act: act, Sat: sat}
}

func parseName(line []string) string {
	return line[0]
}

func parseCityLine(line []string) (string, float64, float64) {
	lat, _ := strconv.ParseFloat(line[4], 64)
	lng, _ := strconv.ParseFloat(line[5], 64)
	return line[0], lat, lng
}

func parseSchool(r *csv.Reader, name string, id int) (*models.School, string) {
	line, err := r.Read()
	city, lat, lng := parseCityLine(line)
	gpas := make([]models.GpaReq, 0)
	percs := make([]models.PercReq, 0)
	for {
		if isDataLine(line) {
			if isGpaLine(line) {
				gpas = append(gpas, parseGpaLine(line))
			} else {
				percs = append(percs, parsePercLine(line))
			}
		} else {
			break
		}
		line, err = r.Read()
		//more serious error checking and our own error definition are required
		if err != nil {
			return &models.School{Name: name, City: city, Lat: lat, Lng: lng, Id: id, Percs: percs, Gpas: gpas, Images: make([]string, 0)}, ""
		}
	}
	// if err != nil {
	// 	println(err)
	// }
	return &models.School{Name: name, City: city, Lat: lat, Lng: lng, Id: id, Percs: percs, Gpas: gpas, Images: make([]string, 0)}, parseName(line)
}

func parseCsv() {
	path := "src/tamil_font_demo/resources/datas/data.csv"
	firstDataCsv, _ := os.Open(path)
	r := csv.NewReader(firstDataCsv)

	_, err := r.Read()
	_, err = r.Read()
	line, err := r.Read()
	var name *string = &line[0]
	*name = parseName(line)
	id := 0
	for {
		var school *models.School
		school, *name = parseSchool(r, *name, id)
		schools = append(schools, *school)
		id++
		if *name == "" {
			break
		}
	}
	if err != nil {
		println("WE ARE HERE")
		log.Fatal(err)
	}
}

func picParse() {
	//path, _ := filepath.Abs("freshdb.csv")
	data, _ := os.Open(homeDir + "/.tamil_font_demo/datas/freshdb.csv")
	r := csv.NewReader(data)
	lineCounter := -2
	for {
		line, err := r.Read()
		if err != nil {
			break
		}
		if line[13] == "" {
			continue
		}
		if line[13][0] == 'P' {
			lineCounter++
		}
		if lineCounter == 49 {
			break
		}
		if []rune(line[13])[0] == 'h' {
			schools[lineCounter].Images = append(schools[lineCounter].Images, line[13])
			schools[lineCounter].Images = append(schools[lineCounter].Images, line[14])

			schools[lineCounter].AcceptanceRate = line[7]
			schools[lineCounter].TotalStudents = line[8]
			
			
			male_female_ratio := strings.Split(line[9],"/")
			
			if len(male_female_ratio) == 2 {
				schools[lineCounter].MaleRatio = male_female_ratio[0]
				schools[lineCounter].FemaleRatio = male_female_ratio[1]
			}

			schools[lineCounter].StudentPopulation = line[10]
			schools[lineCounter].FreshmanClassSize = line[11]
			schools[lineCounter].AvgFinancialAidPackage = line[12]
		}
	}
}
