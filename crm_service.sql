-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: localhost    Database: crm_service
-- ------------------------------------------------------
-- Server version	8.0.33

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
-- Table structure for table `actors`
--

DROP TABLE IF EXISTS `actors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `actors` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  `is_verified` enum('false','true') DEFAULT 'false',
  `is_active` enum('false','true') DEFAULT 'false',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `salt` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `fk_actors_role_id` (`role_id`),
  CONSTRAINT `actors_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actors`
--

LOCK TABLES `actors` WRITE;
/*!40000 ALTER TABLE `actors` DISABLE KEYS */;
INSERT INTO `actors` VALUES (1,'superadmin','$2a$10$yY17w5JZZ7H5SQGU7rHJt.hWofM/clQtk9xjJlug7uN.n8Lzqh4QC',1,'true','true','2023-05-30 01:57:31','2023-06-05 08:32:11','5c953b55388c74c95cb89738567a4476'),(2,'newadmin','$2a$10$CE7oGuzzZLaaDVg90wwESuy64v9Uu8iOV5f7RIbwhSi9RpocW9v46',2,'true','true','2023-06-03 08:17:34','2023-06-07 07:30:07','68a3d6deb5ea63af8f60112b2579aa13'),(7,'fourthdadmin','fourthdadmin',2,'false','false','2023-06-04 00:43:15','2023-06-04 00:43:15',NULL),(8,'fifthadmin','fifthadmin',2,'false','false','2023-06-04 00:43:30','2023-06-04 00:43:30',NULL),(9,'sixthadmin','sixthadmin',2,'false','false','2023-06-04 00:43:39','2023-06-04 00:43:39',NULL),(10,'seventhadmin','seventhadmin',2,'false','false','2023-06-04 00:43:49','2023-06-04 00:43:49',NULL),(11,'newestAdmin','newestAdmin',2,'false','false','2023-06-04 00:46:00','2023-06-04 00:46:00',NULL),(16,'fromadmin','fromadmin',2,'false','false','2023-06-04 03:30:08','2023-06-07 06:46:55',NULL),(17,'fromadmin2','fromadmin2',2,'false','false','2023-06-04 03:36:28','2023-06-04 03:36:28',NULL),(18,'hashedadmin','$2a$10$iIPU4Mp6NvWKWEp6XQv5l.MWfBYGk5aaUqAbLF05g9nFCT/VTaMIi',2,'false','false','2023-06-05 08:08:05','2023-06-07 06:47:27','769c4b03fbbef241b65a8d66fe28e28c'),(20,'adminhuhu','$2a$10$X0BQ/pqKIjvjqqkEl351EebidNEObOCg0L1e7zPeM.tliaBCHObQ.',2,'false','false','2023-06-07 04:58:27','2023-06-07 04:58:27','e7f71de9097ae7233e0079eafb789520'),(21,'adminhehe','$2a$10$WFrlXPNhbu14aoWU1r1OLu3YBCdscMpTQnxoHGqQObwJ5kM0/v6JK',2,'false','false','2023-06-07 05:01:40','2023-06-07 05:01:40','8dfe029c03b6428281224a3e1954555f'),(22,'adminhoho','$2a$10$MlFrpbkvsIl4HrQZuAevqelgMIGtbHW25GVJV4fr2IgE7p7.LnfS6',2,'false','false','2023-06-07 05:03:11','2023-06-07 05:03:11','48de333d58aa44e318430f6f75518cc7'),(24,'admini','$2a$10$bDKXSwxS0KIk5u90QW2IWuxx2zuXjbt5y/e46k4NWSJHSgkmDl1au',2,'false','false','2023-06-07 06:47:56','2023-06-07 06:47:56','4e03eec28b3c583fe2c02f0ba0c29f96'),(29,'adminR3','$2a$10$wlyBm69bdWRip1VpJMkfkOXzH4xKd6CBMrsQJXYN.MxcT4C7S4xW2',2,'false','false','2023-06-07 07:30:47','2023-06-07 07:31:01','443be2f1372b8147f306a5f363ddc3df'),(30,'adminR1','$2a$10$J84OB5RIgkExcv.c9fYbfuAe5GmTLDazEu8JIkihjgX4FfHnC2uEO',2,'false','false','2023-06-07 07:31:24','2023-06-07 07:31:24','3c6a55df6c2d5021c8dcd6fd1add58bb');
/*!40000 ALTER TABLE `actors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(100) DEFAULT '',
  `last_name` varchar(100) DEFAULT '',
  `email` varchar(100) NOT NULL,
  `avatar` varchar(255) DEFAULT '',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES (1,'Pascal','Panatagama','pascal@gmail.com','url','2023-06-04 06:49:58','2023-06-04 06:49:58'),(2,'Lionel','Messi','messi@gmail.com','url','2023-06-04 06:50:39','2023-06-04 06:50:39'),(3,'Neymar','Jr','neymar@gmail.com','url','2023-06-04 06:50:53','2023-06-04 06:50:53'),(4,'Vinicius','Jr','vinicius@gmail.com','url','2023-06-04 06:51:04','2023-06-04 06:51:04'),(6,'Michaeliu','Lawson','michael.lawson@reqres.in','url','2023-06-06 21:09:36','2023-06-07 07:17:55');
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `register_approval`
--

DROP TABLE IF EXISTS `register_approval`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `register_approval` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` bigint unsigned DEFAULT NULL,
  `super_admin_id` bigint unsigned DEFAULT NULL,
  `status` varchar(20) DEFAULT 'pending',
  PRIMARY KEY (`id`),
  KEY `fk_actors_id_idx` (`admin_id`),
  CONSTRAINT `fk_actors_id` FOREIGN KEY (`admin_id`) REFERENCES `actors` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `register_approval`
--

LOCK TABLES `register_approval` WRITE;
/*!40000 ALTER TABLE `register_approval` DISABLE KEYS */;
INSERT INTO `register_approval` VALUES (11,29,1,'rejected'),(12,30,NULL,'pending');
/*!40000 ALTER TABLE `register_approval` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `role_name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'superadmin'),(2,'admin');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-07 21:49:07
