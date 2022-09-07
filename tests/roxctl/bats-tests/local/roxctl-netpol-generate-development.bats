#!/usr/bin/env bats

load "../helpers.bash"

out_dir=""

setup_file() {
    command -v yq >/dev/null || skip "Tests in this file require yq"
    echo "Using yq version: '$(yq4.16 --version)'" >&3
    # as of Aug 2022, we run yq version 4.16.2
    # remove binaries from the previous runs
    rm -f "$(roxctl-development-cmd)" "$(roxctl-release-cmd)"
    echo "Testing roxctl version: '$(roxctl-development version)'" >&3
}

setup() {
  out_dir="$(mktemp -d -u)"
  ofile="$(mktemp)"
  export ROX_ROXCTL_NETPOL_GENERATE='true'
}

teardown() {
  rm -rf "$out_dir"
  rm -f "$ofile"
}

@test "roxctl-development generate netpol should respect ROX_ROXCTL_NETPOL_GENERATE feature-flag at runtime" {
  export ROX_ROXCTL_NETPOL_GENERATE=false
  run roxctl-development generate netpol "$out_dir"
  assert_failure
  assert_line --partial 'unknown command "generate"'
}


@test "roxctl-development generate netpol should return error on empty or non-existing directory" {
  run roxctl-development generate netpol "$out_dir"
  assert_failure
  assert_line --regexp "[eE]rror synthesizing policies from folder: no deployment objects discovered in the repository"

  run roxctl-development generate netpol
  assert_failure
  assert_line --partial "accepts 1 arg(s), received 0"
}

@test "roxctl-development generate netpol generates network policies" {
  assert_file_exist "${test_data}/np-guard/scenario-minimal-service/frontend.yaml"
  assert_file_exist "${test_data}/np-guard/scenario-minimal-service/backend.yaml"
  echo "Writing network policies to ${ofile}" >&3
  run roxctl-development generate netpol "${test_data}/np-guard/scenario-minimal-service"
  assert_success

  echo "$output" > "$ofile"
  assert_file_exist "$ofile"
  yaml_valid "$ofile"

  # There must be at least 2 yaml documents in the output
  # yq version 4.16.2 has problems with handling 'document_index', thus we use 'di'
  run yq e 'di' "${ofile}"
  assert_line '0'
  assert_line '1'

  # Ensure that both yaml docs are of kind 'NetworkPolicy'
  run yq e '.kind | ({"match": ., "doc": di})' "${ofile}"
  assert_line --index 0 'match: NetworkPolicy'
  assert_line --index 1 'doc: 0'
  assert_line --index 2 'match: NetworkPolicy'
  assert_line --index 3 'doc: 1'
}

