#!/bin/sh
sudo rsync -Pcav -e "ssh -i ~/.ssh/pp.dev" /mnt/d/go/src/eman/passport/daily/bin/dev/ ubuntu@35.178.108.73:/var/www/services/launch;
ssh -i ~/.ssh/pp.dev ubuntu@35.178.108.73 sudo supervisorctl restart all