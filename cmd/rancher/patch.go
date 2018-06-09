package rancher

import (
	"encoding/json"
	"fmt"
	"time"
	"strconv"
)

type Patch struct {
	Spec `json:"spec"`
}

type Spec struct {
	Template `json:"template"`
}

type Template struct {
	Metadata `json:"metadata"`
}

type Metadata struct {
	Labels `json:"labels"`
}

type Labels struct {
	UpdatedAt string `json:"updated-at"`
}

func GetDeploymentPatchString() string {
	patch := Patch{
		Spec{
			Template{
				Metadata{
					Labels{
						UpdatedAt: strconv.FormatInt(time.Now().Unix(), 10),
					},
				},
			},
		},
	}

	jsonByteArray, err := json.Marshal(patch)

	if err != nil {
		panic(fmt.Errorf("Could not marshal as json: %s \n", err))
		return ""
	}

	return string(jsonByteArray[:])

}
