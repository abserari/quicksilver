package main

import (
	"fmt"

	s2 "github.com/golang/geo/s2"
)

func main() {

	latlng := s2.LatLngFromDegrees(31.232135, 121.41321700000003)
	cellID := s2.CellIDFromLatLng(latlng)
	cell := s2.CellFromCellID(cellID) //9279882742634381312

	fmt.Println(cell)
	// cell.Level()
	fmt.Println("latlng = ", latlng)
	fmt.Println("cell level = ", cellID.Level())
	fmt.Printf("cell = %d\n", cellID)
	smallCell := s2.CellFromCellID(cellID.Parent(10))
	fmt.Printf("smallCell level = %d\n", smallCell.Level())
	fmt.Printf("smallCell id = %b\n", smallCell.ID())
	fmt.Printf("smallCell ApproxArea = %v\n", smallCell.ApproxArea())
	fmt.Printf("smallCell AverageArea = %v\n", smallCell.AverageArea())
	fmt.Printf("smallCell ExactArea = %v\n", smallCell.ExactArea())
}
