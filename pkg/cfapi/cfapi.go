package cfapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/shreyashree-17/project/pkg/models"
	"go.uber.org/zap"
)

type CodeforcesAPI interface {
	RecentActions(maxCount int) ([]models.RecentAction, error)
}

type CodeforcesClient struct {
	client http.Client
}

func (cfClient *CodeforcesClient) RecentActions(maxCount int) ([]models.RecentAction, error) {

	resp, err := cfClient.client.Get("https://codeforces.com/api/recentActions?maxCount=100")
	if err != nil {
		zap.S().Errorf("Error occured while calling cf api: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error occurred while reading the resp body")
		return nil, err
	}

	//zap.S().Info(string(data))

	wrapper := struct {
		Status string
		Result []models.RecentAction
	}{}

	if err = json.Unmarshal(data, &wrapper); err != nil {
		zap.S().Errorf("Error while unmarshalling data from cfapi : %v", err)
	}

	return wrapper.Result, err
}

func NewCodeforcesClient() CodeforcesAPI {
	obj := new(CodeforcesClient)
	return obj
}
