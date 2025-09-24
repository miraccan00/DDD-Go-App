// Deprecated: moved to application/k8s_service.go for DDD structure
// This file is no longer used.

package k8s

import (
	"context"
	"sort"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type K8sService struct {
	Client *kubernetes.Clientset
}

func NewK8sService(client *kubernetes.Clientset) *K8sService {
	return &K8sService{Client: client}
}

func (s *K8sService) GetUnusedConfigMaps(ctx context.Context) ([]string, error) {
	// Tüm ConfigMap'leri ve kullanılanları bul
	allCM := make(map[string]struct{})
	usedCM := make(map[string]struct{})

	nsl, err := s.Client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, ns := range nsl.Items {
		cms, err := s.Client.CoreV1().ConfigMaps(ns.Name).List(ctx, metav1.ListOptions{})
		if err != nil {
			continue
		}
		for _, cm := range cms.Items {
			allCM[ns.Name+"/"+cm.Name] = struct{}{}
		}
		// Kullanılanları bulmak için workload'ları gezmek gerekir (örnek: sadece Deployment)
		deps, err := s.Client.AppsV1().Deployments(ns.Name).List(ctx, metav1.ListOptions{})
		if err == nil {
			for _, d := range deps.Items {
				for _, v := range d.Spec.Template.Spec.Volumes {
					if v.ConfigMap != nil && v.ConfigMap.Name != "" {
						usedCM[ns.Name+"/"+v.ConfigMap.Name] = struct{}{}
					}
				}
			}
		}
	}
	var unused []string
	for k := range allCM {
		if _, ok := usedCM[k]; !ok {
			unused = append(unused, k)
		}
	}
	sort.Strings(unused)
	return unused, nil
}

func (s *K8sService) GetUnusedPVCs(ctx context.Context) ([]string, error) {
	allPVC := make(map[string]struct{})
	usedPVC := make(map[string]struct{})

	nsl, err := s.Client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, ns := range nsl.Items {
		pvcs, err := s.Client.CoreV1().PersistentVolumeClaims(ns.Name).List(ctx, metav1.ListOptions{})
		if err != nil {
			continue
		}
		for _, pvc := range pvcs.Items {
			allPVC[ns.Name+"/"+pvc.Name] = struct{}{}
		}
		deps, err := s.Client.AppsV1().Deployments(ns.Name).List(ctx, metav1.ListOptions{})
		if err == nil {
			for _, d := range deps.Items {
				for _, v := range d.Spec.Template.Spec.Volumes {
					if v.PersistentVolumeClaim != nil && v.PersistentVolumeClaim.ClaimName != "" {
						usedPVC[ns.Name+"/"+v.PersistentVolumeClaim.ClaimName] = struct{}{}
					}
				}
			}
		}
	}
	var unused []string
	for k := range allPVC {
		if _, ok := usedPVC[k]; !ok {
			unused = append(unused, k)
		}
	}
	sort.Strings(unused)
	return unused, nil
}

func (s *K8sService) GetUnusedSecrets(ctx context.Context) ([]string, error) {
	allSecrets := make(map[string]struct{})
	usedSecrets := make(map[string]struct{})

	nsl, err := s.Client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, ns := range nsl.Items {
		secrets, err := s.Client.CoreV1().Secrets(ns.Name).List(ctx, metav1.ListOptions{})
		if err != nil {
			continue
		}
		for _, secret := range secrets.Items {
			allSecrets[ns.Name+"/"+secret.Name] = struct{}{}
		}
		deps, err := s.Client.AppsV1().Deployments(ns.Name).List(ctx, metav1.ListOptions{})
		if err == nil {
			for _, d := range deps.Items {
				for _, v := range d.Spec.Template.Spec.Volumes {
					if v.Secret != nil && v.Secret.SecretName != "" {
						usedSecrets[ns.Name+"/"+v.Secret.SecretName] = struct{}{}
					}
				}
			}
		}
	}
	var unused []string
	for k := range allSecrets {
		if _, ok := usedSecrets[k]; !ok {
			unused = append(unused, k)
		}
	}
	sort.Strings(unused)
	return unused, nil
}
