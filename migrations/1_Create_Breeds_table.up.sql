CREATE TABLE IF NOT EXISTS `Breeds`(
  `Id` int NOT NULL AUTO_INCREMENT,
  `Search` varchar(200) NOT NULL,
  `JsonData` text NOT NULL,
  PRIMARY KEY (`Id`)
);
