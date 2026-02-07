package ctrl

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/EagleLizard/ezd-daemon/internal/lib/config"
	"github.com/EagleLizard/ezd-daemon/internal/lib/logging"
	gh "github.com/EagleLizard/ezd-daemon/internal/lib/model/github"
)

/*
Listener endpoint for github webhooks.

	See:
		- https://docs.github.com/en/webhooks/webhook-events-and-payloads#ping
		- https://github.com/GitbookIO/go-github-webhook/blob/master/handler.go
*/
func PostGhHook(cfg *config.EzdDConfigType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("hello\n")
		event := r.Header.Get("x-github-event")
		fmt.Printf("event: %s\n", event)
		signature := r.Header.Get("x-hub-signature-256")
		fmt.Printf("%s\n", signature)
		if signature == "" || !strings.HasPrefix(signature, "sha256=") {
			logging.Logger.Sugar().Error("Invalid or missing signature")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logging.Logger.Sugar().Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		validPayload, err := validatePayload(body, signature)
		if err != nil {
			logging.Logger.Sugar().Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !validPayload {
			logging.Logger.Sugar().Error(errors.New("Invalid signature"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		/* pass to relevant handler based on action */
		switch event {
		case "push":
			pushHandler(body)
		default:
			logging.Logger.Sugar().Infof("Unhandled event: %s", event)
		}
		w.WriteHeader(http.StatusOK)
	}
}

func pushHandler(body []byte) error {
	fmt.Print("Push Event\n")
	var ghp = gh.GhPushPayload{}
	err := json.Unmarshal(body, &ghp)
	if err != nil {
		return err
	}
	fmt.Printf("repo: %s\nref: %s\n", ghp.Repository.Full_name, ghp.Ref)
	/*
		- lookup config for repo that was pushed
			- if not found, return error
		- config will provide info on how to handle the push
			- to start, point at relevant local deploy script(s) for the repo
			- deploy validation:
				- on failure, rollback. Nice to have.
	*/
	switch ghp.Repository.Name {
	case "ezd-api-rc2":
		ezdApiRc2(ghp)
	default:
		logging.Logger.Sugar().Infow(fmt.Sprintf("no config for repo '%s'", ghp.Repository.Name))
	}
	// cmd := exec.Command("docker", "ps")
	// out, err := cmd.Output()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%s\n", out)
	return nil
}

func ezdApiRc2(ghp gh.GhPushPayload) {
	fmt.Printf("ezd api rc2 cfg !!! \n")
}

func validatePayload(body []byte, signature string) (bool, error) {
	sigHash := strings.Split(signature, "=")[1]
	hm := hmac.New(sha256.New, []byte(config.EzdDConfig.EzdGhWebhookSecret))
	_, err := hm.Write(body)
	if err != nil {
		return false, err
	}
	sum := hm.Sum(nil)
	h := fmt.Sprintf("%x", sum)
	return hmac.Equal([]byte(h), []byte(sigHash)), nil
}
