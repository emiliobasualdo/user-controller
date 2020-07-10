set -e
echo "\033[34mUPLOADING CONFIG FILES\033[0m"
cat config-default.yaml | eb ssh --command 'cat > $HOME/config-default.yaml'
cat config-prod.yaml | eb ssh --command 'cat > $HOME/config-prod.yaml'
echo "\033[34mDEPLOYING\033[0m"
eb deploy