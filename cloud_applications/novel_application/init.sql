CREATE TABLE `novel` (
  `id` varchar(36) NOT NULL,
  `title` varchar(50) NOT NULL,
  `author` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `chapter` (
  `id` varchar(36) NOT NULL,
  `novel_id` varchar(36) NOT NULL,
  `title` varchar(50) NOT NULL,
  `index` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `chapter_novel_id_idx` (`novel_id`),
  CONSTRAINT `chapter_novel_id` FOREIGN KEY (`novel_id`) REFERENCES `novel` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `novel_source` (
  `id` varchar(36) NOT NULL,
  `source_id` varchar(30) NOT NULL,
  `novel_id` varchar(36) DEFAULT NULL,
  `preference` int(11) NOT NULL,
  `vip` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_source_index` (`source_id`,`novel_id`),
  UNIQUE KEY `unique_preference_index` (`novel_id`,`preference`),
  CONSTRAINT `source_novel_id` FOREIGN KEY (`novel_id`) REFERENCES `novel` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `chapter_source` (
  `novel_source_id` varchar(36) DEFAULT NULL,
  `chapter_id` varchar(36) NOT NULL,
  `url` varchar(100) DEFAULT NULL,
  `text_length` int(11) DEFAULT '0',
  UNIQUE KEY `url_UNIQUE` (`url`),
  KEY `chapter_source_novel_source_id_idx` (`novel_source_id`),
  KEY `chapter_source_chapter_id_idx` (`chapter_id`),
  CONSTRAINT `chapter_source_chapter_id` FOREIGN KEY (`chapter_id`) REFERENCES `chapter` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `chapter_source_novel_source_id` FOREIGN KEY (`novel_source_id`) REFERENCES `novel_source` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci