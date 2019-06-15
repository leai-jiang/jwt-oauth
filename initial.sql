SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for OAuthGithub
-- ----------------------------
DROP TABLE IF EXISTS `OAuthGithub`;
CREATE TABLE `OAuthGithub`  (
  `id` bigint(63) NOT NULL,
  `name` char(63) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `avatar` char(63) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `company` char(63) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `blog` char(63) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `email` char(63) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `location` varchar(63) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `create_time` datetime(0) DEFAULT NULL,
  `update_time` datetime(0) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
