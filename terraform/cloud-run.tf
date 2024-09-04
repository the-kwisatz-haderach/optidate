resource "google_cloud_run_v2_service" "default" {
  name     = "optidate"
  location = "europe-west1"
  ingress  = "INGRESS_TRAFFIC_ALL"

  traffic {
    type    = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
    percent = 100
  }

  template {
    service_account = google_service_account.cloud_run_invoker.email
    scaling {
      max_instance_count = 2
      min_instance_count = 0
    }
    containers {
      image = "europe-west1-docker.pkg.dev/optidate/optidate/${google_artifact_registry_repository.default.name}"
      resources {
        limits = {
          cpu    = "1000m"
          memory = "256Mi"
        }
        cpu_idle          = true
        startup_cpu_boost = false
      }
      ports {
        container_port = 8000
      }
      startup_probe {
        initial_delay_seconds = 0
        timeout_seconds       = 240
        period_seconds        = 240
        failure_threshold     = 1
        tcp_socket {
          port = 8000
        }
      }
    }
  }

  depends_on = [google_artifact_registry_repository.default]
  lifecycle {
    ignore_changes = [
      client,
      client_version,
      template[0].containers[0].image,
    ]
  }
}

resource "google_cloud_run_service_iam_binding" "default" {
  location = google_cloud_run_v2_service.default.location
  service  = google_cloud_run_v2_service.default.name
  role     = "roles/run.invoker"
  members = [
    "allUsers"
  ]
}
