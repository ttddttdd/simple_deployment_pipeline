#  The Simple_deployment_pipeline

This is simple deployment pipeline for Golang application and CockroachDB.

## System Requirements:

    OS: Linux Ubuntu 20.10
    Ansible: 2.9
    Golang: 1.16
    CockroachDB: 20.2.5

## How to deploy the application and CockroachDB

1. Build the application for deployment using Makefile:
   
   > cd application && make build

2. The Ansible playbook deploys the application and CockroachDB on one server, create required certs, 
   keys, databases, tables, load data.
   Run Ansible playbook to deploy on any predefined environments `QA/STAGE/PRODUCTION` 

   > cd ansible && ansible-playbook -i inventory/hosts.ini playbooks/new_service.yaml -e "servers=QA" -Kk -u user
   
   For each environment you can modify specific variables in `ansible/inventory/group_vars` directory

## How to check working application 
   
   After deployment the application available on port 8088 (i.e. 192.168.88.11:8088) 
   and you can see in the browser: 
   
   > Initial balances: 
   > 
   > 1 - 1000 
   >
   > 2 - 250