# graph-study

## MySQL

### Config
```
$ cd mysql
$ bash build.sh
$ bash run.sh
$ docker ps // obtain the <container_id>
$ bash initdb.sh
Container ID: <container_id>
```

## Neo4j

### Config
```
$ cd neo4j
$ bash build.sh
$ bash run.sh
$ docker ps // obtain the <container_id>
$ bash initdb.sh
Container ID: <container_id>
```

## Queries

### Find All Nodes Recursively Starting From an Id
**SQL**
```sql
WITH RECURSIVE get_vertices AS (
    SELECT
        c.id as vertex_id
         , cp.parent_id
    FROM vertex as c
             LEFT JOIN edge cp on cp.vertex_id = c.id
    WHERE
            c.id = '77d0b81c-612a-4a5a-919f-2f7f6e4de91d'
    UNION ALL
    SELECT
        cp2.vertex_id
         , cp2.parent_id
    FROM edge cp2
             INNER JOIN get_vertices o on o.vertex_id = cp2.parent_id
)
SELECT c.id, c.type, parent_id
FROM get_vertices gc
         JOIN vertex c on c.id = gc.vertex_id
```
Response: Array

**Cypher**
```cypher
MATCH p = (r:Vertex {id: "77d0b81c-612a-4a5a-919f-2f7f6e4de91d"})-[*]->(x)
WHERE NOT ((x)-->())
RETURN nodes(p) AS Vertices, relationships(p) AS Edges
```
Response: Tree