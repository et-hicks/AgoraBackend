#!/bin/zsh

#gcloud functions deploy "$1" --runtime go116 --trigger-http --allow-unauthenticated
gcloud run deploy "$1" --source . --region=us-central1