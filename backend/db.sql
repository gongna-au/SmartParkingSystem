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


-- 创建停车场信息表
CREATE TABLE `parking_lots` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `spaces` int(11) NOT NULL,
  `charge` varchar(50) NOT NULL,
  `latitude` decimal(10, 8) NOT NULL,
  `longitude` decimal(11, 8) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `parking_lots` (`name`, `address`, `spaces`, `charge`, `latitude`, `longitude`) VALUES 
('Pike Place Market Parking Garage', '1531 Western Ave, Seattle, WA 98101 US', 90, '$8/h', 47.609722, -122.342222),
('South Place Market Parking Garage', '3231 Western Ave, Seattle, WA 98198 US', 100, '$7/h', 47.609722, -122.332222),
('West Place Market Parking Garage', '1356 Western Ave, Seattle, WA 12101 US', 80, '$6/h', 47.609722, -122.322222);


-- 假设用户当前的经纬度为(latitude, longitude)
SET @user_lat = 47.609722;
SET @user_long = -122.342222;

SELECT *, 
       (6371 * acos(cos(radians(@user_lat)) 
                    * cos(radians(latitude)) 
                    * cos(radians(longitude) - radians(@user_long)) 
                    + sin(radians(@user_lat)) 
                    * sin(radians(latitude)))) AS distance 
FROM `parking_lots`
ORDER BY distance ASC;

