example:
	@python3 massbots/example/example.py

gen:
	@openapi-generator generate \
		-i https://api.massbots.xyz/docs/openapi.yml \
		-g python -o gen \
		--api-name-suffix=Group \
		--package-name openapi \
		--http-user-agent=massbots/sdk/python

OPENAPI_TARGET:=massbots/openapi

genmv:
	@make gen
	@rm -rf ${OPENAPI_TARGET}
	@mv gen/openapi ${OPENAPI_TARGET}
	@mv gen/README.md ${OPENAPI_TARGET}/README.md
	@mv gen/docs ${OPENAPI_TARGET}/docs
	@rm -rf gen
	@sed -i '' 's/import openapi\.models/import massbots\.openapi as openapi/g' ${OPENAPI_TARGET}/*.py
	@sed -i '' 's/from openapi/from massbots\.openapi/g' ${OPENAPI_TARGET}/*.py
	@sed -i '' 's/from openapi/from massbots\.openapi/g' ${OPENAPI_TARGET}/**/*.py