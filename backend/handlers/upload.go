package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

// UploadFile 上传文件到服务器
func (h *UploadHandler) UploadFile(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "文件上传失败",
		})
		return
	}

	// 验证文件大小（最大10MB）
	maxSize := int64(10 * 1024 * 1024)
	if file.Size > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "文件大小不能超过10MB",
		})
		return
	}

	// 验证文件类型
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
	}

	// 打开文件读取MIME类型
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文件读取失败",
		})
		return
	}
	defer src.Close()

	// 读取文件头512字节用于检测MIME类型
	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文件类型检测失败",
		})
		return
	}

	// 检测MIME类型
	mimeType := http.DetectContentType(buffer)
	if !allowedTypes[mimeType] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "不支持的文件类型，仅支持 JPEG, PNG, GIF, WebP",
		})
		return
	}

	// 重置文件指针
	src.Seek(0, 0)

	// 生成文件名
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		// 根据MIME类型确定扩展名
		switch mimeType {
		case "image/jpeg":
			ext = ".jpg"
		case "image/png":
			ext = ".png"
		case "image/gif":
			ext = ".gif"
		case "image/webp":
			ext = ".webp"
		}
	}

	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d_%s%s", timestamp, generateRandomString(8), ext)

	// 确保uploads目录存在
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "创建上传目录失败",
		})
		return
	}

	// 保存文件
	filepath := filepath.Join(uploadDir, filename)
	dst, err := os.Create(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文件保存失败",
		})
		return
	}
	defer dst.Close()

	// 复制文件内容
	if _, err = io.Copy(dst, src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文件写入失败",
		})
		return
	}

	// 返回文件URL
	// 注意：这里需要根据实际部署的域名修改
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	fileURL := fmt.Sprintf("%s/uploads/%s", baseURL, filename)

	c.JSON(http.StatusOK, gin.H{
		"url":      fileURL,
		"filename": filename,
		"size":     file.Size,
		"mime":     mimeType,
	})
}

// GetCOSCredentials 获取腾讯云COS临时密钥
func (h *UploadHandler) GetCOSCredentials(c *gin.Context) {
	// TODO: 实现腾讯云COS临时密钥获取
	// 参考：https://cloud.tencent.com/document/product/436/14048

	c.JSON(http.StatusOK, gin.H{
		"sessionToken":  "临时Token",
		"authorization": "临时授权",
		"expiredTime":   time.Now().Add(time.Hour).Unix(),
	})
}

// GetOSSCredentials 获取阿里云OSS临时凭证
func (h *UploadHandler) GetOSSCredentials(c *gin.Context) {
	// TODO: 实现阿里云OSS STS临时凭证获取
	// 参考：https://help.aliyun.com/document_detail/100624.html

	c.JSON(http.StatusOK, gin.H{
		"accessKeyId":     "临时AccessKeyId",
		"accessKeySecret": "临时AccessKeySecret",
		"securityToken":   "临时SecurityToken",
		"expiration":      time.Now().Add(time.Hour).Format(time.RFC3339),
		"policy":          "上传策略",
		"signature":       "签名",
	})
}

// generateRandomString 生成随机字符串
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	timestamp := time.Now().UnixNano()

	for i := range result {
		result[i] = charset[(timestamp+int64(i))%int64(len(charset))]
	}

	return string(result)
}

// DeleteFile 删除文件（可选功能）
func (h *UploadHandler) DeleteFile(c *gin.Context) {
	filename := c.Param("filename")

	// 验证文件名，防止路径遍历攻击
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "非法的文件名",
		})
		return
	}

	// 删除文件
	filepath := filepath.Join("uploads", filename)
	if err := os.Remove(filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文件删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "文件删除成功",
	})
}
