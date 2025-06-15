#!/bin/sh
rsync -Pcav -e "ssh -i ~/.ssh/passportpower-key-dev-ec2" /mnt/d/go/src/eman/passport/game/bin/dev/ ubuntu@35.178.108.73:/var/www/services/games;
ssh -i ~/.ssh/passportpower-key-dev-ec2 ubuntu@35.178.108.73 sudo supervisorctl restart all