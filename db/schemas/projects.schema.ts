import {relations, sql} from 'drizzle-orm';
import {
  jsonb,
  pgTable,
  text,
  timestamp,
  uuid,
  varchar,
} from 'drizzle-orm/pg-core';
import { devices } from './devices.schema';

export type ProjectSettings = {
  incomingURL: string;
};

export const projects = pgTable('projects', {
  id: uuid('id')
    .primaryKey()
    .default(sql`gen_random_uuid()`),
  name: varchar('name', {length: 50}).notNull(),
  alias: varchar('alias', {length: 50}).unique().notNull(),
  description: text('description'),
  settings: jsonb('settings')
    .$type<ProjectSettings>()
    .default(sql`'{}'::jsonb`),
  createdAt: timestamp('created_at', {withTimezone: true}).default(sql`now()`),
  updatedAt: timestamp('updated_at', {withTimezone: true}).default(sql`now()`),
});

export const projectRelations = relations(projects, ({one, many}) => ({
  devices: many(devices),
}));