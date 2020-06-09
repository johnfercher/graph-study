# GraphStudy

This repository contains a graph study, handling specifically trees. Here are two approaches to deal with graphs,
the first solution uses a [MySQL](mysql.com) database, which contains tables that represents vertices and edges,
and the second solution uses a [Neo4j](https://neo4j.com/) database, which is graph based.

The idea is to compare the performance from both approaches to solve common graph problems. There is a REST API which
calls both databases and saves information. The API was written in Golang and here is the [postman collection](docs).

## Data
| Id | Type | ParentId |
|---|---|---|
| c37e04a4-01e6-4b6f-bb34-94bc60dd1495 | A (Root) | NULL |
| 77d0b81c-612a-4a5a-919f-2f7f6e4de91d | B | c37e04a4-01e6-4b6f-bb34-94bc60dd1495 |
| 4ab50ac3-6ec6-421a-8b27-4687f7bc6572 | B | c37e04a4-01e6-4b6f-bb34-94bc60dd1495 |
| ... | ... | ... |

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

### Run
```
$ bash run.sh
```

### Stop
```
$ docker ps // obtain the <container_id>
$ docker stop <container_id>
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

### Run
```
$ bash run.sh
```

### Stop
```
$ docker ps // obtain the <container_id>
$ docker stop <container_id>
```

## Queries

### Recursive Tree
**SQL**
```sql
WITH RECURSIVE get_vertices AS (
    SELECT
        c.id as vertex_id
         , cp.parent_id
    FROM vertex as c
             LEFT JOIN edge cp on cp.vertex_id = c.id
    WHERE
            c.id = 'c37e04a4-01e6-4b6f-bb34-94bc60dd1495'
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
MATCH p = (r:Vertex {id: "c37e04a4-01e6-4b6f-bb34-94bc60dd1495"})-[:has]->(x)
WHERE NOT ((x)-->())
RETURN nodes(p) AS Vertices, relationships(p) AS Edges
```
Response: Tree

### Outgoing Relations
**SQL**
```
TODO
```
Response: Number

**Cypher**
```cypher
MATCH (v:Vertex)-[r]->() WHERE v.id = 'c37e04a4-01e6-4b6f-bb34-94bc60dd1495s' RETURN COUNT(r)
```
Response: Number

## TODO
- Access Neo4j from Golang API inside a docker container