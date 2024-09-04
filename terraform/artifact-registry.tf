resource "google_artifact_registry_repository" "default" {
  repository_id = "optidate"
  description   = "repository for optidate docker images"
  format        = "DOCKER"
}
