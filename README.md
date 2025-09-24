# K8s Cleanup API (Hobby Project)

Bu proje, Kubernetes cluster'ındaki kullanılmayan ConfigMap, PVC ve Secret kaynaklarını tespit etmek için REST API sunar. DDD (Domain Driven Design) mimarisiyle geliştirilmiştir ve Go ile yazılmıştır.

## Amaç
Kubernetes ortamlarında zamanla biriken ve kullanılmayan kaynakları (ConfigMap, PVC, Secret) kolayca tespit ederek temizlik ve optimizasyon sağlar. Özellikle büyük ve dinamik ortamlarda gereksiz kaynakların tespiti için idealdir.

## Özellikler
- Kullanılmayan ConfigMap, PVC ve Secret'ları GET endpointleri ile listeler
- Fiber v3 ile hızlı ve modern REST API
- Hem local (kubeconfig) hem de cluster içi bağlantı desteği
- DDD'ye uygun modüler ve test edilebilir yapı

## Kurulum
1. Go 1.20+ yüklü olmalı
2. Bağımlılıkları yükleyin:
   ```bash
   go mod tidy
   ```
3. Uygulamayı derleyin:
   ```bash
   go build -o bin/app ./cmd/main.go
   ```
4. Uygulamayı başlatın:
   ```bash
   ./bin/app
   ```

## API Kullanımı
Aşağıdaki endpointler ile kullanılmayan kaynakları listeleyebilirsiniz:

- Kullanılmayan ConfigMap'ler:
  ```bash
  curl http://localhost:8081/unused-configmaps
  ```
- Kullanılmayan PVC'ler:
  ```bash
  curl http://localhost:8081/unused-pvc
  ```
- Kullanılmayan Secret'lar:
  ```bash
  curl http://localhost:8081/unused-secrets
  ```

## Geliştirici Notları
- Kod DDD'ye uygun olarak application, infrastructure, domain katmanlarına ayrılmıştır.
- K8s ile bağlantı localde kubeconfig ile, cluster içinde ise in-cluster config ile sağlanır.
- Test ve geliştirme için localde kubeconfig dosyanızın olması yeterlidir.

## Katkı
Pull request ve issue açarak katkıda bulunabilirsiniz.

## Lisans
MIT
