package stlib

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
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
	Symbol   *WaypointSymbol
	Type     string
	X        int64
	Y        int64
	Orbitals []Orbital
	Faction  FactionShort
	Traits   []Trait
	Chart    *Chart
}

type WaypointJson struct {
	Symbol       string
	Type         string
	SystemSymbol string
	X            int64
	Y            int64
	Orbitals     []OrbitalJson
	Faction      FactionShort
	Traits       []Trait
	Chart        ChartJson
}

type WaypointsData struct {
	Waypoints []WaypointJson `json:"data"`
}
type WaypointData struct {
	Waypoint WaypointJson `json:"data"`
}

type WaypointSymbol struct {
	Sector   string
	System   string
	Waypoint string
}

type WaypointShort struct {
	Symbol *WaypointSymbol
	Type   string
	X      int64
	Y      int64
}

type Orbital struct {
	Symbol *WaypointSymbol
}
type OrbitalJson struct {
	Symbol string
}

type ChartJson struct {
	WaypointSymbol string
	SubmittedBy    string
	SubmittedOn    string
}
type Chart struct {
	WaypointSymbol *WaypointSymbol
	SubmittedBy    string
	SubmittedOn    time.Time
}

func getWps(wp string) *WaypointSymbol {
	if wp == "" {
		return nil
	}
	wps := strings.Split(wp, "-")

	return &WaypointSymbol{
		Sector:   wps[0],
		System:   wps[0] + "-" + wps[1],
		Waypoint: wp,
	}
}

func (stl *StLib) GetWaypoint(system string, waypoint string) (Waypoint, error) {
	resp, err := stl.GetUrl(fmt.Sprintf("systems/%s/waypoints/%s", system, waypoint))
	if err != nil {
		return Waypoint{}, err
	}

	waypointData := WaypointData{}
	err = json.Unmarshal(resp, &waypointData)
	if err != nil {
		return Waypoint{}, err
	}

	wpj := waypointData.Waypoint
	wpjos := []Orbital{}
	for _, wpdo := range wpj.Orbitals {
		wpjos = append(wpjos, Orbital{
			Symbol: getWps(wpdo.Symbol),
		})
	}

	chartSubOn, err := time.Parse("2006-01-02T15:04:05Z07:00", wpj.Chart.SubmittedOn)
	if err != nil {
		return Waypoint{}, err
	}

	return Waypoint{
		Symbol:   getWps(wpj.Symbol),
		Type:     wpj.Type,
		X:        wpj.X,
		Y:        wpj.Y,
		Orbitals: wpjos,
		Faction: FactionShort{
			Symbol: wpj.Faction.Symbol,
		},
		Traits: wpj.Traits,
		Chart: &Chart{
			WaypointSymbol: getWps(wpj.Chart.WaypointSymbol),
			SubmittedBy:    wpj.Chart.SubmittedBy,
			SubmittedOn:    chartSubOn,
		},
	}, nil
}
func (stl *StLib) ListWaypoints(system string) ([]Waypoint, error) {
	resp, err := stl.GetUrl(fmt.Sprintf("systems/%s/waypoints", system))
	if err != nil {
		return []Waypoint{}, err
	}

	waypointsData := WaypointsData{}
	err = json.Unmarshal(resp, &waypointsData)
	if err != nil {
		return []Waypoint{}, err
	}

	waypoints := []Waypoint{}
	for _, wpd := range waypointsData.Waypoints {
		wpdos := []Orbital{}
		for _, wpdo := range wpd.Orbitals {
			wpdos = append(wpdos, Orbital{
				Symbol: getWps(wpdo.Symbol),
			})
		}

		chartSubOn, err := time.Parse("2006-01-02T15:04:05Z07:00", wpd.Chart.SubmittedOn)
		if err != nil {
			return []Waypoint{}, err
		}

		wp := Waypoint{
			Symbol:   getWps(wpd.Symbol),
			Type:     wpd.Type,
			X:        wpd.X,
			Y:        wpd.Y,
			Orbitals: wpdos,
			Faction: FactionShort{
				Symbol: wpd.Faction.Symbol,
			},
			Traits: wpd.Traits,
			Chart: &Chart{
				WaypointSymbol: getWps(wpd.Chart.WaypointSymbol),
				SubmittedBy:    wpd.Chart.SubmittedBy,
				SubmittedOn:    chartSubOn,
			},
		}

		waypoints = append(waypoints, wp)
	}
	return waypoints, nil
}
