#!/bin/bash

ctrlAddr=127.0.0.1:2048
# create subscription
grpcurl -d @ -plaintext $ctrlAddr  vanus.core.controller.TriggerController.CreateSubscription <<EOF
{
  "filters": [
      {
        "exact": {
            "type":"test"
        }
      }
  ],
  "sink": "http://vance-display.default",
  "eventbus": ""
}
EOF
# delete subscription
grpcurl -d @ -plaintext $ctrlAddr  vanus.core.controller.TriggerController.DeleteSubscription <<EOF
{
  "id": 123
}
EOF



