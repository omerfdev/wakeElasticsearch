# Elasticsearch ve Kibana Başlatma Servisi

Bu basit Go programı, bir HTTP endpoint'i aracılığıyla Elasticsearch ve Kibana'nın başlatılmasını sağlar.

## Nasıl Kullanılır

1. Uygulamayı başlatmak için öncelikle Go diline sahip bir geliştirme ortamına sahip olmanız gerekir.

2. Proje dizinine gidin ve `go run main.go` komutunu çalıştırarak HTTP sunucusunu başlatın.

3. Tarayıcınızda veya API test aracınızda `http://localhost:8080/start` adresine bir GET isteği göndererek Elasticsearch ve Kibana'yı başlatabilirsiniz.

## Önemli Notlar

- Elasticsearch ve Kibana Docker imajlarına dayanmaktadır. Bu nedenle, bu servisin çalışması için Docker'ın kurulu olması gerekir.
- Elasticsearch ve Kibana'nın başlaması için biraz zaman gerekebilir. Bu yüzden, isteğinize cevap vermeden önce birkaç saniye bekleyin.
- Başlangıçta kullanılan portlar Elasticsearch ve Kibana'nın varsayılan portlarına göre ayarlanmıştır. Bu nedenle, yerel olarak çalışan Elasticsearch ve Kibana'ya bağlanmak için herhangi bir ek yapılandırma gerekmez.
