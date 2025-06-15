#!/bin/sh
rsync -Pcav -e "ssh -i ~/.ssh/passportpower-key-dev-ec2" /mnt/d/go/src/eman/passport/sync/bin/dev/ ubuntu@35.178.108.73:/var/www/services/sync