-- MySQL dump 10.13  Distrib 8.0.21, for Win64 (x86_64)
--
-- Host: localhost    Database: testapi
-- ------------------------------------------------------
-- Server version	8.0.21

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
-- Table structure for table `tb_pekerjaan`
--

DROP TABLE IF EXISTS `tb_pekerjaan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_pekerjaan` (
  `id_pekerjaan` int NOT NULL AUTO_INCREMENT,
  `pekerjaan` varchar(50) DEFAULT NULL,
  `status` varchar(5) DEFAULT 'A',
  PRIMARY KEY (`id_pekerjaan`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_pekerjaan`
--

LOCK TABLES `tb_pekerjaan` WRITE;
/*!40000 ALTER TABLE `tb_pekerjaan` DISABLE KEYS */;
INSERT INTO `tb_pekerjaan` VALUES (1,'PNS','A'),(2,'Wirausaha','A'),(3,'Wiraswasta','A');
/*!40000 ALTER TABLE `tb_pekerjaan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_pendidikan`
--

DROP TABLE IF EXISTS `tb_pendidikan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_pendidikan` (
  `id_pendidikan` int NOT NULL AUTO_INCREMENT,
  `pendidikan` varchar(50) DEFAULT NULL,
  `status` varchar(5) DEFAULT 'A',
  PRIMARY KEY (`id_pendidikan`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_pendidikan`
--

LOCK TABLES `tb_pendidikan` WRITE;
/*!40000 ALTER TABLE `tb_pendidikan` DISABLE KEYS */;
INSERT INTO `tb_pendidikan` VALUES (1,'SD','A'),(2,'SMP','A'),(3,'SMA','A'),(4,'Diploma','A'),(5,'Sarjana','A');
/*!40000 ALTER TABLE `tb_pendidikan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_user`
--

DROP TABLE IF EXISTS `tb_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_user` (
  `id_user` int NOT NULL AUTO_INCREMENT,
  `nama` varchar(150) NOT NULL,
  `tanggal_lahir` varchar(50) DEFAULT NULL,
  `no_ktp` varchar(20) DEFAULT NULL,
  `id_pekerjaan` int DEFAULT NULL,
  `id_pendidikan` int DEFAULT NULL,
  `status` varchar(5) DEFAULT 'A',
  PRIMARY KEY (`id_user`),
  KEY `id_pekerjaan` (`id_pekerjaan`),
  KEY `id_pendidikan` (`id_pendidikan`),
  CONSTRAINT `tb_user_ibfk_1` FOREIGN KEY (`id_pekerjaan`) REFERENCES `tb_pekerjaan` (`id_pekerjaan`),
  CONSTRAINT `tb_user_ibfk_2` FOREIGN KEY (`id_pendidikan`) REFERENCES `tb_pendidikan` (`id_pendidikan`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_user`
--

LOCK TABLES `tb_user` WRITE;
/*!40000 ALTER TABLE `tb_user` DISABLE KEYS */;
INSERT INTO `tb_user` VALUES (3,'Agus','1999-15-11','3573331221000022',1,2,'A'),(4,'Budi','1999-10-12','357003223400001',2,4,'A'),(8,'Sujarwo','1999-10-05','357003223400002',3,4,'A');
/*!40000 ALTER TABLE `tb_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-03-01 12:42:59
