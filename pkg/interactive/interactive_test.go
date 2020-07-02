package interactive

import (
	"fmt"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestNewInteractive(t *testing.T) {

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	paths := []string{fmt.Sprintf("%s/testutils/kubeconfig", pwd)}

	config := Config{
		SelectCluster:   false,
		KubeConfigPaths: paths,
	}
	interactiveManager, err := NewInteractive(&config)
	log.Info(err)
	log.Info(interactiveManager)
	log.Info(interactiveManager.ctx.Name)

	exemptedContex := "cluster-b"
	if interactiveManager.ctx.Name != exemptedContex {
		t.Fatalf("unexpected context name, got %s expected %s", interactiveManager.ctx.Name, exemptedContex)
	}

}
