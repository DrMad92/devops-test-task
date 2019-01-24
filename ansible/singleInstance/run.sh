#!/bin/bash
echo "Enter vault password"
read -s VAULT_PASS
mkfifo -m 600 ./vault_pass

echo ${VAULT_PASS} > ./vault_pass &
ansible-playbook  infrastructure.yaml --vault-password-file=./vault_pass;
echo ${VAULT_PASS} > ./vault_pass &
ansible-playbook -i ec2_hosts deploy.yaml --vault-password-file=./vault_pass;

rm -f ./vault_pass
echo "${SECONDS} elapsed"