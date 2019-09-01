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

            /usr/local/go/bin/go get go.mongodb.org/mongo-driver

            /vagrant/build.sh
        SHELL
    end

    # NGINX reverse proxy
    config.vm.define 'nginx' do |nginx|
        nginx.vm.hostname = "nginx"
        nginx.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"
        nginx.vm.network "private_network", ip: "192.168.2.11"
        nginx.vm.synced_folder ".", "/vagrant", owner: "vagrant", group: "vagrant", mount_options: ["dmode=555,fmode=555"]

        nginx.vm.provision 'shell', inline: <<-SHELL
            apt-get update
            apt-get install nginx -y
            service nginx enable
            service nginx start
            cp /vagrant/nginx.conf /etc/nginx/nginx.conf
            chmod 644 /etc/nginx/nginx.conf
            service nginx restart
        SHELL
    end

    # MongoDB server
    config.vm.define 'mongo' do |mongo|
        mongo.vm.hostname = "mongo"
        mongo.vm.network "private_network", ip: "192.168.2.14"
        mongo.vm.synced_folder ".", "/vagrant", owner: "vagrant", group: "vagrant", mount_options: ["dmode=555,fmode=555"]

        mongo.vm.provision 'shell', inline: <<-SHELL
            wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | apt-key add -
            echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/4.2 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-4.2.list
            apt-get update
            apt-get install -y mongodb-org
            systemctl enable mongod
            service mongod start
            
            while netstat -lnt | awk '$4 ~ /:27017$/ {exit 1}'; do sleep 1; done

            mongo < /vagrant/data/load.js
        
            cp /vagrant/mongod.conf /etc/mongod.conf
            service mongod restart
        SHELL
    end

    # Go webapp server 1
    config.vm.define 'app1' do |app1|
        app1.vm.hostname = "app1"
        app1.vm.network "private_network", ip: "192.168.2.12"
        app1.vm.synced_folder ".", "/vagrant", owner: "vagrant", group: "vagrant", mount_options: ["dmode=555,fmode=555"]

        app1.vm.provision 'shell', inline: <<-SHELL
            cp /vagrant/app.service /lib/systemd/system/app.service
            echo -e "Environment=SERVER_NAME=app1\nEnvironment=CONTENT_PATH=/vagrant/webapp/static/\nEnvironment=DB_HOST=192.168.2.14:27017" >> /lib/systemd/system/app.service
            systemctl start app
            systemctl enable app
        SHELL
    end

    # Go webapp server 2
    config.vm.define 'app2' do |app2|
        app2.vm.hostname = "app2"
        app2.vm.network "private_network", ip: "192.168.2.13"
        app2.vm.synced_folder ".", "/vagrant", owner: "vagrant", group: "vagrant", mount_options: ["dmode=555,fmode=555"]

        app2.vm.provision 'shell', inline: <<-SHELL
            cp /vagrant/app.service /lib/systemd/system/app.service
            echo -e "Environment=SERVER_NAME=app2\nEnvironment=CONTENT_PATH=/vagrant/webapp/static/\nEnvironment=DB_HOST=192.168.2.14:27017" >> /lib/systemd/system/app.service
            systemctl start app
            systemctl enable app
        SHELL
    end
end