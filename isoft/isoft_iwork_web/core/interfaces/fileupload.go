package interfaces

type IFileUpload interface {
	SaveFile(suffixs []string) (tempFileName, fileName, tempFilePath string)
}
