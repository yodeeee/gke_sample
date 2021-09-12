package main

deny[msg] {
  resources := ["DaemonSet", "Deployment", "ReplicaSet"]
  apis := ["apps/v1beta1", "apps/v1beta2", "extensions/v1beta1"]
  input.apiVersion == "extensions/v1beta1"
  input.kind == resources[_]
  input.apiVersion == apis[_]
  msg := sprintf("%s/%s: API %s はデフォルトで提供されなくなりました。代わりに apps/v1 を使用してください。", [input.kind, input.metadata.name, input.apiVersion])
}

deny[msg] {
  docker_images := input.jobs[_].docker[_].image
  not startswith(docker_images, "circleci/")
  msg = "Only use official CircleCI images"
}

deny[msg] {
  docker_images := input.jobs[_].docker[_].image
  tag_is_latest(split(docker_images, ":"))
  msg = "Do not use `latest` container image tags"
}

# helpers
tag_is_latest(images) {
  count(images) < 2
}

tag_is_latest([_, tag]) {
  tag == "latest"
}
