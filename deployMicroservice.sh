#!/bin/zsh

gcloud functions deploy "$1" --runtime go116 --trigger-http --allow-unauthenticated