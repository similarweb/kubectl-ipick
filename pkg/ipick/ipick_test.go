package ipick

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
	interactiveManager, _ := NewIpick(&config)

	exemptedContex := "cluster-b"
	if interactiveManager.ctx.Name != exemptedContex {
		t.Fatalf("unexpected context name, got %s expected %s", interactiveManager.ctx.Name, exemptedContex)
	}

}

func TestRandomInteger(t *testing.T) {

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	paths := []string{fmt.Sprintf("%s/testutils/kubeconfig", pwd)}

	config := Config{
		SelectCluster:   false,
		KubeConfigPaths: paths,
	}
	interactiveManager, _ := NewIpick(&config)

	randomNumber := interactiveManager.randomInteger(1, 3)
	if randomNumber != 1 && randomNumber != 2 && randomNumber != 3 {
		t.Fatalf("unexpected random number, got %d expected %s", randomNumber, "1-3")
	}

}

func TestAddRightPadding(t *testing.T) {

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	paths := []string{fmt.Sprintf("%s/testutils/kubeconfig", pwd)}

	config := Config{
		SelectCluster:   false,
		KubeConfigPaths: paths,
	}
	interactiveManager, _ := NewIpick(&config)

	str := interactiveManager.addRightPadding("test", 20, "-")
	if str != "test----------------" {
		t.Fatalf("unexpected padding string, got %s expected %s", str, "test----------------")
	}

}
