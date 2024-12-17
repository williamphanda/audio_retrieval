#!/bin/bash

# do chmod +x upload_valid.sh
# ./upload_valid.sh to run upload + retrieve convertion after

# UPLOAD CURL
URL="http://127.0.0.1:8000/audio/user/999/phrase/999"
AUDIO_FILE="./sample.wav"

curl --request POST "$URL" \
     --form "audio_file=@$AUDIO_FILE"

if [ $? -eq 0 ]; then
    echo "Request sent successfully."
else
    echo "Failed to send the request."
fi