#!/bin/bash
ansible-playbook --ask-vault-pass infrastructure.yaml;
ansible-playbook -i ec2_hosts deploy.yaml;
echo "${SECONDS} elapsed"