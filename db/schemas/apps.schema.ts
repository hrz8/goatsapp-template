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

export type AppSettings = {
  incomingURL: string;
};

export const apps = pgTable('apps', {
  id: uuid('id')
    .primaryKey()
    .default(sql`gen_random_uuid()`),
  name: varchar('name', {length: 50}).notNull(),
  description: text('description'),
  settings: jsonb('settings')
    .$type<AppSettings>()
    .default(sql`'{}'::jsonb`),
  createdAt: timestamp('created_at', {withTimezone: true}).default(sql`now()`),
  updatedAt: timestamp('updated_at', {withTimezone: true}).default(sql`now()`),
});

export const appRelations = relations(apps, ({one, many}) => ({
  devices: many(devices),
}));