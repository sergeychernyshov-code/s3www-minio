package main

import (
    "context"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    minioClient *minio.Client
    bucket      string

    downloadCounter = prometheus.NewCounter(
        prometheus.CounterOpts{
            Name: "s3_file_download_total",
            Help: "Total number of files downloaded from S3",
        },
    )

    requestDuration = prometheus.NewHistogram(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Duration of HTTP requests.",
            Buckets: prometheus.DefBuckets,
        },
    )

    activeRequests = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "http_active_requests",
            Help: "Current number of active requests.",
        },
    )

    errorCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "s3_file_errors_total",
            Help: "Total number of errors during file download",
        },
        []string{"type"},
    )
)

func init() {
    prometheus.MustRegister(downloadCounter)
    prometheus.MustRegister(requestDuration)
    prometheus.MustRegister(activeRequests)
    prometheus.MustRegister(errorCounter)
}

func handler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    activeRequests.Inc()
    defer func() {
        requestDuration.Observe(time.Since(start).Seconds())
        activeRequests.Dec()
    }()

    key := strings.TrimPrefix(r.URL.Path, "/")
    if key == "" {
        http.Error(w, "Missing file key in path", http.StatusBadRequest)
        errorCounter.WithLabelValues("missing_key").Inc()
        return
    }

    object, err := minioClient.GetObject(context.Background(), bucket, key, minio.GetObjectOptions{})
    if err != nil {
        http.Error(w, "Failed to get object", http.StatusInternalServerError)
        log.Printf("Error fetching object: %v", err)
        errorCounter.WithLabelValues("minio_get_error").Inc()
        return
    }

    stat, err := object.Stat()
    if err != nil {
        http.Error(w, "Object not found", http.StatusNotFound)
        log.Printf("Error stating object: %v", err)
        errorCounter.WithLabelValues("minio_stat_error").Inc()
        return
    }

    downloadCounter.Inc()
    w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename="%s"", key))
    w.Header().Set("Content-Type", stat.ContentType)
    w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size))
    io.Copy(w, object)
}

func main() {
    endpoint := os.Getenv("MINIO_ENDPOINT")
    accessKey := os.Getenv("MINIO_ACCESS_KEY")
    secretKey := os.Getenv("MINIO_SECRET_KEY")
    useSSL := os.Getenv("MINIO_USE_SSL") == "true"
    bucket = os.Getenv("S3_BUCKET")

    var err error
    minioClient, err = minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        log.Fatalf("Failed to initialize MinIO client: %v", err)
    }

    http.HandleFunc("/", handler)
    http.Handle("/metrics", promhttp.Handler())

    port := "8080"
    log.Printf("Starting server on :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}