package models

import ()

type School struct {
	Name   string    `json:"name"`
	City   string    `json:"city"`
	Lat    float64   `json:"lat"`
	Lng    float64   `json:"lng"`
	Id     int       `json:"id"`
	Percs  []PercReq `json:"percs"`
	Gpas   []GpaReq  `json: "gpas"`
	Images []string  `json:"images"`
	AcceptanceRate string `json:"acceptance_rate"`
	TotalStudents string `json:"total_students"`
	MaleRatio string `json:"male_ratio"`
	FemaleRatio string `json:"female_ratio"`
	StudentPopulation string `json:"student_population"`
	FreshmanClassSize  string `json:"freshman_class_size"`
	AvgFinancialAidPackage string `json:"avg_financial_aid_package"`
}

type GpaReq struct {
	Gpa float64 `json:"gpa"`
	Act int     `json:"act"`
	Sat int     `json:"sat"`
}

type PercReq struct {
	Perc int `json:"perc"`
	Act  int `json:"act"`
	Sat  int `json:"sat"`
}

type Stats struct {
	Perc int
	Gpa  float64
	Sat  int
	Act  int
}

type req interface {
	Decide(stats *Stats)
}

func (g *GpaReq) Decide(stats *Stats) bool {
	if g.Gpa <= stats.Gpa && ((g.Act <= stats.Act) || (g.Sat <= stats.Sat)) {
		return true
	} else {
		return false
	}
}

func (p *PercReq) Decide(stats *Stats) bool {
	if stats.Perc <= p.Perc && ((p.Act <= stats.Act) || (p.Sat <= stats.Sat)) {
		return true
	} else {
		return false
	}
}
