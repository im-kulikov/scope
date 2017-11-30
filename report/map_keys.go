package report

/* Lookup table to allow msgpack/json decoder to avoid heap allocation
   for common ps.Map keys. The map is static so we don't have to lock
   access from multiple threads and don't have to worry about it
   getting clogged with values that are only used once.
*/
var commonKeys = map[string]string{
	"addr":                             "addr",
	"cmdline":                          "cmdline",
	"conntracked":                      "conntracked",
	"container":                        "container",
	"container_image":                  "container_image",
	"control_probe_id":                 "control_probe_id",
	"deployment":                       "deployment",
	"docker_attach_container":          "docker_attach_container",
	"docker_container_command":         "docker_container_command",
	"docker_container_created":         "docker_container_created",
	"docker_container_hostname":        "docker_container_hostname",
	"docker_container_id":              "docker_container_id",
	"docker_container_ips":             "docker_container_ips",
	"docker_container_ips_with_scopes": "docker_container_ips_with_scopes",
	"docker_container_name":            "docker_container_name",
	"docker_container_network_mode":    "docker_container_network_mode",
	"docker_container_restart_count":   "docker_container_restart_count",
	"docker_container_state":           "docker_container_state",
	"docker_container_state_human":     "docker_container_state_human",
	"docker_container_uptime":          "docker_container_uptime",
	"docker_exec_container":            "docker_exec_container",
	"docker_image_id":                  "docker_image_id",
	"docker_image_name":                "docker_image_name",
	"docker_image_size":                "docker_image_size",
	"docker_image_virtual_size":        "docker_image_virtual_size",
	"docker_is_in_host_network":        "docker_is_in_host_network",
	"docker_pause_container":           "docker_pause_container",
	"docker_remove_container":          "docker_remove_container",
	"docker_restart_container":         "docker_restart_container",
	"docker_start_container":           "docker_start_container",
	"docker_stop_container":            "docker_stop_container",
	"docker_unpause_container":         "docker_unpause_container",
	"does_not_make_connections":        "does_not_make_connections",
	"host":                                "host",
	"host_node_id":                        "host_node_id",
	"kubernetes_available_replicas":       "kubernetes_available_replicas",
	"kubernetes_created":                  "kubernetes_created",
	"kubernetes_desired_replicas":         "kubernetes_desired_replicas",
	"kubernetes_fully_labeled_replicas":   "kubernetes_fully_labeled_replicas",
	"kubernetes_ip":                       "kubernetes_ip",
	"kubernetes_labels_name":              "kubernetes_labels_name",
	"kubernetes_labels_pod-template-hash": "kubernetes_labels_pod-template-hash",
	"kubernetes_name":                     "kubernetes_name",
	"kubernetes_namespace":                "kubernetes_namespace",
	"kubernetes_observed_generation":      "kubernetes_observed_generation",
	"kubernetes_replicas":                 "kubernetes_replicas",
	"kubernetes_scale_down":               "kubernetes_scale_down",
	"kubernetes_scale_up":                 "kubernetes_scale_up",
	"kubernetes_strategy":                 "kubernetes_strategy",
	"kubernetes_unavailable_replicas":     "kubernetes_unavailable_replicas",
	"kubernetes_updated_replicas":         "kubernetes_updated_replicas",
	"name":              "name",
	"pid":               "pid",
	"pod":               "pod",
	"port":              "port",
	"ppid":              "ppid",
	"procspied":         "procspied",
	"reverse_dns_names": "reverse_dns_names",
	"snooped_dns_names": "snooped_dns_names",
	"threads":           "threads",
}

// lookupCommonKey tris to avoid an allocation in populating the key
// by looking it up
func lookupCommonKey(key *string, b []byte) {
	var ok bool
	if *key, ok = commonKeys[string(b)]; !ok {
		*key = string(b)
	}
}
