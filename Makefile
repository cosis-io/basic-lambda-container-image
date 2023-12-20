SHELL=/bin/bash
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS = -uec

ifndef IMAGE_TAG
	IMAGE_TAG := $(shell git rev-parse --short HEAD | awk '{$$1=$$1};1')
endif

.PHONY: build-local
build-local:
	docker buildx build \
	--platform linux/arm64 \
	--build-arg=AWS_REGION=$${AWS_REGION} \
  --build-arg=ENVIRONMENT=$${ENVIRONMENT} \
  --load --target=prod -t sch00l.io/$${ENVIRONMENT}-<lambda-handler> .

.PHONY: build-image
build-image:
	docker buildx build \
	--platform linux/arm64 --load \
	--build-arg=AWS_REGION=$${AWS_REGION} \
	--build-arg=ENVIRONMENT=$${ENVIRONMENT} \
	--target=prod \
	-t $${REGISTRY}/$${REPOSITORY}:$${ENVIRONMENT}-<lambda-handler>-$${IMAGE_TAG} .

.PHONY: push-image
push-image: build-image
	docker push \
  $${REGISTRY}/$${REPOSITORY}:$${ENVIRONMENT}-<lambda-handler>-$${IMAGE_TAG}

.PHONY: test-image-build
test-image-build:
	docker run -d -v ~/.aws-lambda-rie:/aws-lambda -p 8800:8080 \
  --entrypoint /aws-lambda/aws-lambda-rie \
  sch00l.io/$${ENVIRONMENT}-<lambda-handler>:latest \
      /usr/ci-svc-usr/main

.PHONY: test-image-response
test-image-response:
	curl -X POST -H "Content-Type: application/json" \
  	-d '{ "item1": true, "item2": "12345abc", "item3": "1d28cf980738850c7ba914eef4ee" }' \
  	http://localhost:8800/2015-03-31/functions/function/invocations
