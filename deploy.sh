#!/bin/sh

## Validate valid usage
case "$1" in
        dev)
            declare env=dev
            declare -a ips=("34.245.108.85")
            ;;
        prod)
            declare env=prod
            declare -a ips=("18.130.254.226")
            ;;         
        *)
            echo $"Usage: $0 {dev|prod}"
            exit 1 
esac

## Define different upload procedures

upload_passport_game() {
  for host in "${ips[@]}"
  do
    rsync -Pcav --exclude dev --exclude prod -e "ssh -i ~/.ssh/passportpower-key-$env-ec2" /mnt/d/go/src/eman/passport/game/bin/ ubuntu@$host:/var/www/services/games;
    rsync -Pcav -e "ssh -i ~/.ssh/passportpower-key-$env-ec2" /mnt/d/go/src/eman/passport/game/bin/$env/ ubuntu@$host:/var/www/services/games;
  done
}

upload_passport_daily() {
  for host in "${ips[@]}"
  do
    rsync -Pcav --exclude dev --exclude prod -e "ssh -i ~/.ssh/passportpower-key-$env-ec2" /mnt/d/go/src/eman/passport/daily/bin/ ubuntu@$host:/var/www/services/launch;
    rsync -Pcav -e "ssh -i ~/.ssh/passportpower-key-$env-ec2" /mnt/d/go/src/eman/passport/daily/bin/$env/ ubuntu@$host:/var/www/services/launch;
  done
}

upload_passport_sync() {
  for host in "${ips[@]}"
  do
    rsync -Pcav --exclude dev --exclude prod -e "ssh -i ~/.ssh/passportpower-key-$env-ec2" /mnt/d/go/src/eman/passport/sync/bin/ ubuntu@$host:/var/www/services/sync
    rsync -Pcav -e "ssh -i ~/.ssh/passportpower-key-$env-ec2" /mnt/d/go/src/eman/passport/sync/bin/$env/.env ubuntu@$host:/var/www/services/sync;
  done
}


## Decide if building or only uploading
BUILD=1
while getopts "u" OPTION; do
    case $OPTION in
    u)
        BUILD=0 #Upload only
        ;;
    esac
done

if [ $BUILD == 1 ]; then
   ## Build everything
   make -C /mnt/d/go/src/eman/passport/game
   make -C /mnt/d/go/src/eman/passport/daily
   make -C /mnt/d/go/src/eman/passport/sync
fi

# Upload new services
upload_passport_game
upload_passport_daily
upload_passport_sync

# Restart services
ssh -i ~/.ssh/passportpower-key-$env-ec2 ubuntu@$host sudo supervisorctl restart all

