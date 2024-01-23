drop database if exists `backend`;

create database `backend`;

use  `backend`;

-- 创建用户表
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `phone` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `users` ( `phone`, `password`) VALUES 
('12378978963', 'mypassword'),
( '12378978961', 'mypassword'),
( '12378978962', 'mypassword');


Select * from `users`;
