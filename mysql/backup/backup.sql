-- MySQL dump 10.13  Distrib 8.0.25, for Linux (x86_64)
--
-- Host: localhost    Database: app
-- ------------------------------------------------------
-- Server version	8.0.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `attendance`
--

DROP TABLE IF EXISTS `attendance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `attendance` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `date` date DEFAULT NULL,
  `flag` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `attendance`
--

LOCK TABLES `attendance` WRITE;
/*!40000 ALTER TABLE `attendance` DISABLE KEYS */;
/*!40000 ALTER TABLE `attendance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `attendance_tmp`
--

DROP TABLE IF EXISTS `attendance_tmp`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `attendance_tmp` (
  `user_id` int DEFAULT NULL,
  `flag` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `attendance_tmp`
--

LOCK TABLES `attendance_tmp` WRITE;
/*!40000 ALTER TABLE `attendance_tmp` DISABLE KEYS */;
INSERT INTO `attendance_tmp` VALUES (1,0),(2,0),(3,0),(4,0),(5,0),(6,0),(7,0),(8,0),(9,0),(10,0),(11,0),(12,0),(13,0),(14,0),(15,0),(16,0),(17,0),(18,0),(19,0),(20,0),(21,0),(22,0),(23,0),(24,0),(25,0),(26,0),(27,0);
/*!40000 ALTER TABLE `attendance_tmp` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log`
--

DROP TABLE IF EXISTS `log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `room_id` int DEFAULT NULL,
  `start_at` datetime DEFAULT NULL,
  `end_at` datetime DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `rssi` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log`
--

LOCK TABLES `log` WRITE;
/*!40000 ALTER TABLE `log` DISABLE KEYS */;
INSERT INTO `log` VALUES (1,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(2,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(3,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(4,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(5,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(6,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(7,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(8,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(9,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(10,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(11,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(12,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(13,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(14,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(15,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(16,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(17,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(18,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(19,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(20,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(21,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(22,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(23,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(24,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(25,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(26,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(27,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(28,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(29,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(30,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(31,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(32,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(33,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(34,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(35,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(36,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(37,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(38,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(39,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(40,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(41,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(42,1,'2015-01-01 00:00:00','2015-01-01 00:00:00',1,-50),(43,1,'2022-06-06 01:31:48','2022-06-06 01:32:32',2,-50),(44,1,'2022-06-06 01:31:48','2022-06-06 01:32:32',3,-50),(45,1,'2022-06-06 01:31:48','2022-06-06 01:32:32',1,-50),(46,1,'2022-06-02 02:00:00','2022-06-02 12:00:00',1,-50),(47,1,'2022-06-02 09:00:00','2022-06-02 15:00:00',2,-50),(48,1,'2022-06-01 22:00:00','2022-06-02 05:00:00',2,-50),(49,1,'2022-06-02 13:00:00','2022-06-02 15:00:00',2,-50),(50,1,'2022-06-02 01:00:00','2022-06-02 15:00:00',2,-50),(51,2,'2022-06-02 01:00:00','2022-06-02 15:00:00',2,-50),(52,2,'2022-06-02 01:00:00','2022-06-02 15:00:00',3,-50),(53,1,'2022-06-02 01:00:00','2022-06-02 15:00:00',3,-50);
/*!40000 ALTER TABLE `log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `room`
--

DROP TABLE IF EXISTS `room`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `room` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `room`
--

LOCK TABLES `room` WRITE;
/*!40000 ALTER TABLE `room` DISABLE KEYS */;
INSERT INTO `room` VALUES (1,'梶研-学生部屋'),(2,'梶研-スマートルーム'),(3,'梶研-院生室'),(4,'梶研-FA部屋'),(5,'梶研-先生部屋');
/*!40000 ALTER TABLE `room` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stayer`
--

DROP TABLE IF EXISTS `stayer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `stayer` (
  `user_id` int DEFAULT NULL,
  `room_id` int DEFAULT NULL,
  `rssi` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stayer`
--

LOCK TABLES `stayer` WRITE;
/*!40000 ALTER TABLE `stayer` DISABLE KEYS */;
/*!40000 ALTER TABLE `stayer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag`
--

DROP TABLE IF EXISTS `tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tag` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag`
--

LOCK TABLES `tag` WRITE;
/*!40000 ALTER TABLE `tag` DISABLE KEYS */;
INSERT INTO `tag` VALUES (1,'梶研'),(2,'ロケーション'),(3,'インタラクション'),(4,'センシング'),(5,'B1'),(6,'B2'),(7,'B3'),(8,'B4'),(9,'M1'),(10,'M2'),(11,'Professor');
/*!40000 ALTER TABLE `tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag_map`
--

DROP TABLE IF EXISTS `tag_map`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tag_map` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `tag_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag_map`
--

LOCK TABLES `tag_map` WRITE;
/*!40000 ALTER TABLE `tag_map` DISABLE KEYS */;
INSERT INTO `tag_map` VALUES (1,1,1),(2,1,11),(3,2,1),(4,2,2),(5,2,10),(6,3,1),(7,3,2),(8,3,10),(9,4,1),(10,4,4),(11,4,9),(12,5,1),(13,5,4),(14,5,9),(15,6,1),(16,6,2),(17,6,9),(18,7,1),(19,7,2),(20,7,8),(21,8,1),(22,8,3),(23,8,8),(24,9,1),(25,9,2),(26,9,8),(27,10,1),(28,10,4),(29,10,8),(30,11,1),(31,11,2),(32,11,8),(33,12,1),(34,12,4),(35,12,8),(36,13,1),(37,13,3),(38,13,8),(39,14,1),(40,14,2),(41,15,1),(42,15,8),(43,16,1),(44,16,2),(45,16,8),(46,17,1),(47,17,3),(48,17,8),(49,18,1),(50,18,7),(51,18,8),(52,19,1),(53,19,7),(54,20,1),(55,20,7),(56,21,1),(57,21,7),(58,22,1),(59,22,7),(60,23,1),(61,23,7),(62,24,1),(63,24,7),(64,25,1),(65,25,7),(66,26,1),(67,26,7),(68,27,1),(69,27,7),(70,28,1),(71,28,5);
/*!40000 ALTER TABLE `tag_map` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` varchar(255) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'e7d61ea3f8dd49c88f2ff2484c07ac00','kaji'),(2,'e7d61ea3f8dd49c88f2ff2484c07ac01','ogane'),(3,'e7d61ea3f8dd49c88f2ff2484c07ac02','Miyagawa-san'),(4,'e7d61ea3f8dd49c88f2ff2484c07ac03','ayato'),(5,'e7d61ea3f8dd49c88f2ff2484c07ac04','ken'),(6,'e7d61ea3f8dd49c88f2ff2484c07ac05','suzaki'),(7,'e7d61ea3f8dd49c88f2ff2484c07ac06','akito'),(8,'e7d61ea3f8dd49c88f2ff2484c07ac07','fueta'),(9,'e7d61ea3f8dd49c88f2ff2484c07ac08','kameda'),(10,'e7d61ea3f8dd49c88f2ff2484c07ac09','maruyama'),(11,'e7d61ea3f8dd49c88f2ff2484c07ac0a','ohashi'),(12,'e7d61ea3f8dd49c88f2ff2484c07ac0b','rui'),(13,'e7d61ea3f8dd49c88f2ff2484c07ac0c','ukai'),(14,'e7d61ea3f8dd49c88f2ff2484c07ac0d','terada'),(15,'e7d61ea3f8dd49c88f2ff2484c07ac0e','toyama'),(16,'e7d61ea3f8dd49c88f2ff2484c07ac0f','shamoto'),(17,'e7d61ea3f8dd49c88f2ff2484c07ac10','isiguro'),(18,'e7d61ea3f8dd49c88f2ff2484c07ac11','ao'),(19,'e7d61ea3f8dd49c88f2ff2484c07ac12','fuma'),(20,'e7d61ea3f8dd49c88f2ff2484c07ac13','ueji'),(21,'e7d61ea3f8dd49c88f2ff2484c07ac14','oiwa'),(22,'e7d61ea3f8dd49c88f2ff2484c07ac15','togawa'),(23,'e7d61ea3f8dd49c88f2ff2484c07ac16','yada'),(24,'e7d61ea3f8dd49c88f2ff2484c07ac17','yokoyama'),(25,'e7d61ea3f8dd49c88f2ff2484c07ac18','kazuo'),(26,'e7d61ea3f8dd49c88f2ff2484c07ac19','sakai'),(27,'e7d61ea3f8dd49c88f2ff2484c07ac1a','iwaguti'),(28,'e7d61ea3f8dd49c88f2ff2484c07ac1b','makino');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-13  7:02:45
