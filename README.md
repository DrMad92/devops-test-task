# devops-test-task
Wallester test task

# Description
`webapp`: Simple webapp, that communicates with PostgreSQL database, and allows CREATE and DELETE operations.<br>
`ansible`: Ansible playbooks to deploy webapp and infrastructure on AWS. Infrastructure deployed using Clouformation template, and webapp with docker container.
# Initial setup
Before deploying, change `aws_access_key` and `aws_secret_key` variables in `ansible/vars.yml` file.<br>
To edit file: `ansible-vault edit ansible/vars.yml`<br>
Password: **wallester**

# Start deployment
Execute `ansible/run.sh`