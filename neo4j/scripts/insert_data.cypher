CREATE (a:Vertex {id: "c37e04a4-01e6-4b6f-bb34-94bc60dd1495", type: "A"});

CREATE (b:Vertex {id: "77d0b81c-612a-4a5a-919f-2f7f6e4de91d", type: "B"});
CREATE (b:Vertex {id: "4ab50ac3-6ec6-421a-8b27-4687f7bc6572", type: "B"});

MATCH (a:Vertex), (b:Vertex) WHERE a.id = "c37e04a4-01e6-4b6f-bb34-94bc60dd1495" AND b.id = "77d0b81c-612a-4a5a-919f-2f7f6e4de91d" CREATE (a)-[:has]->(b);
MATCH (a:Vertex), (b:Vertex) WHERE a.id = "c37e04a4-01e6-4b6f-bb34-94bc60dd1495" AND b.id = "4ab50ac3-6ec6-421a-8b27-4687f7bc6572" CREATE (a)-[:has]->(b);

CREATE (c:Vertex {id: "bf3e0288-0870-46e5-8e8f-d1053caece66", type: "C"});
CREATE (c:Vertex {id: "a20e2747-6118-407b-b1bf-ba37c0b8c7f5", type: "C"});
CREATE (c:Vertex {id: "55c228cd-1978-4a78-9455-be763f3ee2d6", type: "C"});
CREATE (c:Vertex {id: "31b9fca5-3c23-473e-9d67-aec4a3478e7e", type: "C"});
CREATE (c:Vertex {id: "381dedd3-3b35-42d1-bf09-0ad6a2a5e21e", type: "C"});

MATCH (b:Vertex), (c:Vertex) WHERE b.id = '77d0b81c-612a-4a5a-919f-2f7f6e4de91d' AND c.id = 'bf3e0288-0870-46e5-8e8f-d1053caece66' CREATE (b)-[:has]->(c);
MATCH (b:Vertex), (c:Vertex) WHERE b.id = '77d0b81c-612a-4a5a-919f-2f7f6e4de91d' AND c.id = 'a20e2747-6118-407b-b1bf-ba37c0b8c7f5' CREATE (b)-[:has]->(c);
MATCH (b:Vertex), (c:Vertex) WHERE b.id = '4ab50ac3-6ec6-421a-8b27-4687f7bc6572' AND c.id = '55c228cd-1978-4a78-9455-be763f3ee2d6' CREATE (b)-[:has]->(c);
MATCH (b:Vertex), (c:Vertex) WHERE b.id = '4ab50ac3-6ec6-421a-8b27-4687f7bc6572' AND c.id = '31b9fca5-3c23-473e-9d67-aec4a3478e7e' CREATE (b)-[:has]->(c);
MATCH (b:Vertex), (c:Vertex) WHERE b.id = '4ab50ac3-6ec6-421a-8b27-4687f7bc6572' AND c.id = '381dedd3-3b35-42d1-bf09-0ad6a2a5e21e' CREATE (b)-[:has]->(c);

CREATE (d:Vertex {id: "f6043837-d32a-4a59-9c2f-ce8647a37a01", type: "D"});
CREATE (d:Vertex {id: "ca1d2d28-93ae-4370-b1d1-50ded080a26b", type: "D"});
CREATE (d:Vertex {id: "00f48479-7ce2-48ed-bf1a-f063eb8f7589", type: "D"});
CREATE (d:Vertex {id: "5f916897-007a-4343-9047-bc1aaecb1a02", type: "D"});
CREATE (d:Vertex {id: "35a5592c-b4ff-4178-be59-1598be97a250", type: "D"});
CREATE (d:Vertex {id: "1f935058-9f66-4d31-967f-3e3c0a313602", type: "D"});
CREATE (d:Vertex {id: "eb7974aa-738a-4edf-a555-f1e3c25db1ba", type: "D"});
CREATE (d:Vertex {id: "2d00d5d0-eeda-4e80-a030-d001eff66ac7", type: "D"});
CREATE (d:Vertex {id: "8e457f33-73c3-4ca3-8cb3-3ac577cb86e0", type: "D"});
CREATE (d:Vertex {id: "c69454bc-de12-49e9-98b5-1259516ae473", type: "D"});

MATCH (c:Vertex), (d:Vertex) WHERE c.id = 'bf3e0288-0870-46e5-8e8f-d1053caece66' AND d.id = 'f6043837-d32a-4a59-9c2f-ce8647a37a01' CREATE (c)-[:has]->(d);
MATCH (c:Vertex), (d:Vertex) WHERE c.id = 'bf3e0288-0870-46e5-8e8f-d1053caece66' AND d.id = 'ca1d2d28-93ae-4370-b1d1-50ded080a26b' CREATE (c)-[:has]->(d);
MATCH (c:Vertex), (d:Vertex) WHERE c.id = 'a20e2747-6118-407b-b1bf-ba37c0b8c7f5' AND d.id = '00f48479-7ce2-48ed-bf1a-f063eb8f7589' CREATE (c)-[:has]->(d);
MATCH (c:Vertex), (d:Vertex) WHERE c.id = 'a20e2747-6118-407b-b1bf-ba37c0b8c7f5' AND d.id = '5f916897-007a-4343-9047-bc1aaecb1a02' CREATE (c)-[:has]->(d);
MATCH (c:Vertex), (d:Vertex) WHERE c.id = '55c228cd-1978-4a78-9455-be763f3ee2d6' AND d.id = '35a5592c-b4ff-4178-be59-1598be97a250' CREATE (c)-[:has]->(d);
MATCH (c:Vertex), (d:Vertex) WHERE c.id = '55c228cd-1978-4a78-9455-be763f3ee2d6' AND d.id = '1f935058-9f66-4d31-967f-3e3c0a313602' CREATE (c)-[:has]->(d);
MATCH (c:Vertex), (d:Vertex) WHERE c.id = '31b9fca5-3c23-473e-9d67-aec4a3478e7e' AND d.id = 'eb7974aa-738a-4edf-a555-f1e3c25db1ba' CREATE (c)-[:has]->(d);
MATCH (c:Vertex), (d:Vertex) WHERE c.id = '31b9fca5-3c23-473e-9d67-aec4a3478e7e' AND d.id = '2d00d5d0-eeda-4e80-a030-d001eff66ac7' CREATE (c)-[:has]->(d);
MATCH (c:Vertex), (d:Vertex) WHERE c.id = '381dedd3-3b35-42d1-bf09-0ad6a2a5e21e' AND d.id = '8e457f33-73c3-4ca3-8cb3-3ac577cb86e0' CREATE (c)-[:has]->(d);
MATCH (c:Vertex), (d:Vertex) WHERE c.id = '381dedd3-3b35-42d1-bf09-0ad6a2a5e21e' AND d.id = 'c69454bc-de12-49e9-98b5-1259516ae473' CREATE (c)-[:has]->(d);