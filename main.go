package steamwheedle_cartel_download_auctions

import (
	"context"
	"os"

	"github.com/sotah-inc/steamwheedle-cartel/pkg/logging"
	"github.com/sotah-inc/steamwheedle-cartel/pkg/state/fn"
)

var (
	projectId = os.Getenv("GCP_PROJECT")

	sta fn.DownloadAuctionsState
	err error
)

func init() {
	sta, err = fn.NewDownloadAuctionsState(fn.DownloadAuctionsStateConfig{ProjectId: projectId})
	if err != nil {
		logging.WithField("error", err.Error()).Fatal("Failed to establish state")

		return
	}
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func DownloadAuctions(_ context.Context, m PubSubMessage) error {
	return sta.Run(string(m.Data))
}
