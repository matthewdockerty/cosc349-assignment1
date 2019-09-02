echo "Connecting to build server..."
vagrant ssh build -c "exec /vagrant/build.sh"

if [ $? == 0 ]; then
    echo "Connecting to app1 server & restarting service..."
    vagrant ssh app1 -c "sudo systemctl restart app"
    echo "Connecting to app2 server & restarting service..."
    vagrant ssh app2 -c "sudo systemctl restart app"
fi

