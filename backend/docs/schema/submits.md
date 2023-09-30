# submits

## Description

<details>
<summary><strong>Table Definition</strong></summary>

```sql
CREATE TABLE `submits` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `language_submits` bigint(20) NOT NULL,
  `task_submits` bigint(20) NOT NULL,
  `status` varchar(255) DEFAULT NULL,
  `exec_time` bigint(20) DEFAULT NULL,
  `exec_memory` bigint(20) DEFAULT NULL,
  `score` bigint(20) DEFAULT NULL,
  `submitted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `user_submits` bigint(20) NOT NULL,
  `contest_submits` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `submits_contests_submits` (`contest_submits`),
  KEY `submits_languages_submits` (`language_submits`),
  KEY `submits_tasks_submits` (`task_submits`),
  KEY `submits_users_submits` (`user_submits`),
  CONSTRAINT `submits_contests_submits` FOREIGN KEY (`contest_submits`) REFERENCES `contests` (`id`) ON DELETE SET NULL,
  CONSTRAINT `submits_languages_submits` FOREIGN KEY (`language_submits`) REFERENCES `languages` (`id`) ON DELETE NO ACTION,
  CONSTRAINT `submits_tasks_submits` FOREIGN KEY (`task_submits`) REFERENCES `tasks` (`id`) ON DELETE NO ACTION,
  CONSTRAINT `submits_users_submits` FOREIGN KEY (`user_submits`) REFERENCES `users` (`id`) ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
```

</details>

## Columns

| Name | Type | Default | Nullable | Extra Definition | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | ---------------- | -------- | ------- | ------- |
| id | bigint(20) |  | false | auto_increment | [testcase_results](testcase_results.md) [user_submits](user_submits.md) |  |  |
| language_submits | bigint(20) |  | false |  |  | [languages](languages.md) |  |
| task_submits | bigint(20) |  | false |  |  | [tasks](tasks.md) |  |
| status | varchar(255) | NULL | true |  |  |  |  |
| exec_time | bigint(20) | NULL | true |  |  |  |  |
| exec_memory | bigint(20) | NULL | true |  |  |  |  |
| score | bigint(20) | NULL | true |  |  |  |  |
| submitted_at | timestamp | NULL | true |  |  |  |  |
| created_at | timestamp | NULL | true |  |  |  |  |
| updated_at | timestamp | NULL | true |  |  |  |  |
| user_submits | bigint(20) |  | false |  |  | [users](users.md) |  |
| contest_submits | bigint(20) | NULL | true |  |  | [contests](contests.md) |  |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| PRIMARY | PRIMARY KEY | PRIMARY KEY (id) |
| submits_contests_submits | FOREIGN KEY | FOREIGN KEY (contest_submits) REFERENCES contests (id) |
| submits_languages_submits | FOREIGN KEY | FOREIGN KEY (language_submits) REFERENCES languages (id) |
| submits_tasks_submits | FOREIGN KEY | FOREIGN KEY (task_submits) REFERENCES tasks (id) |
| submits_users_submits | FOREIGN KEY | FOREIGN KEY (user_submits) REFERENCES users (id) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| submits_contests_submits | KEY submits_contests_submits (contest_submits) USING BTREE |
| submits_languages_submits | KEY submits_languages_submits (language_submits) USING BTREE |
| submits_tasks_submits | KEY submits_tasks_submits (task_submits) USING BTREE |
| submits_users_submits | KEY submits_users_submits (user_submits) USING BTREE |
| PRIMARY | PRIMARY KEY (id) USING BTREE |

## Relations

![er](submits.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)