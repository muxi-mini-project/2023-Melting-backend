package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

type hook struct {
	CallbackUrl string `json:"callback_url"`
	PushData    struct {
		PushedAt int    `json:"pushed_at"`
		Pusher   string `json:"pusher"`
		Tag      string `json:"tag"`
	} `json:"push_data"`
	Repository struct {
		CommentCount    int    `json:"comment_count"`
		DateCreated     int    `json:"date_created"`
		Description     string `json:"description"`
		Dockerfile      string `json:"dockerfile"`
		FullDescription string `json:"full_description"`
		IsOfficial      bool   `json:"is_official"`
		IsPrivate       bool   `json:"is_private"`
		IsTrusted       bool   `json:"is_trusted"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		Owner           string `json:"owner"`
		RepoName        string `json:"repo_name"`
		RepoUrl         string `json:"repo_url"`
		StarCount       int    `json:"star_count"`
		Status          string `json:"status"`
	} `json:"repository"`
}

type callback struct {
	State       string `json:"state"`
	Description string `json:"description"`
	Context     string `json:"context"`
}

func main() {
	r := gin.New()
	r.POST("", func(r *gin.Context) {
		data := new(hook)
		err := r.ShouldBindJSON(&data)
		if err != nil {
			return
		}
		output, err := exec.Command("./run.sh").Output()
		if err != nil {
			res := callback{
				State:       "failure",
				Description: err.Error(),
				Context:     string(output),
			}
			tmp, _ := json.Marshal(res)
			cli := http.Client{}
			req, _ := http.NewRequest("POST", data.CallbackUrl, bytes.NewBuffer(tmp))
			cli.Do(req)
			return
		}
		res := callback{
			State:       "success",
			Description: "Deployed",
			Context:     string(output),
		}
		tmp, _ := json.Marshal(res)
		cli := http.Client{}
		req, _ := http.NewRequest("POST", data.CallbackUrl, bytes.NewBuffer(tmp))
		cli.Do(req)
		return
	})
	r.Run(":65002")
}
