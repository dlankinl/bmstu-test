go-test-command:
	go test -v ./...

test:
	rm -rf allure-results
	export ALLURE_OUTPUT_PATH="/Users/dmitry/Desktop/bmstu/7sem/bmstu-test/backend" && go test ./... --race --parallel 8
	cp environment.properties allure-results

allure:
	[ -d allure-reports ] && cp -rf allure-reports/history allure-results || :
	rm -rf allure-reports
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

report: test allure

.PHONY: go-test-command test allure report