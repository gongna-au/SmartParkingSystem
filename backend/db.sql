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

INSERT INTO `users` (`phone`, `password`) VALUES 
('12378978963', 'mypassword'),
('12378978961', 'mypassword'),
('12378978962', 'mypassword'),
('12378978964', 'mypassword'),
('12378978965', 'mypassword'),
('12378978966', 'mypassword'),
('12378978967', 'mypassword'),
('12378978968', 'mypassword'),
('12378978969', 'mypassword'),
('12378978970', 'mypassword'),
('12378978971', 'mypassword'),
('12378978972', 'mypassword'),
('12378978973', 'mypassword'),
('12378978974', 'mypassword'),
('12378978977', 'mypassword'),
('15002597117', '123'),
('12378978975', 'mypassword');



Select * from `users`;

CREATE TABLE `bank_cards_bound` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `card_number` varchar(20) NOT NULL,
  `card_password` varchar(255) NOT NULL, -- 注意：实际应用中不应直接存储密码，而应存储密码的哈希值
  `bank_name` varchar(100) NOT NULL,
  `expiry_date`  varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `bank_cards_bound` (`user_id`, `card_number`, `card_password`, `bank_name`, `expiry_date`) VALUES 
(1, '4580967032145897', '784589', 'Global Bank', '2025-12-31'),
(1, '5523145801234567', '092388', 'City Finance Bank', '2026-11-30'),
(2, '4532157894561230', '667511', 'National Trust', '2027-10-31'),
(2, '4024007198765432', '889211', 'Pacific Bank', '2028-09-30'),
(3, '4916971532487601', '453111', 'Unity Finance', '2029-08-31'),
(3, '4556737586899855', '786211', 'Heritage Bank', '2030-07-31'),
(1, '4485637482910543', '112411', 'Prosperity Bank', '2031-06-30'),
(2, '4716351123456789', '334511', 'Metro Credit Union', '2032-05-31'),
(3, '4539871234567891', '556711', 'Highland Bank', '2033-04-30'),
(14, '4539871234567852', '785211', 'Highland Bank', '2033-04-30'),
(14, '4539871234565248', '524811', 'Highland Bank', '2033-04-30'),
(16, '4539871234567852', '785211', 'Highland Bank', '2033-04-30'),
(16, '4539871234565248', '524811', 'Highland Bank', '2033-04-30'),
(1, '4024007131234567', '778811', 'Crest Bank', '2034-03-31');


-- 创建停车场信息表
CREATE TABLE `parking_lots` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `imageUrl` VARCHAR(255) NOT NULL,
  `spaces` int(11) NOT NULL,
  `charge` varchar(50) NOT NULL,
  `latitude` decimal(10, 8) NOT NULL,
  `longitude` decimal(11, 8) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `parking_lots` (`name`, `address`, `spaces`, `charge`, `latitude`, `longitude`, `imageUrl`) VALUES 
('Pike Place Market Parking Garage', '1531 Western Ave, Seattle, WA 98101 US', 90, '$8/h', 47.609722, -122.342222, 'https://cdn.pixabay.com/photo/2014/04/02/10/45/parking-304465_1280.png'),
('South Place Market Parking Garage', '3231 Western Ave, Seattle, WA 98198 US', 100, '$7/h', 47.609722, -122.332222, 'https://cdn.pixabay.com/photo/2021/07/02/09/39/cars-6381364_1280.jpg'),
('West Place Market Parking Garage', '1356 Western Ave, Seattle, WA 12101 US', 80, '$6/h', 47.609722, -122.322222, 'https://cdn.pixabay.com/photo/2021/07/02/09/39/cars-6381364_1280.jpg'),
('East Place Market Parking Garage', '2356 Eastern Ave, Seattle, WA 98102 US', 85, '$9/h', 47.609722, -122.312222, 'https://cdn.pixabay.com/photo/2021/07/02/09/39/cars-6381364_1280.jpg'),
('North Place Market Parking Garage', '4356 Northern Ave, Seattle, WA 98103 US', 75, '$5/h', 47.609722, -122.302222, 'https://cdn.pixabay.com/photo/2021/07/02/09/39/cars-6381364_1280.jpg'),
('Central Place Market Parking Garage', '536 Central Ave, Seattle, WA 98104 US', 95, '$10/h', 47.609722, -122.292222, 'https://cdn.pixabay.com/photo/2021/07/02/09/39/cars-6381364_1280.jpg'),
('Harbor Place Market Parking Garage', '636 Harbor Ave, Seattle, WA 98105 US', 85, '$4/h', 47.609722, -122.282222, 'https://cdn.pixabay.com/photo/2016/03/22/00/45/multi-storey-car-park-1271917_1280.jpg'),
('River Place Market Parking Garage', '736 River Ave, Seattle, WA 98106 US', 65, '$3/h', 47.609722, -122.272222, 'https://cdn.pixabay.com/photo/2016/03/22/00/45/multi-storey-car-park-1271917_1280.jpg'),
('Ocean Place Market Parking Garage', '836 Ocean Ave, Seattle, WA 98107 US', 55, '$2/h', 47.609722, -122.262222, 'https://cdn.pixabay.com/photo/2016/03/22/00/45/multi-storey-car-park-1271917_1280.jpg'),
('Mountain Place Market Parking Garage', '936 Mountain Ave, Seattle, WA 98108 US', 45, '$1/h', 47.609722, -122.252222, 'https://cdn.pixabay.com/photo/2016/03/22/00/45/multi-storey-car-park-1271917_1280.jpg');


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

