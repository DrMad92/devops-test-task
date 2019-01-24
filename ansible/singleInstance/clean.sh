#!/bin/bash

ansible-playbook --ask-vault-pass cleanup.yaml;
echo "${SECONDS} elapsed"