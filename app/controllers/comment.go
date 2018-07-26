package controllers

import (
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"webapp/app/models"
)

type CommentController struct {
	BaseController
}

func (c CommentController) Index() revel.Result {
	var (
		comments []models.Comment
		err      error
	)
	sort := c.GetParam("sort", "-created_at")
	limit, _ := strconv.Atoi(c.GetParam("limit", "10"))

	comments, err = models.GetComments(limit, sort)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(comments)
}

func (c CommentController) Show(id string) revel.Result {
	var (
		comment   models.Comment
		err       error
		commentID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid comment id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	commentID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid comment id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	comment, err = models.GetComment(commentID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(comment)
}

func (c CommentController) Create() revel.Result {
	var (
		comment models.Comment
		err     error
	)

	err = c.Params.BindJSON(&comment)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	comment, err = models.AddComment(comment)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(comment)
}

func (c CommentController) Update() revel.Result {
	var (
		comment models.Comment
		err     error
	)
	err = c.Params.BindJSON(&comment)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = comment.UpdateComment()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(comment)
}

func (c CommentController) Delete(id string) revel.Result {
	var (
		err       error
		comment   models.Comment
		commentID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid comment id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	commentID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid comment id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	comment, err = models.GetComment(commentID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = comment.DeleteComment()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}
