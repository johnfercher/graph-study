#!/bin/bash

read -p 'Container ID: ' containerId
echo
echo Loading data into $containerId.

cat scripts/insert_data.cypher | docker exec -i $containerId ./bin/cypher-shell -u neo4j -p supersecret
cat scripts/insert_data2.cypher | docker exec -i $containerId ./bin/cypher-shell -u neo4j -p supersecret
