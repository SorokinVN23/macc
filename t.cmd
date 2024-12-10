cd C:\dev\projects\Yandex\macc

$SERVER_PORT = 8080
$ADDRESS = "localhost:$SERVER_PORT"
$TEMP_FILE = "./test.tmp"

.\metricstest-windows-amd64.exe "-test.v" "-test.run" "^TestIteration5$" -agent-binary-path ./cmd/agent/agent.exe -binary-path ./cmd/server/server.exe -server-port $SERVER_PORT -source-path .