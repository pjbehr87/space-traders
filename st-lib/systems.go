package stlib

import (
	"encoding/json"
	"fmt"
	"strings"
)

type System struct {
	Symbol       string
	SectorSymbol string `json:"sectorSymbol"`
	Type         string
	X            int64
	Y            int64
	Waypoints    []WaypointShort
	Factions     []Faction
}

type SystemData struct {
	System System `json:"data"`
}

type Waypoint struct {
	Symbol       string
	Type         string
	SystemSymbol string
	X            int64
	Y            int64
	Orbitals     []Orbital
	Faction      FactionShort
	Traits       []Trait
	Chart        Chart
}

type WaypointData struct {
	Waypoint []Waypoint `json:"data"`
}

type WaypointSymbol struct {
	Sector   string
	System   string
	Waypoint string
}

type WaypointShort struct {
	Symbol string
	Type   string
	X      int64
	Y      int64
}

type Orbital struct {
	Symbol string
}

type Chart struct {
	WaypointSymbol string
	SubmittedBy    string
	SubmittedOn    string
}

func getWps(wp string) WaypointSymbol {
	wps := strings.Split(wp, "-")

	return WaypointSymbol{
		Sector:   wps[0],
		System:   wps[0] + "-" + wps[1],
		Waypoint: wp,
	}
}

func (stl *StLib) GetWaypoint(system string, waypoint string) ([]Waypoint, error) {
	resp, err := stl.GetUrl(fmt.Sprintf("systems/%s/waypoints/%s", system, waypoint))
	if err != nil {
		return []Waypoint{}, err
	}

	waypointData := WaypointData{}
	fmt.Printf("%s", resp)
	err = json.Unmarshal(resp, &waypointData)
	if err != nil {
		return []Waypoint{}, err
	}
	fmt.Printf("%+v", waypointData)

	return waypointData.Waypoint, nil
}
