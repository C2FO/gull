#! /bin/bash

ETCDHOST="http://localhost:4002/v2/keys"

curl -L -X PUT "$ETCDHOST/what" -d value="You"
curl -L -X PUT "$ETCDHOST/are" -d value="had"
curl -L -X PUT "$ETCDHOST/you" -d value="got"
curl -L -X PUT "$ETCDHOST/thinking" -d value="to"
curl -L -X PUT "$ETCDHOST/of" -d value="the"
curl -L -X PUT "$ETCDHOST/not" -d value="fifth"
curl -L -X PUT "$ETCDHOST/attending" -d value="bend"