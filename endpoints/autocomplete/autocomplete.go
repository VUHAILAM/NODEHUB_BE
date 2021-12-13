package autocomplete

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"gitlab.com/hieuxeko19991/job4e_be/services/autocomplete"
	"go.uber.org/zap"
)

type AutocompleteSerialize struct {
	JobTrie *autocomplete.Trie
	CanTrie *autocomplete.Trie
	RecTrie *autocomplete.Trie

	Logger *zap.Logger
}

func (s *AutocompleteSerialize) AutocompleteJob(ginCtx *gin.Context) {
	req := models.RequestGetAutocomplete{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request get autocomplete error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := s.JobTrie.Search(req.Text)
	if req.Limit != 0 && req.Limit <= int64(len(res)) {
		ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, res[:req.Limit])
	} else {
		ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, res)
	}
}

func (s *AutocompleteSerialize) AutocompleteRec(ginCtx *gin.Context) {
	req := models.RequestGetAutocomplete{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request get autocomplete error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := s.RecTrie.Search(req.Text)
	if req.Limit != 0 && req.Limit <= int64(len(res)) {
		ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, res[:req.Limit])
	} else {
		ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, res)
	}
}

func (s *AutocompleteSerialize) AutocompleteCan(ginCtx *gin.Context) {
	req := models.RequestGetAutocomplete{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request get autocomplete error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := s.CanTrie.Search(req.Text)
	if req.Limit != 0 && req.Limit <= int64(len(res)) {
		ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, res[:req.Limit])
	} else {
		ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, res)
	}
}
