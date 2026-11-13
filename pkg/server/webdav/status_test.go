package webdav

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/sirrobot01/decypharr/internal/nntp"
)

func TestStreamErrorStatus(t *testing.T) {
	missing := fmt.Errorf("%w: %w", nntp.ErrAllProvidersFailed, &nntp.Error{Type: nntp.ErrorTypeArticleNotFound, Code: 430})
	yenc := fmt.Errorf("stream: %w", &nntp.Error{Type: nntp.ErrorTypeYencDecode})
	busy := &nntp.Error{Type: nntp.ErrorTypeServerBusy}
	for err, want := range map[error]int{missing: http.StatusNotFound, yenc: http.StatusNotFound, busy: http.StatusInternalServerError, errors.New("x"): http.StatusInternalServerError} {
		if got := streamErrorStatus(err); got != want {
			t.Errorf("%v: %d, want %d", err, got, want)
		}
	}
}
