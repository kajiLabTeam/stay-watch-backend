use app;


CREATE TABLE IF NOT EXISTS user
(
  id        VARCHAR(100),
  name      VARCHAR(50),
  team      VARCHAR(50)
);

-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', 'ohashi', 'ロケーション');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-2', 'ukai', 'インタラクション');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', 'toyama', 'ロケーション');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-7436-17873', 'ogane', 'ロケーション');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-7474-59294', 'kaji', 'Professor');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-99', 'isiguro','B2');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2019-60556', 'Miyagawa-san', 'ロケーション');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-5', 'kameda', 'ロケーション');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-22823-42602', 'suzaki', 'ロケーション');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-200', '電池持ちの検証', 'ロケーション');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-7', 'rui	', 'センシング');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-1-19', 'ayato', 'センシング');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-10', 'maruyama', 'センシング');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-1-121', 'ken', 'M1');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-6', 'akito', 'B4');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-11', 'iwao', 'B3');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-4', 'shamoto', 'B4');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-8', 'terada', 'B4');
-- INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-9', 'fueta', 'B4');



INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac00', 'kaji', 'Professor');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac01', 'ogane', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac02', 'Miyagawa-san', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac03', 'ayato', 'センシング');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac04', 'ken', 'M1');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac05', 'suzaki', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac06', 'akito', 'B4');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac07', 'fueta', 'B4');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac08', 'kameda', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac09', 'maruyama', 'センシング');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0a', 'ohashi', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0b', 'rui	', 'センシング');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0c', 'ukai', 'インタラクション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0d', 'terada', 'B4');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0e', 'toyama', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0f', 'shamoto', 'B4');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac10', 'isiguro','B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac11', 'ao','B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac12', 'fuma','B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac13', 'ueji','B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac14', 'oiwa','B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac15', 'togawa','B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac16', 'yada','B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac17', 'yokoyama','B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac18', 'kazuo','B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac19', 'sakai','B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac1a', 'iwaguti','B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac1b', 'makino','B1');








CREATE TABLE IF NOT EXISTS log
(
  id int(11) primary key AUTO_INCREMENT,
  room_id  int(11),
  start_at  DATETIME,
  end_at    DATETIME,
  user_id   VARCHAR(100),
  rssi      int(11)
);

-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 08:00:00', '2022-04-30 10:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 11:00:00', '2022-04-30 15:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 13:00:00', '2022-04-30 16:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 14:00:00', '2022-04-30 15:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 14:00:00', '2022-04-30 15:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-7436-17873', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 16:00:00', '2022-04-30 18:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-22823-42602', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 14:00:00', '2022-04-30 15:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-5', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-05-01 08:00:00', '2022-05-01 10:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 11:00:00', '2022-05-01 15:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-05-01 13:00:00', '2022-05-01 16:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 14:00:00', '2022-05-01 15:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 09:00:00', '2022-05-01 12:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-7436-17873', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 10:00:00', '2022-05-01 13:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-22823-42602', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 8:00:00', '2022-05-01 13:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-5', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 13:00:00', '2022-05-01 15:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-7436-17873', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 15:00:00', '2022-05-01 18:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-22823-42602', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 14:00:00', '2022-05-01 16:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-5', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-12 10:00:00', '2022-05-12 16:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', -60);
-- INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-12 11:00:00', '2022-05-12 16:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', -60);




CREATE TABLE IF NOT EXISTS stayer
(
  user_id        VARCHAR(100),
  room_id     int(11),
  rssi      int(11)
);


CREATE TABLE IF NOT EXISTS room
(
  id        int(11) primary key AUTO_INCREMENT,
  name      VARCHAR(50)
);

INSERT INTO room (name) VALUES ('梶研-学生部屋');
INSERT INTO room (name) VALUES ('梶研-スマートルーム');
INSERT INTO room (name) VALUES ('梶研-院生室');
INSERT INTO room (name) VALUES ('梶研-FA部屋');
INSERT INTO room (name) VALUES ('梶研-先生部屋');


CREATE TABLE IF NOT EXISTS tag
(
  id        int(11) primary key AUTO_INCREMENT,
  name      VARCHAR(50)
);

INSERT INTO tag (name) VALUES ('梶研');
INSERT INTO tag (name) VALUES ('ロケーション');
INSERT INTO tag (name) VALUES ('インタラクション');
INSERT INTO tag (name) VALUES ('センシング');
INSERT INTO tag (name) VALUES ('B1');
INSERT INTO tag (name) VALUES ('B2');
INSERT INTO tag (name) VALUES ('B3');
INSERT INTO tag (name) VALUES ('B4');
INSERT INTO tag (name) VALUES ('M1');
INSERT INTO tag (name) VALUES ('M2');
INSERT INTO tag (name) VALUES ('Professor');


CREATE TABLE IF NOT EXISTS tag_map
(
  id       int(11) primary key AUTO_INCREMENT,
  user_id  VARCHAR(50),
  tag_id   int(11)
);


INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac00', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac00', 11);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac01', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac01', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac01', 10);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac02', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac02', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac02', 10);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac03', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac03', 4);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac03', 9);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac04', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac04', 4);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac04', 9);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac05', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac05', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac05', 9);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac06', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac06', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac06', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac07', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac07', 3);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac07', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac08', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac08', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac08', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac09', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac09', 4);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac09', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0a', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0a', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0a', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0b', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0b', 4);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0b', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0c', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0c', 3);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0c', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0d', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0d', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0d', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0e', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0e', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0e', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0f', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0f', 3);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0f', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac10', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac10', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac10', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac11', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac11', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac12', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac12', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac13', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac13', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac14', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac14', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac15', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac15', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac16', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac16', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac17', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac17', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac18', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac18', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac19', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac19', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac1a', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac1a', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac1b', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac1b', 5);







CREATE TABLE IF NOT EXISTS `attendance` (
  `id` int(11) NOT NULL primary key AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL,
  `date` date NOT NULL,
  `exit` BIT NOT NULL
)





























