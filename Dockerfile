# Build the manager binary
FROM golang:1.16 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY pkg/ pkg/
COPY version/ version/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM registry.access.redhat.com/ubi8:8.5-200 AS base-env

RUN yum --setopt=tsflags=nodocs install -y unzip tar rsync shadow-utils
RUN yum --setopt=tsflags=nodocs install -y java-1.8.0-openjdk-devel
RUN yum --setopt=tsflags=nodocs install -y hostname libaio python2
RUN yum clean all && [ ! -d /var/cache/yum ] || rm -rf /var/cache/yum

WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]

LABEL \
      com.redhat.component="amq-broker-rhel8-operator-container"  \
      com.redhat.delivery.appregistry="false" \
      description="ActiveMQ Artemis Broker Operator"  \
      io.k8s.description="An associated operator that handles broker installation, updates and scaling."  \
      io.k8s.display-name="ActiveMQ Artemis Broker Operator"  \
      io.openshift.expose-services=""  \
      io.openshift.s2i.scripts-url="image:///usr/local/s2i"  \
      io.openshift.tags="messaging,amq,integration,operator,golang"  \
      maintainer="Roddie Kieley <rkieley@redhat.com>"  \
      name="amq7/amq-broker-rhel8-operator" \
      summary="ActiveMQ Artemis Broker Operator"  \
      version="0.20"