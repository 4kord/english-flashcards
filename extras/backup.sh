#!/bin/bash

# Location where you want to keep your db dump
backup_folder_path=~/db_backups


# File name i.e: dump-2020-06-24.sql
file_name="dump-"`date "+%d-%m-%Y_%H_%M_%S"`".sql"


# ensure the location exists
mkdir -p ${backup_folder_path}


#change database name, username and docker container name
dbname=english
username=postgres
container=dad60e8bbf3d

backup_file=${backup_folder_path}/${file_name}

sudo docker exec -it ${container} pg_dump -U ${username} -d ${dbname} > ${backup_file}
../third_party/gdrive upload ${backup_file}
rm ${backup_file}

echo "Dump successful"
