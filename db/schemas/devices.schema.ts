import {relations, sql} from 'drizzle-orm';
import {
  boolean as bool,
  jsonb,
  pgTable,
  timestamp,
  uuid,
  varchar,
} from 'drizzle-orm/pg-core';

import {apps} from './apps.schema';

export type DeviceCustomAttributes = Record<string, any>;

export const devices = pgTable('devices', {
  id: uuid('id')
    .primaryKey()
    .default(sql`gen_random_uuid()`),
  appId: uuid('app_id')
    .references(() => apps.id, {onDelete: 'cascade'})
    .notNull(),
  integrationId: varchar('integration_id', {length: 50}).notNull(),
  name: varchar('name', {length: 50}).notNull(),
  phoneNumber: varchar('phone_number', {length: 20}),
  description: varchar('description', {length: 140}),
  isActive: bool('is_active').notNull(),
  customAttributes: jsonb('custom_attributes')
    .$type<DeviceCustomAttributes>()
    .default(sql`'{}'::jsonb`),
  createdAt: timestamp('created_at', {withTimezone: true}).default(sql`now()`),
  updatedAt: timestamp('updated_at', {withTimezone: true}).default(sql`now()`),
});

export const whatsappDeviceRelations = relations(
  devices,
  ({one}) => ({
    app: one(apps, {
      fields: [devices.appId],
      references: [apps.id],
    }),
  }),
);