#!/bin/bash

sudo mysql < sql-scripts/clean_data.sql
sudo mysql < sql-scripts/drop_tables.sql
sudo mysql < sql-scripts/drop_database.sql
sudo mysql < sql-scripts/create_database.sql
sudo mysql < sql-scripts/create_tables.sql
sudo mysql < sql-scripts/insert_data.sql
sudo mysql < sql-scripts/create_user.sql