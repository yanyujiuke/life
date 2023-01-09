package response

import "life/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
