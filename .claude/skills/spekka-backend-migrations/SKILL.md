---
name: Backend Migrations
description: Your approach to handling backend migrations. Use this skill when working on files where backend migrations comes into play.
---

# Backend Migrations

This Skill provides Claude Code with specific guidance on how it should handle backend migrations.

## Instructions

- **Reversible Migrations**: Always implement rollback/down methods to enable safe migration reversals
- **Small, Focused Changes**: Keep each migration focused on a single logical change for clarity and easier troubleshooting
- **Zero-Downtime Deployments**: Consider deployment order and backwards compatibility for high-availability systems
- **Separate Schema and Data**: Keep schema changes separate from data migrations for better rollback safety
- **Index Management**: Create indexes on large tables carefully, using concurrent options when available to avoid locks
- **Naming Conventions**: Use clear, descriptive names that indicate what the migration does
- **Version Control**: Always commit migrations to version control and never modify existing migrations after deployment
