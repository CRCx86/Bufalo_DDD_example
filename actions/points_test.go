package actions

import (
	"location_service_v1/ls_v2/models"
)

func (as *ActionSuite) Test_PointsResource_List() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_PointsResource_Show() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_PointsResource_Create() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_PointsResource_Update() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_PointsResource_Destroy() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_PointsResource_New() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_PointsResource_Edit() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_GetPickPointsList() {
	pointsDB := make([]*models.Point, 0)
	point := models.Point{
		Name: "1",
	}
	pointsDB = append(pointsDB, &point)
	println(pointsDB)
	println(1)
}
