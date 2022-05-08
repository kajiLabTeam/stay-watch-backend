use app;


CREATE TABLE user
(
  id        VARCHAR(100),
  name      VARCHAR(50),
  team      VARCHAR(50)
);

INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', 'ohashi', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-2', 'ukai', 'インタラクション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', 'toyama', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-7436-17873', 'ogane', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-7474-59294', 'kaji', 'Professor');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-99', 'isiguro','B2');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2019-60556', 'Miyagawa-san', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-5', 'kameda', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-22823-42602', 'suzaki', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-200', '電池持ちの検証', 'ロケーション');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-7', 'rui	', 'センシング');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-1-19', 'ayato', 'センシング');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-10', 'maruyama', 'センシング');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-1-121', 'ken', 'M1');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-6', 'akito', 'B4');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-11', 'iwao', 'B3');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-4', 'shamoto', 'B4');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-8', 'terada', 'B4');
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-9', 'fueta', 'B4');

CREATE TABLE log
(
  id int(11) primary key AUTO_INCREMENT,
  room_id  int(11),
  start_at  DATETIME,
  end_at    DATETIME,
  user_id   VARCHAR(100),
  rssi      int(11)
);

INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 08:00:00', '2022-04-30 10:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 11:00:00', '2022-04-30 15:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 13:00:00', '2022-04-30 16:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 14:00:00', '2022-04-30 15:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 14:00:00', '2022-04-30 15:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-7436-17873', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 16:00:00', '2022-04-30 18:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-22823-42602', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-04-30 14:00:00', '2022-04-30 15:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-5', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-05-01 08:00:00', '2022-05-01 10:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 11:00:00', '2022-05-01 15:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (1, '2022-05-01 13:00:00', '2022-05-01 16:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 14:00:00', '2022-05-01 15:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 09:00:00', '2022-05-01 12:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-7436-17873', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 10:00:00', '2022-05-01 13:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-22823-42602', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 8:00:00', '2022-05-01 13:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-5', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 13:00:00', '2022-05-01 15:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-7436-17873', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 15:00:00', '2022-05-01 18:30:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-22823-42602', -60);
INSERT INTO log (room_id, start_at, end_at, user_id, rssi) VALUES (2, '2022-05-01 14:00:00', '2022-05-01 16:00:00', 'e7d61ea3f8dd49c88f2ff2484c07acb9-2021-5', -60);



CREATE TABLE stayer
(
  user_id        VARCHAR(100),
  room_id     int(11),
  rssi      int(11)
);


CREATE TABLE room
(
  id        int(11) primary key AUTO_INCREMENT,
  name      VARCHAR(50)
);

INSERT INTO room (name) VALUES ('梶研-学生部屋');
INSERT INTO room (name) VALUES ('梶研-スマートルーム');
INSERT INTO room (name) VALUES ('梶研-院生室');
INSERT INTO room (name) VALUES ('梶研-FA部屋');
INSERT INTO room (name) VALUES ('梶研-先生部屋');


CREATE TABLE tag
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


CREATE TABLE tag_map
(
  id       int(11) primary key AUTO_INCREMENT,
  user_id  VARCHAR(50),
  tag_id   int(11)
);

INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-2', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-2', 3);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-2', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-3', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-7436-17873', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-7436-17873', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-7436-17873', 10);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-7474-59294', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-7474-59294', 11);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-99', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-99', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2019-60556', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2019-60556', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2019-60556', 10);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-5', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-5', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-5', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-22823-42602', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-22823-42602', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-22823-42602', 9);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-200', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-200', 2);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-7', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-7', 4);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-7', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-1-19', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-1-19', 4);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-1-19', 9);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-10', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-10', 4);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-10', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-1-121', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-1-121', 9);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-6', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-6', 4);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-6', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-11', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-11', 7);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-4', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-4', 4);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-4', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-8', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-8', 4);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-8', 8);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-9', 1);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-9', 3);
INSERT INTO tag_map (user_id, tag_id) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-9', 8);



























