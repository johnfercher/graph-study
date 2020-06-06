USE RussianDoll;

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;

INSERT INTO russian_doll
(`id`, `name`)
VALUES
('c37e04a4-01e6-4b6f-bb34-94bc60dd1495', '727021724710');

INSERT INTO russian_doll
(`id`, `name`)
VALUES
('77d0b81c-612a-4a5a-919f-2f7f6e4de91d', '151818555780');

INSERT INTO russian_doll_parent
(`parent_id`, `russian_doll_id`)
VALUES
('c37e04a4-01e6-4b6f-bb34-94bc60dd1495', '77d0b81c-612a-4a5a-919f-2f7f6e4de91d')