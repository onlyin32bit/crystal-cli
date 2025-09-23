package wd

import "crystal-cli/utils"

const initContent string = `# Crystal ORM Schema
# This file defines your database models, enums, and settings.` +
	"\n# Edit this file and run `crystal-cli generate` or `crystal-cli migrate`\n" +
	`# to create migrations and generate Go code.

provider: postgres # Database provider (postgres, mysql, sqlite)
source: env.DATABASE_URL # The connection string, or path to the .db file (sqlite)

enums:
 Role: [ADMIN, USER]

models:
  baseModel: &baseModel
    id:
      $type: int        # field data type
      $id: true         # primary key
      $default: AUTOINC # auto increment
    createdAt:
      $type: dateTime
      $default: NOW     # default to current timestamp
    updatedAt:
      $type: dateTime
      $onUpdate: NEWDATETIME  # auto-updated on changes
    deletedAt:
      $type: dateTime?  # nullable

  User:
    <<: *baseModel # reusable Fields
    username:
      $type: string
      $unique: true
    passwordHash:
      $type: string
    role:
      $type: Role
      $default: USER

  Profile:
    id:
      $type: int
      $id: true
      $default: AUTOINC
    bio:
      $type: string
    user:
      $type: User
      $relation:
        fields: [userId]
        references: [id]
    userId:
      $type: int
      $unique: true
`

func InitCrystal() error {
	if err := utils.CreateFolder("crystal"); err != nil {
		return err
	}
	if err := utils.CreateFileWithContent("crystal/schema.yaml", initContent); err != nil {
		return err
	}
	return nil
}
