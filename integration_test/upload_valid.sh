#!/bin/bash

# do chmod +x upload_valid.sh
# ./upload_valid.sh to run upload + retrieve convertion after

# UPLOAD CURL
URL="http://127.0.0.1:8000/audio/user/1/phrase/1"
AUDIO_FILE="./sample.wav"

curl --request POST "$URL" \
     --form "audio_file=@$AUDIO_FILE"

if [ $? -eq 0 ]; then
    echo "Request sent successfully."
else
    echo "Failed to send the request."
fi

# GET CURL
URL="http://127.0.0.1:8000/audio/user/1/phrase/1/ogg"
OUTPUT_FILE="./test_response_file_1_1.ogg"

curl --request GET "$URL" -o "$OUTPUT_FILE"

if [ $? -eq 0 ]; then
  echo "File downloaded successfully: $OUTPUT_FILE"
else
  echo "Failed to download the file."
fi
