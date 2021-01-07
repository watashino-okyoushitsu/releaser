VERSION=0.0.1

.PHONY: tag
tag:
	git tag -a "v$(VERSION)" -m "Release $(VERSION)"
	git push --tags
