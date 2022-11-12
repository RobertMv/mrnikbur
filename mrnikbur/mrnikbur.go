package mrnikbur

import (
	"fmt"
	"ssl-test/mrnikbur/persons"
)

type Person string

// MrNikBur represents mrnikbur's package main structure
type MrNikBur struct {
	height      string
	hair        string
	isKind      bool
	isHot       bool
	orientation string
	dick        string
	pants       Pants
}

// New initializes new MrNikBur instance
func New() *MrNikBur {
	return &MrNikBur{
		height:      "башой",
		hair:        "как у рэпера",
		isKind:      true,
		isHot:       true,
		orientation: "straight",
		dick:        "is nice, bro",
	}
}

// Joy method is used to express joy and happiness
func (m *MrNikBur) Joy() {
	fmt.Print("AХАХАХАХАХА")
}

// Amaze is used to express amazement for some events or news
func (m *MrNikBur) Amaze() {
	fmt.Print("Пипяяяяу")
}

// TalkTo is used for talking and negotiation
func (m *MrNikBur) TalkTo(person *Person) {
	if person == nil {
		*person = persons.Zheltock
	}
	fmt.Printf("Значит разговаривал с %s, он сказал...", *person)
}

// AtPool is used for
func (m *MrNikBur) AtPool() {
	m.pants.off()
	m.orientation = "gay"
}

type Pants struct {
}

func (p Pants) on() {
	fmt.Print("Pants are on!")
}

func (p Pants) off() {
	fmt.Print("Pants are off")
}
