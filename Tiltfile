# Tiltfile for go-microservices
# Builds go services, creates images and applies k8s manifests found in k8s/
# Uses the restart_process extension to support fast rebuilds during development.

load('ext://restart_process', 'docker_build_with_restart')

# Apply all k8s manifests in the k8s/ folder
# Tilt's Starlark doesn't expose a glob helper by default, so use a local shell
# to expand the pattern and pass the list of files to k8s_yaml.
yaml_blob = local('bash -lc "ls k8s/*.yaml 2>/dev/null || true"')
# local() may return a CmdResult with a stdout field or a plain blob/string depending on Tilt version.
yaml_list = str(yaml_blob).split()
if len(yaml_list) == 0:
  print("Warning: no k8s manifests found in k8s/; continuing without applying k8s YAMLs")
k8s_yaml(yaml_list)

# ------------------- Auth Service -------------------
auth_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/auth-service ./services/auth-service/cmd'
local_resource(
  'auth-service-compile',
  auth_compile_cmd,
  deps=['./services/auth-service'],
  labels='compiles')

docker_build_with_restart(
  'auth-service',
  '.',
  dockerfile='services/auth-service/Dockerfile',
  entrypoint=['/usr/local/bin/auth-service'],
  only=[
    './build/auth-service',
    './services/auth-service',
  ],
  live_update=[
    sync('./build', '/usr/local/bin'),
  ],
)

k8s_resource('auth-service', port_forwards=[50051], resource_deps=['auth-service-compile'], labels='services')

# ------------------- API Gateway -------------------
api_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/api-gateway ./api-gateway/cmd'
local_resource(
  'api-gateway-compile',
  api_compile_cmd,
  deps=['./api-gateway', './proto'],
  labels='compiles')

docker_build_with_restart(
  'api-gateway',
  '.',
  dockerfile='api-gateway/Dockerfile',
  entrypoint=['/app/main'],
  only=[
    './build/api-gateway',
    './api-gateway',
    './proto',
  ],
  live_update=[
    sync('./build', '/app'),
    sync('./proto', '/app/proto'),
  ],
)

k8s_resource('api-gateway', port_forwards=[8080], resource_deps=['api-gateway-compile'], labels='services')

# ------------------- Post Service -------------------
post_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/post-service ./services/post-service/cmd'
local_resource('post-service-compile', post_compile_cmd, deps=['./services/post-service'], labels='compiles')

docker_build_with_restart(
  'post-service',
  '.',
  dockerfile='services/post-service/Dockerfile',
  entrypoint=['/usr/local/bin/post-service'],
  only=[
    './build/post-service',
    './services/post-service',
  ],
  live_update=[
    sync('./build', '/usr/local/bin'),
  ],
)

k8s_resource('post-service', resource_deps=['post-service-compile'], labels='services')

# ------------------- User Service -------------------
user_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/user-service ./services/user-service/cmd'
local_resource('user-service-compile', user_compile_cmd, deps=['./services/user-service'], labels='compiles')

docker_build_with_restart(
  'user-service',
  '.',
  dockerfile='services/user-service/Dockerfile',
  entrypoint=['/usr/local/bin/user-service'],
  only=[
    './build/user-service',
    './services/user-service',
  ],
  live_update=[
    sync('./build', '/usr/local/bin'),
  ],
)

k8s_resource('user-service', resource_deps=['user-service-compile'], labels='services')

# ------------------- Postgres (port forward) -------------------
k8s_resource('postgres', port_forwards=[5432], labels='datastore')

# ------------------- Other tooling forwards (optional) -------------------
# Forward common tooling ports if present in the k8s manifests
k8s_resource('prometheus', port_forwards=[9090], labels='monitoring')
k8s_resource('grafana', port_forwards=[3000], labels='monitoring')
k8s_resource('jaeger', port_forwards=[16686], labels='tracing')

# End of Tiltfile
