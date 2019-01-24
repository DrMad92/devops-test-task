# devops-test-task
Wallester test task

# Description
`webapp`: Simple webapp, that communicates with PostgreSQL database, and allows CREATE and DELETE operations.<br>
`ansible/singleInstance`: Ansible playbooks to deploy webapp and infrastructure on AWS. Infrastructure deployed using Clouformation templates, and webapp with docker container.<br>
`ansible/autoscale`: Not finished.

# Initial setup
Before deploying, change credential variables in `creds.yml` file.<br>
To edit file: `ansible-vault edit creds.yml`<br>
Password: **wallester**

# Start deployment
Execute `run.sh`<br>
Outputs Load balancer DNS name to access webapp and EC2 public IP for ssh.

# Cleanup
Execute `clean.sh`<br>
This will remove all created stacks and local files.