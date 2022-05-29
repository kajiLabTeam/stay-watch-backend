use app;


CREATE TABLE IF NOT EXISTS user
(
  id        int(11) PRIMARY KEY AUTO_INCREMENT,
  uid     VARCHAR(255),
  name      VARCHAR(50)
);


INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac00', 'kaji');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac01', 'ogane');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac02', 'Miyagawa-san');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac03', 'ayato');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac04', 'ken');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac05', 'suzaki');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac06', 'akito');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac07', 'fueta');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac08', 'kameda');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac09', 'maruyama');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0a', 'ohashi');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0b', 'rui');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0c', 'ukai');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0d', 'terada');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0e', 'toyama');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac0f', 'shamoto');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac10', 'isiguro');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac11', 'ao');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac12', 'fuma');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac13', 'ueji');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac14', 'oiwa');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac15', 'togawa');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac16', 'yada');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac17', 'yokoyama');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac18', 'kazuo');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac19', 'sakai');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac1a', 'iwaguti');
INSERT INTO user (uid, name) VALUES ('e7d61ea3f8dd49c88f2ff2484c07ac1b', 'makino');



CREATE TABLE IF NOT EXISTS log
(
  id int(11) primary key AUTO_INCREMENT,
  room_id  int(11),
  start_at  DATETIME,
  end_at    DATETIME,
  user_id   int(11),
  rssi      int(11)
);



CREATE TABLE IF NOT EXISTS stayer
(
  user_id        int(11),
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
  user_id  int(11),
  tag_id   int(11)
);


INSERT INTO tag_map (user_id, tag_id) VALUES (1, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (1, 11);
INSERT INTO tag_map (user_id, tag_id) VALUES (2, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (2, 2);
INSERT INTO tag_map (user_id, tag_id) VALUES (2, 10);
INSERT INTO tag_map (user_id, tag_id) VALUES (3, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (3, 2);
INSERT INTO tag_map (user_id, tag_id) VALUES (3, 10);
INSERT INTO tag_map (user_id, tag_id) VALUES (4, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (4, 4);
INSERT INTO tag_map (user_id, tag_id) VALUES (4, 9);
INSERT INTO tag_map (user_id, tag_id) VALUES (5, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (5, 4);
INSERT INTO tag_map (user_id, tag_id) VALUES (5, 9);
INSERT INTO tag_map (user_id, tag_id) VALUES (6, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (6, 2);
INSERT INTO tag_map (user_id, tag_id) VALUES (6, 9);
INSERT INTO tag_map (user_id, tag_id) VALUES (7, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (7, 2);
INSERT INTO tag_map (user_id, tag_id) VALUES (7, 8);
INSERT INTO tag_map (user_id, tag_id) VALUES (8, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (8, 3);
INSERT INTO tag_map (user_id, tag_id) VALUES (8, 8);
INSERT INTO tag_map (user_id, tag_id) VALUES (9, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (9, 2);
INSERT INTO tag_map (user_id, tag_id) VALUES (9, 8);
INSERT INTO tag_map (user_id, tag_id) VALUES (10, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (10, 4);
INSERT INTO tag_map (user_id, tag_id) VALUES (10, 8);
INSERT INTO tag_map (user_id, tag_id) VALUES (11, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (11, 2);
INSERT INTO tag_map (user_id, tag_id) VALUES (11, 8);
INSERT INTO tag_map (user_id, tag_id) VALUES (12, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (12, 4);
INSERT INTO tag_map (user_id, tag_id) VALUES (12, 8);
INSERT INTO tag_map (user_id, tag_id) VALUES (13, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (13, 3);
INSERT INTO tag_map (user_id, tag_id) VALUES (13, 8);
INSERT INTO tag_map (user_id, tag_id) VALUES (14, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (14, 2);
INSERT INTO tag_map (user_id, tag_id) VALUES (15, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (15, 8);
INSERT INTO tag_map (user_id, tag_id) VALUES (16, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (16, 2);
INSERT INTO tag_map (user_id, tag_id) VALUES (16, 8);
INSERT INTO tag_map (user_id, tag_id) VALUES (17, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (17, 3);
INSERT INTO tag_map (user_id, tag_id) VALUES (17, 8);
INSERT INTO tag_map (user_id, tag_id) VALUES (18, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (18, 7);
INSERT INTO tag_map (user_id, tag_id) VALUES (18, 8);
INSERT INTO tag_map (user_id, tag_id) VALUES (19, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (19, 7);
INSERT INTO tag_map (user_id, tag_id) VALUES (20, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (20, 7);
INSERT INTO tag_map (user_id, tag_id) VALUES (21, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (21, 7);
INSERT INTO tag_map (user_id, tag_id) VALUES (22, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (22, 7);
INSERT INTO tag_map (user_id, tag_id) VALUES (23, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (23, 7);
INSERT INTO tag_map (user_id, tag_id) VALUES (24, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (24 ,7);
INSERT INTO tag_map (user_id, tag_id) VALUES (25, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (25, 7);
INSERT INTO tag_map (user_id, tag_id) VALUES (26, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (26, 7);
INSERT INTO tag_map (user_id, tag_id) VALUES (27, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (27, 7);
INSERT INTO tag_map (user_id, tag_id) VALUES (28, 1);
INSERT INTO tag_map (user_id, tag_id) VALUES (28, 5);




CREATE TABLE IF NOT EXISTS attendance (
  id int(11) primary key AUTO_INCREMENT,
  user_id int(11),
  date date ,
  flag int(11)
);



CREATE TABLE IF NOT EXISTS attendance_tmp (  
  user_id int(11),
  flag int(11)
);

INSERT INTO attendance_tmp (user_id, flag) VALUES (1, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (2, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (3, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (4, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (5, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (6, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (7, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (8, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (9, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (10, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (11, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (12, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (13, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (14, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (15, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (16, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (17, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (18, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (19, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (20, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (21, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (22, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (23, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (24, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (25, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (26, 0);
INSERT INTO attendance_tmp (user_id, flag) VALUES (27, 0);





































