package stlib

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

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

type Engine struct {
	Symbol       string
	Name         string
	Description  string
	Condition    int
	Speed        int
	Requirements ShipRequirements
}

type FrameJson struct {
	Symbol         string
	Name           string
	Description    string
	Condition      int
	ModuleSlots    int `json:"moduleSlots"`
	MountingPoints int `json:"mountingPoints"`
	FuelCapacity   int `json:"fuelCapacity"`
	Requirements   ShipRequirements
}
type Frame struct {
	Symbol         string
	Name           string
	Description    string
	Condition      int
	ModuleSlots    int
	MountingPoints int
	FuelCapacity   int
	Requirements   ShipRequirements
}

type Module struct {
	Symbol       string
	Capacity     int
	Range        int
	Name         string
	Description  string
	Requirements ShipRequirements
}

type Mount struct {
	Symbol       string
	Name         string
	Description  string
	Strength     int
	Deposits     []string
	Requirements ShipRequirements
}

type OrbitalJson struct {
	Symbol string
}
type Orbital struct {
	Symbol *WaypointSymbol
}

type Reactor struct {
	Symbol       string
	Name         string
	Description  string
	Condition    int
	PowerOutput  int `json:"powerOutput"`
	Requirements ShipRequirements
}

type ShipRequirements struct {
	Power int
	Crew  int
	Slots int
}

type ShipJson struct {
	// "type" is a reserved word
	ShipType      string `json:"type"`
	Name          string
	Description   string
	PurchasePrice int `json:"purchasePrice"`
	Frame         FrameJson
	Reactor       Reactor
	Engine        Engine
	Modules       []Module
	Mounts        []Mount
}
type Ship struct {
	// "type" is a reserved word
	ShipType      string
	Name          string
	Description   string
	PurchasePrice int
	Frame         Frame
	Reactor       Reactor
	Engine        Engine
	Modules       []Module
	Mounts        []Mount
}

type ShipTypeJson struct {
	Type string
}

type ShipyardJson struct {
	Symbol       string
	ShipTypes    []ShipTypeJson `json:"shipTypes"`
	Transactions []TransactionJson
	Ships        []ShipJson
}
type Shipyard struct {
	Symbol       string
	ShipTypes    []string
	Transactions []Transaction
	Ships        []Ship
}

type System struct {
	Symbol       string
	SectorSymbol string `json:"sectorSymbol"`
	Type         string
	X            int
	Y            int
	Waypoints    []WaypointShort
	Factions     []Faction
}

type TransactionJson struct {
	WaypointSymbol string `json:"waypointSymbol"`
	ShipSymbol     string
	Price          int
	AgentSymbol    string `json:"agentSymbol"`
	Timestamp      string
}
type Transaction struct {
	WaypointSymbol *WaypointSymbol
	ShipSymbol     string
	Price          int
	AgentSymbol    string
	Timestamp      time.Time
}

type WaypointJson struct {
	Symbol       string
	Type         string
	SystemSymbol string
	X            int
	Y            int
	Orbitals     []OrbitalJson
	Faction      FactionShort
	Traits       []Trait
	Chart        ChartJson
}
type Waypoint struct {
	Symbol   *WaypointSymbol
	Type     string
	X        int
	Y        int
	Orbitals []Orbital
	Faction  FactionShort
	Traits   []Trait
	Chart    *Chart
}
type WaypointShort struct {
	Symbol *WaypointSymbol
	Type   string
	X      int
	Y      int
}
type WaypointSymbol struct {
	Sector   string
	System   string
	Waypoint string
}

type ShipyardData struct {
	Shipyard ShipyardJson `json:"data"`
}
type SystemData struct {
	System System `json:"data"`
}
type WaypointsData struct {
	Waypoints []WaypointJson `json:"data"`
}
type WaypointData struct {
	Waypoint WaypointJson `json:"data"`
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

func (stl *StLib) GetShipyard(system string, waypoint string) (Shipyard, error) {
	resp, err := stl.GetUrl(fmt.Sprintf("systems/%s/waypoints/%s/shipyard", system, waypoint))
	if err != nil {
		return Shipyard{}, err
	}

	shipyardData := ShipyardData{}
	err = json.Unmarshal(resp, &shipyardData)
	if err != nil {
		return Shipyard{}, err
	}

	sdj := shipyardData.Shipyard

	shipTypes := []string{}
	for _, sdst := range shipyardData.Shipyard.ShipTypes {
		shipTypes = append(shipTypes, sdst.Type)
	}
	transactions := []Transaction{}
	for _, sdstj := range shipyardData.Shipyard.Transactions {
		timestamp, err := time.Parse("2006-01-02T15:04:05Z07:00", sdstj.Timestamp)
		if err != nil {
			return Shipyard{}, err
		}
		transaction := Transaction{
			WaypointSymbol: getWps(sdstj.WaypointSymbol),
			ShipSymbol:     sdstj.ShipSymbol,
			Price:          sdstj.Price,
			AgentSymbol:    sdstj.AgentSymbol,
			Timestamp:      timestamp,
		}

		transactions = append(transactions, transaction)
	}
	ships := []Ship{}
	for _, sdsj := range shipyardData.Shipyard.Ships {
		ships = append(ships, Ship{
			ShipType:      sdsj.ShipType,
			Name:          sdsj.Name,
			Description:   sdsj.Description,
			PurchasePrice: sdsj.PurchasePrice,
			Frame: Frame{
				Symbol:         sdsj.Frame.Symbol,
				Name:           sdsj.Frame.Name,
				Description:    sdsj.Frame.Description,
				Condition:      sdsj.Frame.Condition,
				ModuleSlots:    sdsj.Frame.ModuleSlots,
				MountingPoints: sdsj.Frame.MountingPoints,
				FuelCapacity:   sdsj.Frame.FuelCapacity,
				Requirements:   sdsj.Frame.Requirements,
			},
			Reactor: sdsj.Reactor,
			Engine:  sdsj.Engine,
			Modules: sdsj.Modules,
			Mounts:  sdsj.Mounts,
		})
	}
	return Shipyard{
		Symbol:       sdj.Symbol,
		ShipTypes:    shipTypes,
		Transactions: transactions,
		Ships:        ships,
	}, nil
}
