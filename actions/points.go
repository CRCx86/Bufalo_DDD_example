package actions

import (
	"fmt"
	"location_service_v1/ls_v2/service"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/x/responder"
)

// PointsController is a
type PointsResource struct {
	buffalo.Resource
	pointsService *service.PointsService
}

// NewPointsController is a
func NewPointResource(service *service.PointsService) *PointsResource {
	return &PointsResource{
		pointsService: service,
	}
}

// List gets all Points. This function is mapped to the path
// GET /points
func (ctr PointsResource) List(c buffalo.Context) error {

	points, q, err := ctr.pointsService.List(c)
	if err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)
		c.Set("points", points)
		return c.Render(http.StatusOK, r.HTML("/points/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(points))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(points))
	}).Respond(c)
}

// Show gets the data for one Point. This function is mapped to
// the path GET /points/{point_id}
func (ctr PointsResource) Show(c buffalo.Context) error {

	point, err := ctr.pointsService.Show(c)
	if err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("point", point)

		return c.Render(http.StatusOK, r.HTML("/points/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(point))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(point))
	}).Respond(c)

}

// New renders the form for creating a new Point.
// This function is mapped to the path GET /points/new
func (ctr PointsResource) New(c buffalo.Context) error {
	point := ctr.pointsService.New(c)
	if point == nil {
		return fmt.Errorf("somthing goes worng")
	}
	c.Set("point", point)
	return c.Render(http.StatusOK, r.HTML("/points/new.plush.html"))
}

// Create adds a Point to the DB. This function is mapped to the
// path POST /points
func (ctr PointsResource) Create(c buffalo.Context) error {

	verrs, point, err := ctr.pointsService.Create(c)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("point", point)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/points/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "point.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/points/%v", point.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(point))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(point))
	}).Respond(c)
}

// Edit renders a edit form for a Point. This function is
// mapped to the path GET /points/{point_id}/edit
func (ctr PointsResource) Edit(c buffalo.Context) error {
	point, err := ctr.pointsService.Edit(c)
	if err != nil {
		return err
	}
	c.Set("point", point)
	return c.Render(http.StatusOK, r.HTML("/points/edit.plush.html"))
}

// Update changes a Point in the DB. This function is mapped to
// the path PUT /points/{point_id}
func (ctr PointsResource) Update(c buffalo.Context) error {

	verrs, point, err := ctr.pointsService.Update(c)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("point", point)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/points/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "point.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/points/%v", point.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(point))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(point))
	}).Respond(c)
}

// Destroy deletes a Point from the DB. This function is mapped
// to the path DELETE /points/{point_id}
func (ctr PointsResource) Destroy(c buffalo.Context) error {

	point, err := ctr.pointsService.Destroy(c)
	if err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "point.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/points")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(point))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(point))
	}).Respond(c)

}
