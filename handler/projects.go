package handler

import (
	"main/model"
	"main/model/db"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Getprojects godoc
//
//	 TODO
//	@Summary		Get one's projects
//	@Tags			dev
//	@Description	Get all the projects from user(login required)
//	@Produce		json
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	db.ProposalInfo
//	@Router			/users/myproject [get]
func Getprojects(r *gin.Context) {
	id := r.GetInt("userID")
	data, _ := model.GetProposals(id)
	r.JSON(200, data)
}

// Getproject godoc
//
//	@Summary		Get a project
//	@Description	Get a project with its id
//	@Param			info_id			query	string	true	"the id of the project"
//	@Param			Authorization	header	string	true	"token"
//	@Produce		json
//	@Success		200	{object}	db.ProposalInfo
//	@Failure		404	{object}	handler.Response
//	@Router			/project [get]
func Getproject(r *gin.Context) {
	q := r.Query("info_id")
	id, err := strconv.Atoi(q)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusNotFound)
		return
	}
	data := db.ProposalInfo{
		InfoID: int32(id),
	}
	data = model.GetSth(data)
	SendResponse(r, nil, data)
}

// UpdateProject godoc
//
//	@Summary		Update one's project
//	@Description	Update user's project(login required)
//	@Accept			application/json
//	@Param			id				query	string	true	"the id of the project"
//	@Param			Authorization	header	string	true	"token"
//	@Param			newproject		body	object	true	"operating project"
//	@Produce		json
//	@Success		200	{object}	handler.Response
//	@Router			/project [put]
func UpdateProject(r *gin.Context) {
	sid := r.Query("id")
	if sid == "" {
		SendBadRequest(r, model.ErrBadRequest, nil, model.ErrorSender())
		return
	}
	id, _ := strconv.Atoi(sid)
	uid := r.GetInt("userID")
	data := db.ProposalInfo{
		InfoID: int32(id),
	}
	data = model.GetSth(data)
	if data.UID != int32(uid) {
		SendBadRequest(r, model.ErrNotAuthrized, nil, model.ErrorSender())
		return
	}
	r.ShouldBindJSON(&data)
	model.UpdateSth(data)
	SendResponse(r, nil, model.NoResponse)
}

// GetTemplate godoc
//
//	@Summary		Get a templte
//	@Description	Get a template with its id
//	@Param			id				query	string	true	"the id of the template"
//	@Param			Authorization	header	string	true	"token"
//	@Produce		json
//	@Success		200	{object}	db.Template
//	@Failure		404	{object}	handler.Response
//	@Router			/project/template [get]
func GetTemplate(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil || id == 0 {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
	}
	data := db.Template{Temid: int32(id)}
	data = model.GetSth(data)
	SendResponse(r, nil, data)
}

// CreateProject godoc
//
//	@Summary		Create a new project
//	@Description	Create user's project(login required)
//	@Accept			application/json
//	@Param			Authorization	header	string	true	"token"
//	@Param			newproject		body	object	true	"operating project"
//	@Produce		json
//	@Success		200	{object}	handler.Response
//	@Router			/newproject [post]
func CreateProject(r *gin.Context) {
	data := new(db.ProposalInfo)
	r.ShouldBindJSON(&data)
	data.UID = int32(r.GetInt("userID"))
	data.Time = time.Now()
	if data.Budget == "" {
		data.Budget = "{}"
	}
	if data.Nodes == "" {
		data.Nodes = "{}"
	}
	model.CreateSth(*data)
	SendResponse(r, nil, data)
}

func GameSelect(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("game_id"))
	if err != nil || id == 0 {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
	}
	//搜索
	data := new(db.Game)
	data.Gameid = int32(id)
	SendResponse(r, nil, model.GetSth(*data))
}

func FindGames(r *gin.Context) {
	data := new(db.Game)
	r.ShouldBindJSON(&data)
	SendResponse(r, nil, model.GetGames(*data))
}
