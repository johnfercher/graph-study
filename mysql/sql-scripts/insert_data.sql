USE Graph;

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;

INSERT INTO vertex (`id`, `type`) VALUES ('c37e04a4-01e6-4b6f-bb34-94bc60dd1495', 'A');

INSERT INTO vertex (`id`, `type`) VALUES ('77d0b81c-612a-4a5a-919f-2f7f6e4de91d', 'B');
INSERT INTO vertex (`id`, `type`) VALUES ('4ab50ac3-6ec6-421a-8b27-4687f7bc6572', 'B');

INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('c37e04a4-01e6-4b6f-bb34-94bc60dd1495', '77d0b81c-612a-4a5a-919f-2f7f6e4de91d');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('c37e04a4-01e6-4b6f-bb34-94bc60dd1495', '4ab50ac3-6ec6-421a-8b27-4687f7bc6572');

INSERT INTO vertex (`id`, `type`) VALUES ('bf3e0288-0870-46e5-8e8f-d1053caece66', 'C');
INSERT INTO vertex (`id`, `type`) VALUES ('a20e2747-6118-407b-b1bf-ba37c0b8c7f5', 'C');
INSERT INTO vertex (`id`, `type`) VALUES ('55c228cd-1978-4a78-9455-be763f3ee2d6', 'C');
INSERT INTO vertex (`id`, `type`) VALUES ('31b9fca5-3c23-473e-9d67-aec4a3478e7e', 'C');
INSERT INTO vertex (`id`, `type`) VALUES ('381dedd3-3b35-42d1-bf09-0ad6a2a5e21e', 'C');

INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('77d0b81c-612a-4a5a-919f-2f7f6e4de91d', 'bf3e0288-0870-46e5-8e8f-d1053caece66');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('77d0b81c-612a-4a5a-919f-2f7f6e4de91d', 'a20e2747-6118-407b-b1bf-ba37c0b8c7f5');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('4ab50ac3-6ec6-421a-8b27-4687f7bc6572', '55c228cd-1978-4a78-9455-be763f3ee2d6');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('4ab50ac3-6ec6-421a-8b27-4687f7bc6572', '31b9fca5-3c23-473e-9d67-aec4a3478e7e');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('4ab50ac3-6ec6-421a-8b27-4687f7bc6572', '381dedd3-3b35-42d1-bf09-0ad6a2a5e21e');

INSERT INTO vertex (`id`, `type`) VALUES ('f6043837-d32a-4a59-9c2f-ce8647a37a01', 'D');
INSERT INTO vertex (`id`, `type`) VALUES ('ca1d2d28-93ae-4370-b1d1-50ded080a26b', 'D');
INSERT INTO vertex (`id`, `type`) VALUES ('00f48479-7ce2-48ed-bf1a-f063eb8f7589', 'D');
INSERT INTO vertex (`id`, `type`) VALUES ('5f916897-007a-4343-9047-bc1aaecb1a02', 'D');
INSERT INTO vertex (`id`, `type`) VALUES ('35a5592c-b4ff-4178-be59-1598be97a250', 'D');
INSERT INTO vertex (`id`, `type`) VALUES ('1f935058-9f66-4d31-967f-3e3c0a313602', 'D');
INSERT INTO vertex (`id`, `type`) VALUES ('eb7974aa-738a-4edf-a555-f1e3c25db1ba', 'D');
INSERT INTO vertex (`id`, `type`) VALUES ('2d00d5d0-eeda-4e80-a030-d001eff66ac7', 'D');
INSERT INTO vertex (`id`, `type`) VALUES ('8e457f33-73c3-4ca3-8cb3-3ac577cb86e0', 'D');
INSERT INTO vertex (`id`, `type`) VALUES ('c69454bc-de12-49e9-98b5-1259516ae473', 'D');

INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('bf3e0288-0870-46e5-8e8f-d1053caece66', 'f6043837-d32a-4a59-9c2f-ce8647a37a01');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('bf3e0288-0870-46e5-8e8f-d1053caece66', 'ca1d2d28-93ae-4370-b1d1-50ded080a26b');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('a20e2747-6118-407b-b1bf-ba37c0b8c7f5', '00f48479-7ce2-48ed-bf1a-f063eb8f7589');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('a20e2747-6118-407b-b1bf-ba37c0b8c7f5', '5f916897-007a-4343-9047-bc1aaecb1a02');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('55c228cd-1978-4a78-9455-be763f3ee2d6', '35a5592c-b4ff-4178-be59-1598be97a250');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('55c228cd-1978-4a78-9455-be763f3ee2d6', '1f935058-9f66-4d31-967f-3e3c0a313602');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('31b9fca5-3c23-473e-9d67-aec4a3478e7e', 'eb7974aa-738a-4edf-a555-f1e3c25db1ba');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('31b9fca5-3c23-473e-9d67-aec4a3478e7e', '2d00d5d0-eeda-4e80-a030-d001eff66ac7');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('381dedd3-3b35-42d1-bf09-0ad6a2a5e21e', '8e457f33-73c3-4ca3-8cb3-3ac577cb86e0');
INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('381dedd3-3b35-42d1-bf09-0ad6a2a5e21e', 'c69454bc-de12-49e9-98b5-1259516ae473');