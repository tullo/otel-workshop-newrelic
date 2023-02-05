go build -o newrelic
env --debug $(cat .env | grep -v '^#') ./newrelic
