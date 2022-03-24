USE `jim`;
ALTER TABLE `message_content`
    ALTER `id` DROP DEFAULT;
ALTER TABLE `message_content`
    CHANGE COLUMN `id` `id` BIGINT UNSIGNED NOT NULL FIRST;