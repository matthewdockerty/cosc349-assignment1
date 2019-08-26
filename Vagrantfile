# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/xenial64"

    # Build server
    config.vm.define 'build' do |build|
        build.vm.hostname = "build"
        build.vm.network "private_network", ip: "192.168.2.10"
        build.vm.synced_folder ".", "/vagrant", owner: "vagrant", group: "vagrant", mount_options: ["dmode=777,fmode=777"]

        build.vm.provision 'shell', inline: <<-SHELL
            wget https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz
            tar -C /usr/local -xzf go1.12.9.linux-amd64.tar.gz

            cd /vagrant/webapp
            /usr/local/go/bin/go build server.go
        SHELL
    end

    # NGINX reverse proxy
    config.vm.define 'nginx' do |nginx|
        nginx.vm.hostname = "nginx"
        nginx.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"
        nginx.vm.network "private_network", ip: "192.168.2.11"
        nginx.vm.synced_folder ".", "/vagrant", owner: "vagrant", group: "vagrant", mount_options: ["dmode=555,fmode=555"]

        # TODO: Configure NGINX to start on boot
        nginx.vm.provision 'shell', inline: <<-SHELL
            sudo apt-get update
            sudo apt-get install nginx -y
            sudo service nginx start

           # sudo cp /vagrant/nginx.conf /etc/nginx/nginx.conf
           # sudo chmod 644 /etc/nginx/nginx.conf
           # sudo service nginx restart
        SHELL
    end

    # Go webapp server 1
    config.vm.define 'app1' do |app1|
        app1.vm.hostname = "app1"
        app1.vm.network "private_network", ip: "192.168.2.12"
        app1.vm.synced_folder ".", "/vagrant", owner: "vagrant", group: "vagrant", mount_options: ["dmode=555,fmode=555"]

        app1.vm.provision 'shell', inline: <<-SHELL
            sudo cp /vagrant/app.service /lib/systemd/system/app.service
            sudo echo "Environment=SERVER_NAME=app1" >> /lib/systemd/system/app.service
            sudo systemctl start app
            sudo systemctl enable app
            sudo systemctl status app
        SHELL
    end

    # Go webapp server 2
    config.vm.define 'app2' do |app2|
        app2.vm.hostname = "app2"
        app2.vm.network "private_network", ip: "192.168.2.13"
        app2.vm.synced_folder ".", "/vagrant", owner: "vagrant", group: "vagrant", mount_options: ["dmode=555,fmode=555"]

        app2.vm.provision 'shell', inline: <<-SHELL
            sudo cp /vagrant/app.service /lib/systemd/system/app.service
            sudo echo "Environment=SERVER_NAME=app2" >> /lib/systemd/system/app.service
            sudo systemctl start app
            sudo systemctl enable app
            sudo systemctl status app
        SHELL
    end
end