#!/usr/bin/env bats

load "../helpers.bash"

out_dir=""

setup_file() {
  echo "Testing roxctl version: '$(roxctl-development version)'" >&3
  command -v yq || skip "Tests in this file require yq"
  assert_file_exist "$test_data/helm-output-secured-cluster-services/ca-config.yaml"
  assert_file_exist "$test_data/helm-output-secured-cluster-services/cluster-init-bundle.yaml"
}

setup() {
  CLUSTER_NAME="CLUSTER-1"
  out_dir="$(mktemp -d -u)"
}

teardown() {
  rm -rf "$out_dir"
}

@test "roxctl-release helm output secured-cluster-services --image-defaults=development_build should fail" {
  run roxctl-release helm output secured-cluster-services --image-defaults=development_build --remove --output-dir "$out_dir"
  assert_failure
  assert_output --partial "unexpected value 'development_build', allowed values are"
}

@test "roxctl-release helm output secured-cluster-services (default flavor) should use registry.redhat.io registry (default collector)" {
  run roxctl-release helm output secured-cluster-services --remove --output-dir "$out_dir"
  assert_success
  assert_output --partial "Written Helm chart secured-cluster-services to directory"

  helm_template_secured_cluster "$out_dir" "$out_dir/rendered" "$CLUSTER_NAME"
  assert_components_registry "$out_dir/rendered" "registry.redhat.io" "$any_version" 'collector-slim' 'admission-controller' 'sensor'
}

@test "roxctl-release helm output secured-cluster-services --image-defaults=stackrox.io should use stackrox.io registry (default collector)" {
  run roxctl-release helm output secured-cluster-services --image-defaults=stackrox.io --remove --output-dir "$out_dir"
  assert_success
  assert_output --partial "Written Helm chart secured-cluster-services to directory"

  helm_template_secured_cluster "$out_dir" "$out_dir/rendered" "$CLUSTER_NAME"
  assert_components_registry "$out_dir/rendered" "stackrox.io" "$any_version" 'collector-slim' 'admission-controller' 'sensor'
}

@test "roxctl-release helm output secured-cluster-services --image-defaults=stackrox.io should use stackrox.io registry (fat collector)" {
  run roxctl-release helm output secured-cluster-services --image-defaults=stackrox.io --remove --output-dir "$out_dir"
  assert_success
  assert_output --partial "Written Helm chart secured-cluster-services to directory"

  helm_template_secured_cluster "$out_dir" "$out_dir/rendered" "$CLUSTER_NAME" "--set" "collector.slimMode=false"
  assert_components_registry "$out_dir/rendered" "stackrox.io" "$any_version" 'collector' 'admission-controller' 'sensor'
}

@test "roxctl-release helm output secured-cluster-services --image-defaults=stackrox.io should use stackrox.io registry (slim collector)" {
  run roxctl-release helm output secured-cluster-services --image-defaults=stackrox.io --remove --output-dir "$out_dir"
  assert_success
  assert_output --partial "Written Helm chart secured-cluster-services to directory"

  helm_template_secured_cluster "$out_dir" "$out_dir/rendered" "$CLUSTER_NAME" "--set" "collector.slimMode=true"
  assert_components_registry "$out_dir/rendered" "stackrox.io" "$any_version" 'collector-slim' 'admission-controller' 'sensor'
}

# RHACS

@test "roxctl-release helm output secured-cluster-services --image-defaults=rhacs should use registry.redhat.io registry (default collector)" {
  run roxctl-release helm output secured-cluster-services --image-defaults=rhacs --remove --output-dir "$out_dir"
  assert_success
  assert_output --partial "Written Helm chart secured-cluster-services to directory"

  helm_template_secured_cluster "$out_dir" "$out_dir/rendered" "$CLUSTER_NAME"
  assert_components_registry "$out_dir/rendered" "registry.redhat.io" "$any_version" 'collector-slim' 'admission-controller' 'sensor'
}

@test "roxctl-release helm output secured-cluster-services --image-defaults=rhacs should use registry.redhat.io registry (slim collector)" {
  run roxctl-release helm output secured-cluster-services --image-defaults=rhacs --remove --output-dir "$out_dir"
  assert_success
  assert_output --partial "Written Helm chart secured-cluster-services to directory"

  helm_template_secured_cluster "$out_dir" "$out_dir/rendered" "$CLUSTER_NAME" "--set" "collector.slimMode=true"
  assert_components_registry "$out_dir/rendered" "registry.redhat.io" "$any_version" 'collector-slim' 'admission-controller' 'sensor'
}

@test "roxctl-release helm output secured-cluster-services --image-defaults=rhacs should use registry.redhat.io registry (fat collector)" {
  run roxctl-release helm output secured-cluster-services --image-defaults=rhacs --remove --output-dir "$out_dir"
  assert_success
  assert_output --partial "Written Helm chart secured-cluster-services to directory"

  helm_template_secured_cluster "$out_dir" "$out_dir/rendered" "$CLUSTER_NAME" "--set" "collector.slimMode=false"
  assert_components_registry "$out_dir/rendered" "registry.redhat.io" "$any_version" 'collector' 'admission-controller' 'sensor'
}
