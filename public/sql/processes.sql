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
-- Table structure for table `processes`
--

DROP TABLE IF EXISTS `processes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `processes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `flow_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '''流程id''',
  `process_name` varchar(191) NOT NULL DEFAULT '' COMMENT '''步骤名称''',
  `limit_time` bigint NOT NULL DEFAULT '0' COMMENT '''限定时间,单位秒''',
  `type` varchar(191) NOT NULL DEFAULT 'operation' COMMENT '''流程图显示操作框类型''',
  `icon` varchar(191) DEFAULT '' COMMENT '''流程图显示图标''',
  `process_to` varchar(191) NOT NULL DEFAULT '',
  `style` text,
  `style_color` varchar(191) NOT NULL DEFAULT '#78a300',
  `style_height` bigint NOT NULL DEFAULT '30',
  `style_width` bigint NOT NULL DEFAULT '30',
  `position_left` varchar(191) NOT NULL DEFAULT '100px',
  `position_top` varchar(191) NOT NULL DEFAULT '200px',
  `position` bigint NOT NULL DEFAULT '1' COMMENT '''步骤位置：1正常步骤2：转入子流程0：第一步 当为2时 child_flow_id child_after child_back_process 可设置''',
  `child_flow_id` bigint NOT NULL DEFAULT '0' COMMENT '''子流程id''',
  `child_after` bigint NOT NULL DEFAULT '2' COMMENT '''子流程结束后 1.同时结束父流程 2.返回父流程''',
  `child_back_process` bigint NOT NULL DEFAULT '0' COMMENT '''子流程结束后返回父流程进程''',
  `description` varchar(191) NOT NULL DEFAULT '' COMMENT '''步骤描述''',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `processes`
--

LOCK TABLES `processes` WRITE;
/*!40000 ALTER TABLE `processes` DISABLE KEYS */;
INSERT INTO `processes` VALUES (1,'2025-05-13 23:46:27.217','2025-05-13 23:46:27.217',1,'新建流程',0,'operation','','','width:200px;height:48px;line-height:30px;color:#66CDAA;left:105px;top:26px;','#78a300',48,200,'105px','26px',1,0,2,0,''),(2,'2025-05-13 23:46:30.882','2025-05-13 23:46:30.882',1,'新建流程',0,'operation','','','width:200px;height:48px;line-height:30px;color:#66CDAA;left:69px;top:18px;','#78a300',48,200,'69px','18px',1,0,2,0,'');
/*!40000 ALTER TABLE `processes` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-05-13 23:55:45
