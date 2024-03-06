package tools

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/quic-go-no-crypto/quic-go-no-crypto"
	"github.com/quic-go-no-crypto/quic-go-no-crypto/internal/utils"
	"github.com/quic-go-no-crypto/quic-go-no-crypto/logging"
	"github.com/quic-go-no-crypto/quic-go-no-crypto/qlog"
)

func NewQlogger(logger io.Writer) func(context.Context, logging.Perspective, quic.ConnectionID) *logging.ConnectionTracer {
	return func(_ context.Context, p logging.Perspective, connID quic.ConnectionID) *logging.ConnectionTracer {
		role := "server"
		if p == logging.PerspectiveClient {
			role = "client"
		}
		filename := fmt.Sprintf("log_%s_%s.qlog", connID, role)
		fmt.Fprintf(logger, "Creating %s.\n", filename)
		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("failed to create qlog file: %s", err)
			return nil
		}
		bw := bufio.NewWriter(f)
		return qlog.NewConnectionTracer(utils.NewBufferedWriteCloser(bw, f), p, connID)
	}
}
