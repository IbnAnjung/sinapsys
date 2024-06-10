CREATE DATABASE synapsis;

use synapsis;

-- synapsis.users definition

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `phone_number` varchar(13) NOT NULL,
  `password` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`phone_number`),
  UNIQUE KEY `phone_number` (`phone_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- synapsis.product_categories definition

CREATE TABLE `product_categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- synapsis.products definition

CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `product_category_id` bigint unsigned NOT NULL,
  `name` varchar(50) NOT NULL,
  `description` text,
  `price` double(14,2) NOT NULL DEFAULT '0.00',
  PRIMARY KEY (`id`),
  KEY `idx_product_category_id` (`product_category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- synapsis.carts definition

CREATE TABLE `carts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `product_id` bigint unsigned NOT NULL,
  `quantity` double(8,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`product_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- synapsis.orders definition

CREATE TABLE `orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `created_date` timestamp NOT NULL,
  `expired_date` timestamp NOT NULL,
  `total_price` double(14,2) NOT NULL DEFAULT '0.00',
  `status` tinyint NOT NULL COMMENT '0:created;9:canceled;10:completed',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- synapsis.order_products definition

CREATE TABLE `order_products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_id` bigint unsigned NOT NULL,
  `product_id` bigint unsigned NOT NULL,
  `product_name` varchar(50) NOT NULL,
  `quantity` double(8,2) NOT NULL DEFAULT '0.00',
  `price` double(14,2) NOT NULL DEFAULT '0.00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- synapsis.order_payments definition

CREATE TABLE `order_payments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_id` bigint unsigned NOT NULL,
  `value` double(14,2) DEFAULT NULL,
  `type` tinyint unsigned NOT NULL COMMENT '1:manual transfer',
  `status` tinyint unsigned NOT NULL COMMENT '1:created;9:canceled;10:verified',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- synapsis.payment_manual_transfers definition

CREATE TABLE `payment_manual_transfers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_payment_id` bigint unsigned NOT NULL,
  `bank_account_number` varchar(25) DEFAULT NULL,
  `bank_account_name` varchar(50) DEFAULT NULL,
  `date` date NOT NULL,
  `value` double(14,2) DEFAULT NULL,
  `status` tinyint NOT NULL COMMENT '1:created;9:invalid;10:valid',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
