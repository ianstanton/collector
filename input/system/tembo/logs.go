package tembo

import (
	"context"
	"fmt"
	"github.com/pganalyze/collector/state"
	"github.com/pganalyze/collector/util"
	"github.com/pkg/errors"
	"net/url"
)

// DownloadLogFiles - Gets log files for a Tembo instance
func DownloadLogFiles(ctx context.Context, server *state.Server, logger *util.Logger) (state.PersistedLogState, []state.LogFile, []state.PostgresQuerySample, error) {
	//var err error
	//var psl state.PersistedLogState = server.LogPrevState
	//var logFiles []state.LogFile
	//var samples []state.PostgresQuerySample

	// If tembo_api_token is not set, return an error
	if server.Config.TemboAPIToken == "" {
		return server.LogPrevState, nil, nil, errors.New("tembo_api_token not set")
	}
	// If tembo_instance_id is not set, return an error
	if server.Config.TemboInstanceID == "" {
		return server.LogPrevState, nil, nil, errors.New("tembo_instance_id not set")
	}
	// If tembo_org_id is not set, return an error
	if server.Config.TemboOrgID == "" {
		return server.LogPrevState, nil, nil, errors.New("tembo_org_id not set")
	}

	// Construct query for Tembo Logs API
	query := "{tembo_instance_id=\"" + server.Config.TemboInstanceID + "\"}"

	// URI encode query
	encodedQuery := url.QueryEscape(query)

	// Construct URL for Tembo Logs API wss://api.data-1.use1.tembo.io/loki/api/v1/tail?query=$URL_ENCODED_QUERY
	websocketUrl := "wss://api.data-1.use1.tembo.io/loki/api/v1/tail?query=" + encodedQuery
	fmt.Println(websocketUrl)

	return server.LogPrevState, nil, nil, nil
}
