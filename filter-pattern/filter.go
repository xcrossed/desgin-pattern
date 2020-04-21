package main

import (
	"fmt"
)

type Filter interface {
	DoFilter(fc FilterChain, val *RequestObj) error
}

type FilterChain interface {
	AddFilter(filter Filter)
	DoFilter(*RequestObj) error
}

type filterChain struct {
	Filters []Filter
	next    int
}

func NewFilterChain() FilterChain {
	return &filterChain{
		Filters: make([]Filter, 0, 10),
		next:    0,
	}
}

func (fc *filterChain) AddFilter(filter Filter) {
	fc.Filters = append(fc.Filters, filter)
}

func (fc *filterChain) DoFilter(requestObj *RequestObj) error {
	if fc.next < len(fc.Filters) {
		cuIdx := fc.next
		fc.next = fc.next + 1
		fc.Filters[cuIdx].DoFilter(fc, requestObj)
	}
	return nil
}

type RequestObj struct {
	Name string
	Age  int
	Sex  int
	Desc string
}

type AgeFilter struct{}

func (ageFilter *AgeFilter) DoFilter(fc FilterChain, val *RequestObj) error {
	if val.Age > 10 && val.Age < 17 {
		val.Desc = "少年"
	} else if val.Age >= 17 {
		val.Desc += "成年人 "
	} else {
		val.Desc += "儿童 "
	}
	fmt.Println("AgeFilter req")
	fc.DoFilter(val)
	fmt.Println("AgeFilter resp")
	return nil
}

type SexFilter struct{}

func (sexFilter *SexFilter) DoFilter(fc FilterChain, val *RequestObj) error {
	if val.Sex == 1 {
		val.Desc += "女 "
	} else if val.Sex == 2 {
		val.Desc += "男 "
	} else {
		val.Desc += "未知 "
	}
	fmt.Println("SexFilter req")
	fc.DoFilter(val)
	fmt.Println("SexFilter resp")
	return nil
}

func main() {
	filterChain := NewFilterChain()

	filterChain.AddFilter(new(AgeFilter))
	filterChain.AddFilter(new(SexFilter))
	reqObj := &RequestObj{
		Name: "Tom",
		Age:  10,
		Sex:  1,
	}
	fmt.Printf("before filter:%#v\n", reqObj)
	filterChain.DoFilter(reqObj)
	fmt.Printf("after filter:%#v\n", reqObj)

	fmt.Println("hello")
}
