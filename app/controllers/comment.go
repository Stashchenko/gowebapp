package controllers

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"webapp/app/models"
)

type CommentController struct {
	BaseController
}

const InvalidIdMessage = "Invalid comment id format"

func (c CommentController) Index() revel.Result {
	var (
		comments []models.Comment
		err      error
	)

	sort := c.GetParam("sort", "-created_at")
	limit, _ := strconv.Atoi(c.GetParam("limit", "10"))

	comments, err = models.GetComments(limit, sort)
	if err != nil {
		return c.buildError(err.Error(), 500)
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
	commentID, err = convertToObjectIdHex(id)
	if id == "" || err != nil {
		return c.buildError(InvalidIdMessage, 400)
	}

	comment, err = models.GetComment(commentID)
	if err != nil {
		return c.buildError(err.Error(), 500)
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
		return c.buildError(err.Error(), 403)
	}

	comment, err = models.AddComment(comment)
	if err != nil {
		return c.buildError(err.Error(), 500)
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
		return c.buildError(err.Error(), 400)
	}

	err = comment.UpdateComment()
	if err != nil {
		return c.buildError(err.Error(), 500)
	}
	return c.RenderJSON(comment)
}

func (c CommentController) Delete(id string) revel.Result {
	var (
		err       error
		comment   models.Comment
		commentID bson.ObjectId
	)
	commentID, err = convertToObjectIdHex(id)
	if id == "" || err != nil {
		return c.buildError(InvalidIdMessage, 400)
	}
	comment, err = models.GetComment(commentID)
	if err != nil {
		return c.buildError(err.Error(), 500)
	}
	err = comment.DeleteComment()
	if err != nil {
		return c.buildError(err.Error(), 500)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}
