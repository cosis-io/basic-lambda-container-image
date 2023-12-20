Metadata {
  description =  "Control or Application Plane "
  service     = "%%env%%-<plane type>-plane-<handler description>"
  version     = "%%version%%"
}

Settings {
  api_version     = "v1"
  charset         = "UTF-8"
  http_timeout    = 80
  is_app_get_only = false
  workflow_stage  = "%%env%%.sch00l.<handler stage>"
}
