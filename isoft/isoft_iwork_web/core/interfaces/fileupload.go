package interfaces

type IFileUpload interface {
	SaveFile() (tempFileName, fileName, tempFilePath string)
}
