package utils

import (
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

func TratarImagemProduto(c *gin.Context, imagem *multipart.FileHeader, produto *models.Produto) error {

	savePath := "./client/assets/images/" + imagem.Filename
	if savePath == produto.Imagem {
		os.Remove(savePath)
	}
	err := c.SaveUploadedFile(imagem, savePath)
	return err
}