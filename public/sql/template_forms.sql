-- MySQL dump 10.13  Distrib 8.0.41, for macos15.2 (arm64)
--
-- Host: 127.0.0.1    Database: workflow
-- ------------------------------------------------------
-- Server version	8.0.41

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `template_forms`
--

DROP TABLE IF EXISTS `template_forms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `template_forms` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `field` varchar(191) NOT NULL DEFAULT '' COMMENT '''表单字段英文名''',
  `field_name` varchar(191) NOT NULL DEFAULT '' COMMENT '''表单字段中文名''',
  `field_type` varchar(191) NOT NULL DEFAULT '' COMMENT '''表单字段类型''',
  `field_value` text COMMENT '''表单字段值，select radio checkbox用''',
  `field_default_value` text COMMENT '''表单字段默认值''',
  `field_rules` longblob,
  `sort` bigint NOT NULL DEFAULT '100' COMMENT '''排序''',
  `template_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '''模板ID''',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `template_forms`
--

LOCK TABLES `template_forms` WRITE;
/*!40000 ALTER TABLE `template_forms` DISABLE KEYS */;
INSERT INTO `template_forms` VALUES (2,'2025-05-13 22:21:37.749','2025-05-13 22:21:37.749','title','标题','text','[]','',_binary '[{\"rule_name\":\"required\",\"rule_title\":\"必填\",\"rule_value\":\"\"}]',100,1),(3,'2025-05-13 22:22:11.030','2025-05-13 22:22:11.030','content','正文','textarea','[]','',_binary '[{\"rule_name\":\"required\",\"rule_title\":\"必填\",\"rule_value\":\"\"}]',1,1),(5,'2025-05-13 22:25:04.661','2025-05-13 22:25:04.661','complete_at','完成时间','date','[]','',_binary '[{\"rule_name\":\"date\",\"rule_title\":\"日期\",\"rule_value\":\"\"}]',2,1);
/*!40000 ALTER TABLE `template_forms` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-05-13 23:53:19
