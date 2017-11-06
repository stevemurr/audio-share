package handler

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"murrman/audio-share/model"
	"murrman/audio-share/storage"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Handler --
type Handler struct {
	DB storage.Service
}

// GetMD5Hash --
func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Upload --
func (h *Handler) Upload(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dst := &bytes.Buffer{}
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	audio := &model.Audio{
		Name:       file.Filename,
		Data:       dst.Bytes(),
		UploadedAt: time.Now(),
		Hash:       getMD5Hash(dst.String()),
		Regions:    model.Regions{},
	}
	h.DB.PutAudio(audio.Hash, audio)
	return c.String(http.StatusOK, audio.Hash)
	// return c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("http://audio-share.2017ditrfest.info/%s", audio.Hash))
}

// List --
func (h *Handler) List(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, h.DB.GetAll(), "  ")
}

// GetFile --
func (h *Handler) GetFile(c echo.Context) error {
	f := c.Param("file")
	return c.JSON(http.StatusOK, h.DB.GetAudio(f))
}

// GetPayload --
func (h *Handler) GetPayload(c echo.Context) error {
	f := c.Param("file")
	return c.Blob(http.StatusOK, "audio/wav", h.DB.GetAudioData(f))
}

// DeleteRegion --
func (h *Handler) DeleteRegion(c echo.Context) error {
	id := c.Param("id")
	var req model.Region
	if err := c.Bind(&req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	h.DB.DeleteAudioRegion(id, req)
	return c.NoContent(http.StatusOK)
}

// UpdateRegions --
func (h *Handler) UpdateRegions(c echo.Context) error {
	id := c.Param("id")
	var req []*model.Region
	if err := c.Bind(&req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	h.DB.PutAudioRegion(id, req)
	return c.NoContent(http.StatusOK)
}
