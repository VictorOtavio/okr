CREATE TABLE IF NOT EXISTS `key_results` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `objective_id` INTEGER NOT NULL,
    `description` TEXT NOT NULL,
    `done` INTEGER NOT NULL DEFAULT 0
);
