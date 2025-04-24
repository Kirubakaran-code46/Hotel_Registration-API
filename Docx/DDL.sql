CREATE TABLE `middleware_req_log` (
  `Id` bigint NOT NULL AUTO_INCREMENT,
  `request_id` varchar(100) NOT NULL,
  `token` varchar(100) NOT NULL,
  `requesteddate` datetime NOT NULL,
  `realip` varchar(240) DEFAULT NULL,
  `forwardedip` varchar(200) DEFAULT NULL,
  `method` varchar(20) DEFAULT NULL,
  `path` longtext,
  `host` varchar(240) DEFAULT NULL,
  `remoteaddr` varchar(240) DEFAULT NULL,
  `header` longtext,
  `body` longtext,
  `endpoint` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`Id`)
)

CREATE TABLE `middleware_resp_log` (
  `Id` bigint NOT NULL AUTO_INCREMENT,
  `request_id` varchar(100) DEFAULT NULL,
  `response` longtext,
  `responseStatus` int DEFAULT NULL,
  `requesteddate` datetime NOT NULL,
  `realip` varchar(240) DEFAULT NULL,
  `forwardedip` varchar(200) DEFAULT NULL,
  `method` varchar(20) DEFAULT NULL,
  `path` longtext,
  `host` varchar(240) DEFAULT NULL,
  `remoteaddr` varchar(240) DEFAULT NULL,
  `header` longtext,
  `body` longtext,
  `endpoint` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`Id`)
)


CREATE TABLE `role_management` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `role` varchar(30) NOT NULL,
  `Description` text,
  `CreatedBy` varchar(20) DEFAULT NULL,
  `createdDate` date DEFAULT NULL,
  `UpdatedBy` varchar(20) DEFAULT NULL,
  `UpdatedDate` date DEFAULT NULL,
  PRIMARY KEY (`Id`)
)