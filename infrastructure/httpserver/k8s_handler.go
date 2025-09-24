package httpserver

import (
	"context"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
	"github.com/miraccan00/k8s-cleanup-api/application"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// K8sHandler handles K8s resource endpoints
type K8sHandler struct {
	Service *application.K8sService
}

func NewK8sHandler() *K8sHandler {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeconfig = ""
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	service := application.NewK8sService(clientset)
	return &K8sHandler{Service: service}
}

func (h *K8sHandler) GetUnusedConfigMaps(c fiber.Ctx) error {
	unused, err := h.Service.GetUnusedConfigMaps(context.Background())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(unused)
}

func (h *K8sHandler) GetUnusedPVCs(c fiber.Ctx) error {
	unused, err := h.Service.GetUnusedPVCs(context.Background())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(unused)
}

func (h *K8sHandler) GetUnusedSecrets(c fiber.Ctx) error {
	unused, err := h.Service.GetUnusedSecrets(context.Background())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(unused)
}