CREATE TABLE parking_history (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    parking_lot_id INT NOT NULL,
    vehicle_number VARCHAR(20) NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL
);

INSERT INTO parking_history (user_id, parking_lot_id, start_time, end_time, vehicle_number) VALUES
(1, 1, '2024-02-20 08:00:00', '2024-02-20 10:30:00', 'ABC123'),
(2, 1, '2024-02-20 09:00:00', '2024-02-20 11:15:00', 'XYZ456'),
(1, 2, '2024-02-21 08:15:00', '2024-02-21 10:45:00', 'DEF789'),
(14, 4, '2024-02-21 07:45:00', '2024-02-21 08:30:00', 'GHI012'),
(14, 8, '2024-02-22 07:45:00', '2024-02-22 08:30:00', 'GHI012'),
(14, 4, '2024-02-23 07:45:00', '2024-02-23 08:30:00', 'GHI012'),
(14, 8, '2024-02-24 07:45:00', '2024-02-24 08:30:00', 'GHI012'),
(14, 4, '2024-02-25 07:45:00', '2024-02-25 08:30:00', 'GHI012'),
(16, 8, '2024-02-24 07:45:00', '2024-02-24 08:30:00', 'EFI012'),
(16, 4, '2024-02-25 07:45:00', '2024-02-25 08:30:00', 'YRI012'),
(14, 8, '2024-02-26 07:45:00', '2024-02-27 08:30:00', 'GHI012');



CREATE TABLE `parking_reservations` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT NOT NULL,
  `parking_lot_id` INT NOT NULL,
  `start_time` DATETIME NOT NULL,
  `end_time` DATETIME NOT NULL,
  `bank_card_id` INT NOT NULL,
  `payment_amount` DECIMAL(10, 2) NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
  FOREIGN KEY (`parking_lot_id`) REFERENCES `parking_lots`(`id`),
  FOREIGN KEY (`bank_card_id`) REFERENCES `bank_cards_bound`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `parking_reservations` (`user_id`, `parking_lot_id`, `start_time`, `end_time`, `bank_card_id`, `payment_amount`) VALUES
(1, 1, '2023-03-01 08:00:00', '2023-03-01 12:00:00', 1, 20.00),
(2, 2, '2023-03-02 09:00:00', '2023-03-02 11:00:00', 2, 15.00),
(3, 3, '2023-03-03 10:00:00', '2023-03-03 14:00:00', 3, 25.00),
(1, 4, '2023-03-04 08:30:00', '2023-03-04 10:30:00', 4, 10.00),
(2, 5, '2023-03-05 09:30:00', '2023-03-05 13:30:00', 5, 18.00),
(3, 1, '2023-03-06 10:30:00', '2023-03-06 12:30:00', 6, 12.00),
(1, 2, '2023-03-07 11:00:00', '2023-03-07 15:00:00', 1, 22.00),
(2, 3, '2023-03-08 12:00:00', '2023-03-08 16:00:00', 2, 24.00),
(3, 4, '2023-03-09 13:00:00', '2023-03-09 17:00:00', 3, 30.00),
(14, 3, '2023-03-05 12:00:00', '2023-03-05 16:00:00', 2, 24.00),
(14, 3, '2023-03-06 12:00:00', '2023-03-06 16:00:00', 2, 24.00),
(14, 3, '2023-03-08 12:00:00', '2023-03-08 16:00:00', 2, 24.00),
(14, 4, '2023-03-09 13:00:00', '2023-03-09 17:00:00', 3, 30.00),
(16, 3, '2023-03-10 12:00:00', '2023-03-06 16:00:00', 2, 24.00),
(16, 3, '2023-03-11 12:00:00', '2023-03-08 16:00:00', 2, 24.00),
(16, 4, '2023-03-12 13:00:00', '2023-03-09 17:00:00', 3, 30.00),
(1, 5, '2023-03-10 14:00:00', '2023-03-10 18:00:00', 4, 16.00);

