set -e
echo "\033[34mUPLOADING CONFIG FILES\033[0m"
cat config-default.yaml | eb ssh --command 'sudo tee /home/webapp/config-default.yaml >/dev/null'
echo "\033[34mDEPLOYING\033[0m"
eb deploy