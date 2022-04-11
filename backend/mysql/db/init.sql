use production_db;


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
INSERT INTO user (id, name, team) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-36624-25369', 'takizawa', 'インタラクション');
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

CREATE TABLE log
(
  id int(11) primary key AUTO_INCREMENT,
  room_id  int(11),
  start_at  DATETIME,
  end_at    DATETIME,
  user_id   VARCHAR(100),
  rssi      int(11)
);


CREATE TABLE stayer
(
  user_id        VARCHAR(100),
  room_id     int(11),
  rssi      int(11)
);

INSERT INTO stayer (user_id, room_id, rssi) VALUES ('e7d61ea3f8dd49c88f2ff2484c07acb9-2021-1', 1, -50);

CREATE TABLE room
(
  id        int(11) primary key AUTO_INCREMENT,
  name      VARCHAR(50)
);

INSERT INTO room (name) VALUES ('梶研究室-学生部屋');
INSERT INTO room (name) VALUES ('梶研究室-スマートルーム');
INSERT INTO room (name) VALUES ('梶研究室-院生室');
INSERT INTO room (name) VALUES ('梶研究室-FA部屋');
INSERT INTO room (name) VALUES ('梶研究室-先生部屋');