package pubnub_communicator

import (
	"os"
)

var PUBLISH_KEY = os.Getenv("PUBNUB_PUBLISH_KEY")
var SUBSCRIBE_KEY = os.Getenv("PUBNUB_SUBSCRIBE_KEY")
var SECRET_KEY = os.Getenv("PUBNUB_SECRET_KEY")
var CIPHER_KEY = ""
var USE_SSL = false
