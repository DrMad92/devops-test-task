#!/bin/bash
ansible-playbook --ask-vault-pass autoscale.yaml
echo "${SECONDS} elapsed"