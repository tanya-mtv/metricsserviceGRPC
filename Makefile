.PHONY: buf
buf:
	@echo "[buf] Running buf..."
	@buf generate --path api/metricservice_v1/service.proto
