package controller

import (
	"example/web-service-gin/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumCotroller struct {
	CreateAlbumUsecase usecase.CreateAlbumUsecase
}

func NewAlbumController(
	createAlbumUsecase usecase.CreateAlbumUsecase,
) *AlbumCotroller {
	return &AlbumCotroller{
		CreateAlbumUsecase: createAlbumUsecase,
	}
}

func (con *AlbumCotroller) CreateAlbum(ctx *gin.Context) {
	var newAlbum request.AlbumCreateRequest
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed bind json"})
		return
	}
	err := con.CreateAlbumUsecase.Execute(
		ctx,
		newAlbum.ID,
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Price,
	)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"gin": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, "OK")
}
